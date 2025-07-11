package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Struct untuk POST request
type Person struct {
	Name    string   `json:"name"`
	Job     string   `json:"job"`
	Citizen string   `json:"citizen"`
	Hobbies []string `json:"hobbies"`
}

// Struct untuk data random
type RandomData struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

// Standard response wrapper
type StandardResponse struct {
	StatusCode   int         `json:"status_code"`
	ResponseBody interface{} `json:"response_body"`
	Description  string      `json:"description"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/all", allHandler)
	http.HandleFunc("/random", randomHandler)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler: GET /health (Always 200)
func healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := StandardResponse{
		StatusCode:   200,
		ResponseBody: "OK",
		Description:  "Health check passed",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// Handler: POST /submit
func submitHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		resp := StandardResponse{
			StatusCode:   http.StatusMethodNotAllowed,
			ResponseBody: "Invalid method",
			Description:  "Only POST method is allowed",
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(resp)
		return
	}

	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		resp := StandardResponse{
			StatusCode:   http.StatusBadRequest,
			ResponseBody: "Invalid body",
			Description:  "Request body format is incorrect",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := StandardResponse{
		StatusCode: http.StatusOK,
		ResponseBody: map[string]interface{}{
			"message": "Data received",
			"data":    person,
		},
		Description: "User data submitted successfully",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// Handler: GET /all
func allHandler(w http.ResponseWriter, r *http.Request) {
	var data []RandomData
	for i := 0; i < 15; i++ {
		data = append(data, RandomData{
			ID:    i + 1,
			Value: randomString(),
		})
	}

	resp := StandardResponse{
		StatusCode:   http.StatusOK,
		ResponseBody: data,
		Description:  "List of 15 random JSON entries generated",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// Handler: GET /random
func randomHandler(w http.ResponseWriter, r *http.Request) {
	status := rand.Intn(2) == 0
	code := rand.Intn(900) + 100
	msg := randomString()

	resp := StandardResponse{
		StatusCode: 200,
		ResponseBody: map[string]interface{}{
			"status": status,
			"code":   code,
			"msg":    msg,
		},
		Description: "Random JSON data generated from static /random endpoint",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// Helper: Random string generator
func randomString() string {
	words := []string{"apple", "banana", "cat", "dog", "echo", "fish", "go", "home", "ice", "jazz"}
	return words[rand.Intn(len(words))]
}
