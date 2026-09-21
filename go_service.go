// 這是用go 內建的standard library(net/http)，沒有特別名字就是純手刻
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func handleGoRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Go 服務已收到請求")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: "Go 收到了"})
}

func main() {
	http.HandleFunc("/api/go", handleGoRequest)

	log.Println("Go Service listening on :8002")
	log.Fatal(http.ListenAndServe(":8002", nil))
}
