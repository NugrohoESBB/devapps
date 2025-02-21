package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func RegisHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/inputUsers.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/admin/dashboard.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT id, d, t, tn, s FROM logsessions ORDER BY id DESC LIMIT 5")
	if err != nil {
		http.Error(w, "Failed to fetch users data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var dash_logs []APILogSessions
	for rows.Next() {
		var u APILogSessions
		rows.Scan(&u.ID, &u.D, &u.T, &u.TN, &u.S)
		dash_logs = append(dash_logs, u)
	}

	tmpl.Execute(w, dash_logs)
}

func DashboardHandler_User(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/user/dashboardUsers.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT id, d, t, n, e, lt, ln FROM users LIMIT 5")
	if err != nil {
		http.Error(w, "Failed to fetch users data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var u_users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.D, &u.T, &u.N, &u.E, &u.LT, &u.LN)
		u_users = append(u_users, u)
	}

	tmpl.Execute(w, u_users)
}

func LogDataHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/admin/logData.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT id, d, t, n, k, l, i, f, a FROM logs")
	if err != nil {
		http.Error(w, "Failed to fetch users data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var logs []APILogData
	for rows.Next() {
		var u APILogData
		rows.Scan(&u.ID, &u.D, &u.T, &u.N, &u.K, &u.L, &u.I, &u.F, &u.A)
		logs = append(logs, u)
	}

	tmpl.Execute(w, logs)
}

func LogSessionsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/admin/logSessions.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT id, d, t, tn, s FROM logsessions")
	if err != nil {
		http.Error(w, "Failed to fetch users data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sessions_logs []APILogSessions
	for rows.Next() {
		var u APILogSessions
		rows.Scan(&u.ID, &u.D, &u.T, &u.TN, &u.S)
		sessions_logs = append(sessions_logs, u)
	}

	tmpl.Execute(w, sessions_logs)
}

func InformationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/admin/informationLog.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT id, d, t, n, e, lt, ln, p, r FROM users")
	if err != nil {
		http.Error(w, "Failed to fetch users data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var information_logs []User
	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.D, &u.T, &u.N, &u.E, &u.LT, &u.LN, &u.P, &u.R)
		information_logs = append(information_logs, u)
	}

	tmpl.Execute(w, information_logs)
}

func LogTasksHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/admin/logTask.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT id, r, dc, rt FROM logtasks ORDER BY id DESC")
	if err != nil {
		http.Error(w, "Failed to fetch users data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var task_logs []APILogTasks
	for rows.Next() {
		var u APILogTasks
		rows.Scan(&u.ID, &u.R, &u.DC, &u.RT)
		task_logs = append(task_logs, u)
	}

	tmpl.Execute(w, task_logs)
}

func InvoiceHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/admin/invoice.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
