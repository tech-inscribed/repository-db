package main

import (
	"fmt"
	"net/http"

	"github.com/techinscribed/repository-db/controllers"
	"github.com/techinscribed/repository-db/repositories"
	"github.com/techinscribed/repository-db/sqldb"
)

func main() {
	db := sqldb.ConnectDB()

	// Create repos
	userRepo := repositories.NewUserRepo(db)

	h := controllers.NewBaseHandler(userRepo)

	http.HandleFunc("/", h.HelloWorld)

	s := &http.Server{
		Addr: fmt.Sprintf("%s:%s", "localhost", "5000"),
	}

	s.ListenAndServe()

}
