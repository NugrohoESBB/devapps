package handlers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	D        string `json:"d"`
	T        string `json:"t"`
	N        string `json:"n"`
	E        string `json:"e"`
	LT       string `json:"lt"`
	LN       string `json:"ln"`
	P        string `json:"p"`
	R        string `json:"r"`
}

// API Handler - For JSON or FORM a Users Input
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost || r.Method == http.MethodGet {
		var u User

		if r.Method == http.MethodPost {
			// Coba baca sebagai JSON
			if r.Header.Get("Content-Type") == "application/json" {
				err := json.NewDecoder(r.Body).Decode(&u)
				if err != nil {
					http.Error(w, "Invalid JSON", http.StatusBadRequest)
					return
				}
			} else {
				// Jika bukan JSON, baca dari form-data atau query parameters
				err := r.ParseForm() // Pastikan bisa membaca form-data
				if err != nil {
					http.Error(w, "Gagal membaca form", http.StatusBadRequest)
					return
				}

				// Ambil data dari form atau query parameters
				u.N 	= r.FormValue("n")
				u.E 	= r.FormValue("e")
				u.LT 	= r.FormValue("lt")
				u.LN 	= r.FormValue("ln")
				u.P 	= r.FormValue("p")
				u.R 	= r.FormValue("r")

				// Pastikan semua field diisi
				if u.N == "" || u.E == "" || u.LT == "" || u.LN == "" || u.P == "" || u.R == "" {
					http.Error(w, "Semua parameter harus diisi", http.StatusBadRequest)
					return
				}
			}
		} else { // GET request tetap pakai query parameters
			u.N 	= r.URL.Query().Get("n")
			u.E 	= r.URL.Query().Get("e")
			u.LT 	= r.URL.Query().Get("lt")
			u.LN 	= r.URL.Query().Get("ln")
			u.P 	= r.URL.Query().Get("p")
			u.R 	= r.URL.Query().Get("r")

			if u.N == "" || u.E == "" || u.LT == "" || u.LN == "" || u.P == "" || u.R == "" {
				http.Error(w, "Semua parameter harus diisi", http.StatusBadRequest)
				return
			}
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.P), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Gagal mengenkripsi password", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO users (d, t, n, e, lt, ln, p, r) VALUES (NOW(), NOW(), ?, ?, ?, ?, ?, ?)", u.N, u.E, u.LT, u.LN, hashedPassword, u.R)
		if err != nil {
			http.Error(w, "Gagal menambah pengguna", http.StatusInternalServerError)
			return
		}

		// Redirect ke halaman sukses
		http.Redirect(w, r, "/regis", http.StatusSeeOther)
	}
}

// Handler GET JSON - For Maps Dashboard
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, d, t, n, e, lt, ln, p, r FROM users LIMIT 2")
	if err != nil {
		http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.D, &u.T, &u.N, &u.E, &u.LT, &u.LN, &u.P, &u.R)
		if err != nil {
			http.Error(w, "Gagal membaca data", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Handler GET JSON - For Users Chart Dashboard
func GetUserStatsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT DATE(d) AS date, COUNT(*) AS count FROM users GROUP BY DATE(d) ORDER BY date ASC LIMIT 7")
	if err != nil {
		http.Error(w, "Gagal mengambil statistik pengguna", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var data []map[string]interface{}
	for rows.Next() {
		var date string
		var count int
		err := rows.Scan(&date, &count)
		if err != nil {
			http.Error(w, "Gagal membaca data", http.StatusInternalServerError)
			return
		}
		data = append(data, map[string]interface{}{"date": date, "count": count})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Handler Authorization - For Role Users
func GetUserRoleHandler(w http.ResponseWriter, r *http.Request) {
	role, ok := r.Context().Value("r").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"r": role})
}

// For hashing p before inserting it into to DB
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
