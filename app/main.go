package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	_http "2023_2_Holi/auth/delivery/http"
	"2023_2_Holi/auth/repository/postgresql"
	"2023_2_Holi/auth/usecase"
	_ "github.com/lib/pq"
)

func fromEnv() string {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return ""
	}

	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
}

// @title Netfilx API
// @version 1.0
// @description API of the nelfix project by holi

// @contact.name Alex Chinaev
// @contact.url https://vk.com/l.chinaev
// @contact.email ax.chinaev@yandex.ru

// @license.name AS IS (NO WARRANTY)

// @host 127.0.0.1E
// @schemes Zhttp
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("starting connect to db")
	db, err := sql.Open("postgres", fromEnv())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	sessionRepository := postgresql.NewSessionPostgresqlRepository(db)
	authRepository := postgresql.NewAuthPostgresqlRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepository, sessionRepository)
	_http.NewAuthHandler(router, authUsecase)

	fmt.Println("starting server at :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server stopped")
}
