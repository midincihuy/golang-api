package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Produk struct {
	ID int `json:"id"`
	Nama string `json:"nama"`
	Harga int `json:"harga"`
	Stok int `json:"stok"`
}

var produk = []Produk{
	{ID: 1, Nama: "Mie Ayam", Harga:15000, Stok:100},
	{ID: 2, Nama: "Mie Bakso", Harga:15000, Stok:100},
	{ID: 3, Nama: "Mie Ayam Bakso", Harga:20000, Stok:100},
}

func main(){

	// GET Produk 
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request){
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}

			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(produkBaru)
		}
	})

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