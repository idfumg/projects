package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"myapp/db"
	"myapp/endpoints"
)

const (
	DBHost = "127.0.0.1"
	DBPort = ":3306"
	DBUser = "citizix_user"
	DBPass = "An0thrS3crt"
	DBBase = "citizix_db"
	PORT   = ":8080"
)

func GetRoutes() *mux.Router {
	routes := mux.NewRouter()

	routes.HandleFunc("/page/{guid:[0-9a-zA-Z\\-]+}", endpoints.ServeBlogPage).Methods("GET").Schemes("https")
	routes.HandleFunc("/", endpoints.RedirIndex).Methods("GET").Schemes("https")
	routes.HandleFunc("/home", endpoints.ServeIndex).Methods("GET").Schemes("https")

	routes.HandleFunc("/register", endpoints.RegisterPOST).Methods("POST").Schemes("https")
	routes.HandleFunc("/login", endpoints.LoginPOST).Methods("POST").Schemes("https")

	routes.HandleFunc("/api/pages", endpoints.APIPagesGET).Methods("GET").Schemes("https")
	routes.HandleFunc("/api/pages/{guid:[0-9a-zA-Z\\-]+}", endpoints.APIPageGET).Methods("GET").Schemes("https")

	routes.HandleFunc("/api/comments", endpoints.APIAddCommentPost).Methods("POST").Schemes("https")
	routes.HandleFunc("/api/comments/{id:[\\w\\d\\-]+}", endpoints.APIUpdateCommentPUT).Methods("PUT").Schemes("https")

	return routes
}

func main() {
	db.GetMySQLConnection(DBUser, DBPass, DBHost, DBBase)

	http.Handle("/", GetRoutes())
	http.ListenAndServeTLS(PORT, "certificate.pem", "key.pem", nil)
}
