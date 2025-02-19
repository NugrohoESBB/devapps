package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserAuth struct {
	ID	int    `json:"id"`
	N 	string `json:"n"`
	P 	string `json:"p"`
	R 	string `json:"r"`
}

type LogSessions struct {
	ID	int    `json:"id"`
	TN 	string `json:"tn"`
	S 	string `json:"s"`
}

// LoginHandler: Handles user login and creates a session
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u UserAuth
	json.NewDecoder(r.Body).Decode(&u)

	var storedPassword, role string
	err := db.QueryRow("SELECT id, p, r FROM users WHERE n = ?", u.N).Scan(&u.ID, &storedPassword, &role)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(u.P))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	sessionToken := fmt.Sprintf("session-%d-%d", u.ID, time.Now().Unix())

	_, err = db.Exec("INSERT INTO sessions (u_id, tn, r, t) VALUES (?, ?, ?, NOW())", u.ID, sessionToken, role)
	_, err = db.Exec("INSERT INTO logsessions (tn, s) VALUES (?, ?)", sessionToken, "Login")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Browser
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"r": role,
	})
}

// LogoutHandler: Handles user logout by deleting the session
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}

	result, err := db.Exec("DELETE FROM sessions WHERE tn = ?", cookie.Value)
	if err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}

	// Browser
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
}

// AuthMiddleware: Middleware to validate session token
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var userID int
		var role string
		err = db.QueryRow("SELECT u_id, r FROM sessions WHERE tn = ?", cookie.Value).Scan(&userID, &role)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "r", role)

		next(w, r.WithContext(ctx))
	}
}
