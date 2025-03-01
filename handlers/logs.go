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

type APILogTasks struct {
	ID	int    `json:"id"`
	D 	string `json:"d"`
	T 	string `json:"t"`
	R 	string `json:"r"`
	DC 	string `json:"dc"`
	RT 	string `json:"rt"`
}

// ======================================== LOGS DATA

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

// Handler GET JSON - For JPL Chart Dashboard
func GetLogDataStatsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT d, k, i, f, a FROM logs ORDER BY d DESC LIMIT 7")
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

// ======================================== SESSIONS LOG

// Handler GET JSON - For User Login Chart Dashboard
func GetLoginStatsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT DATE(d) AS date, COUNT(*) AS count FROM logsessions GROUP BY DATE(d) ORDER BY date DESC LIMIT 7")
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

// ======================================== TASK LOG

// API Handler - JSON or FORM Data Input
func AddLogTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost || r.Method == http.MethodGet {
		var u APILogTasks

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
				
				u.R  = r.FormValue("r")
				u.DC = r.FormValue("dc")
				u.RT = r.FormValue("rt")

				if u.R == "" || u.DC == "" || u.RT == "" {
					http.Error(w, "Semua parameter harus diisi", http.StatusBadRequest)
					return
				}
			}
		} else {
			// GET request tetap pakai query parameters
			u.R  = r.URL.Query().Get("r")
			u.DC = r.URL.Query().Get("dc")
			u.RT = r.URL.Query().Get("rt")

			if u.R == "" || u.DC == "" || u.RT == "" {
				http.Error(w, "Semua parameter harus diisi", http.StatusBadRequest)
				return
			}
		}

		_, err := db.Exec("INSERT INTO logtasks (r, dc, rt) VALUES (?, ?, ?)", &u.R, &u.DC, &u.RT)
		if err != nil {
			http.Error(w, "Gagal menambah pengguna", http.StatusInternalServerError)
			return
		}

		// Redirect ke halaman sukses
		http.Redirect(w, r, "/logtasks", http.StatusSeeOther)
	}
}

// API Handler - Delete JSON or FORM Data
func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	// Ambil ID dari form
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	// Eksekusi query DELETE
	_, err = db.Exec("DELETE FROM logtasks WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Gagal menghapus data", http.StatusInternalServerError)
		return
	}

	// Berikan respons sukses
	w.WriteHeader(http.StatusOK)
}

// API Handler - Notifications From Task Log
func GetNotificationsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, d, t, r, dc, rt FROM logtasks ORDER BY id DESC")
	if err != nil {
		http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var notif []APILogTasks
	for rows.Next() {
		var u APILogTasks
		err := rows.Scan(&u.ID, &u.D, &u.T, &u.R, &u.DC, &u.RT)
		if err != nil {
			http.Error(w, "Gagal membaca data", http.StatusInternalServerError)
			return
		}
		notif = append(notif, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notif)
}
