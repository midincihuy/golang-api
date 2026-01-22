package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status" : "OK",
			"message" : "API Running Well",
		})
	})
	fmt.Println("Server running di localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		fmt.Println("Gagal Running")
	}
}