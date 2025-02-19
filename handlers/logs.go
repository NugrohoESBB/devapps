package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type APILogData struct {
	ID 	int     `json:"id"`
	D 	string  `json:"d"`
	T 	string 	`json:"t"`
	N	string  `json:"n"`
	K 	float64 `json:"k"`
	L 	string  `json:"l"`
	I 	float64 `json:"i"`
	F	float64 `json:"f"`
	A 	float64 `json:"a"`
}

type APILogSessions struct {
	ID	int    `json:"id"`
	D 	string `json:"d"`
	T 	string `json:"t"`
	TN 	string `json:"tn"`
	S 	string `json:"s"`
}

// API Handler - JSON or FORM Data Input
func AddLogDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost || r.Method == http.MethodGet {
		var u APILogData

		if r.Method == http.MethodPost {
			if r.Header.Get("Content-Type") == "application/json" {
				err := json.NewDecoder(r.Body).Decode(&u)
				if err != nil {
					http.Error(w, "Invalid JSON", http.StatusBadRequest)
					return
				}
			} else {
				// Jika bukan JSON, baca dari form-data atau query parameters
				err := r.ParseForm()
				if err != nil {
					http.Error(w, "Gagal membaca form", http.StatusBadRequest)
					return
				}
				
				u.N = r.FormValue("n")
				u.K = parseFloat(r.FormValue("k"))
				u.L = r.FormValue("l")
				u.I = parseFloat(r.FormValue("i"))
				u.F = parseFloat(r.FormValue("f"))
				u.A = parseFloat(r.FormValue("a"))

				if u.N == "" || u.K == 0 || u.L == "" || u.I == 0 || u.F == 0 || u.A == 0 {
					http.Error(w, "Semua parameter harus diisi", http.StatusBadRequest)
					return
				}
			}
		} else {
			// GET request tetap pakai query parameters
			u.N = r.URL.Query().Get("n")
			u.K = parseFloat(r.URL.Query().Get("k"))
			u.L = r.URL.Query().Get("l")
			u.I = parseFloat(r.URL.Query().Get("i"))
			u.F = parseFloat(r.URL.Query().Get("f"))
			u.A = parseFloat(r.URL.Query().Get("a"))

			if u.N == "" || u.K == 0 || u.L == "" || u.I == 0 || u.F == 0 || u.A == 0 {
				http.Error(w, "Semua parameter harus diisi", http.StatusBadRequest)
				return
			}
		}

		_, err := db.Exec("INSERT INTO logs (n, k, l, i, f, a) VALUES (?, ?, ?, ?, ?, ?)", &u.N, &u.K, &u.L, &u.I, &u.F, &u.A)
		if err != nil {
			http.Error(w, "Gagal menambah pengguna", http.StatusInternalServerError)
			return
		}

		// Redirect ke halaman sukses
		http.Redirect(w, r, "/logs", http.StatusSeeOther)
	}
}

// Handler GET JSON - For JPL Chart
func GetLogDataStatsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT d, k, i, f, a FROM logs ORDER BY d ASC LIMIT 7")
	if err != nil {
		http.Error(w, "Gagal mengambil data log", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var data []map[string]interface{}
	for rows.Next() {
		var d string
		var k float64
		var i float64
		var f float64
		var a float64
		err := rows.Scan(&d, &k, &i, &f, &a)
		if err != nil {
			http.Error(w, "Gagal membaca data", http.StatusInternalServerError)
			return
		}
		data = append(data, map[string]interface{}{
			"d"	: d,
			"k"	: k,
			"i"	: i,
			"f"	: f,
			"a"	: a,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Fungsi untuk mengonversi string ke float64
func parseFloat(value string) float64 {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return f
}
