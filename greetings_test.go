package main

import (
    "os"
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "userauthapps/handlers"
    _ "github.com/go-sql-driver/mysql"
)

func TestMain(m *testing.M) {
    handlers.InitDB()

    code := m.Run()

    if handlers.GetDB() != nil {
        handlers.GetDB().Close()
    }

    os.Exit(code)
}

func TestHomeHandler_GET(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    recorder := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.HomeHandler)
    handler.ServeHTTP(recorder, req)

    if status := recorder.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }
}

func TestLoginHandler_POST(t *testing.T) {
    UserAuth := map[string]string{
        "n": "testuser",
        "p": "password123",
    }
    jsonData, _ := json.Marshal(UserAuth)

    req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    recorder := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.LoginHandler)
    handler.ServeHTTP(recorder, req)

    if recorder.Code != http.StatusOK && recorder.Code != http.StatusUnauthorized {
        t.Errorf("Expected status OK or Unauthorized, got %v", recorder.Code)
    }
}

func TestLogoutHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/logout", nil)
    if err != nil {
        t.Fatal(err)
    }

    recorder := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.LogoutHandler)
    handler.ServeHTTP(recorder, req)

    if recorder.Code != http.StatusUnauthorized && recorder.Code != http.StatusOK {
        t.Errorf("Expected status Unauthorized or OK, got %v", recorder.Code)
    }
}

func TestAddUserHandler_POST(t *testing.T) {
    formData := strings.NewReader("n=-&e=-&lt=-&ln=-&p=-")
    req, err := http.NewRequest("POST", "/addUserData", formData)
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    recorder := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.AddUserHandler)
    handler.ServeHTTP(recorder, req)

    if recorder.Code != http.StatusSeeOther {
        t.Errorf("Expected redirect status, got %v", recorder.Code)
    }
}

func TestAddUserHandlerJSON_POST(t *testing.T) {
    user := handlers.User{N: "-", E: "-", LT: "-", LN: "-", P: "-"}
    jsonData, _ := json.Marshal(user)

    req, err := http.NewRequest("POST", "/addUserData", bytes.NewBuffer(jsonData))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    recorder := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.AddUserHandler)
    handler.ServeHTTP(recorder, req)

    if status := recorder.Code; status != http.StatusSeeOther {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
    }
}

func TestGetUsersHandler_GET(t *testing.T) {
    req, err := http.NewRequest("GET", "/api/users", nil)
    if err != nil {
        t.Fatal(err)
    }

    recorder := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.GetUsersHandler)
    handler.ServeHTTP(recorder, req)

    if status := recorder.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var users []handlers.User
    if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
        t.Errorf("failed to parse response JSON: %v", err)
    }
}

func TestAddLogDataHandler_POST(t *testing.T) {
    logData := handlers.APILogData{N: "-", K: 0, L: "-", I: 0, F: 0, A: 0}
    jsonData, _ := json.Marshal(logData)

    req, err := http.NewRequest("POST", "/addLogData", bytes.NewBuffer(jsonData))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    recorder := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.AddLogDataHandler)
    handler.ServeHTTP(recorder, req)

    if status := recorder.Code; status != http.StatusSeeOther {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
    }
}

func TestGetLogDataStatsHandler_GET(t *testing.T) {
    req, err := http.NewRequest("GET", "/api/logdata-stats", nil)
    if err != nil {
        t.Fatal(err)
    }

    recorder := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.GetLogDataStatsHandler)
    handler.ServeHTTP(recorder, req)

    if status := recorder.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }
}
