package main

import (
	"fmt"
	"log"
	"net/http"

	"userauthapps/handlers"
)

func main() {
	handlers.InitDB()

	// Public routes
	http.HandleFunc("/", handlers.HomeHandler)

	// Authentication
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/regis", handlers.RegisHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)

	// Protected pages
	http.HandleFunc("/dashboard", handlers.AuthMiddleware(handlers.DashboardHandler))
	http.HandleFunc("/logs", handlers.AuthMiddleware(handlers.LogDataHandler))
	http.HandleFunc("/logsessions", handlers.AuthMiddleware(handlers.LogSessionsHandler))
	http.HandleFunc("/informationLog", handlers.AuthMiddleware(handlers.InformationHandler))
	http.HandleFunc("/logtasks", handlers.AuthMiddleware(handlers.LogTasksHandler))
	http.HandleFunc("/invoice", handlers.AuthMiddleware(handlers.InvoiceHandler))

	// Protected pages - User
	http.HandleFunc("/dashboard-user", handlers.AuthMiddleware(handlers.DashboardHandler))
	http.HandleFunc("/logs-user", handlers.AuthMiddleware(handlers.LogDataHandler))
	http.HandleFunc("/logtasks-user", handlers.AuthMiddleware(handlers.LogTasksHandler))

	// =============== ALL API ===============
	// User API
	http.HandleFunc("/addUserData", handlers.AddUserHandler)
	http.HandleFunc("/api/users", handlers.GetUsersHandler)
	http.HandleFunc("/api/user-stats", handlers.GetUserStatsHandler) // Optional
	http.HandleFunc("/api/login-stats", handlers.GetLoginStatsHandler)
	http.HandleFunc("/api/user-role", handlers.AuthMiddleware(handlers.GetUserRoleHandler))

	// Log Data API
	http.HandleFunc("/addLogData", handlers.AddLogDataHandler)
	http.HandleFunc("/api/logdata-stats", handlers.GetLogDataStatsHandler)

	// Task API
	http.HandleFunc("/addLogTask", handlers.AddLogTaskHandler)
	http.HandleFunc("/deleteTask", handlers.DeleteTasksHandler)
	http.HandleFunc("/updateStatus", handlers.AuthMiddleware(handlers.UpdateTaskStatusHandler))
	http.HandleFunc("/api/notificationTask", handlers.AuthMiddleware(handlers.GetNotificationsHandler))

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server started at :9926")
	log.Fatal(http.ListenAndServe(":9926", nil))
}
