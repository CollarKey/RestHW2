package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var task string
var DB *gorm.DB

type requestBody struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func InitDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database, %v ", err)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []requestBody

	answer := DB.Find(&tasks)
	if answer.Error != nil {
		http.Error(w, "Failed to get the tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var body requestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	process := DB.Create(&body)
	if process.Error != nil {
		http.Error(w, "Could not create the task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Task added successfully: %v\n", task)
	json.NewEncoder(w).Encode(body)

}

func main() {
	InitDB()

	// Здесь используйте &Task{}
	DB.AutoMigrate(&requestBody{})

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", getHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/tasks", postHandler).Methods(http.MethodPost)

	log.Println("Server started at localhost:8080")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
