package main

import (
	"Go-API/routers"
	"log" // Package message server response, if server get a request log can handle
	"net/http"
)

func main() {
	routers.RegisterRouters() // Sudah terdaftar

	log.Println("Server berjalan di http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}
