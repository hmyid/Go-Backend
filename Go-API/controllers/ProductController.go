package controllers

/*
Controllers merupakan implementasi dalam bentuk kode untuk dapat berkomunikasi
dengan client dan server
*/

import (
	"encoding/json" // Package casting json to go values or go values to json
	"log"
	"net/http" // Package HTTP client and Server
	"strings"

	"strconv" // Memakai Data Type Convertion
	"sync"    // Memakai Mutex
)

type Product struct {
	ID     int
	Name   string
	Weight int
	Price  int
}

// Middleware
func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf("Request %s %s", req.Method, req.URL.Path)
		next(res, req)
	})
}

// Auth
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		uname, pwd, ok := req.BasicAuth()

		if !ok {
			res.Write([]byte("Username Tidak Boleh Kosong"))
		}

		if uname == "admin" && pwd == "123" {
			next.ServeHTTP(res, req)
			return
		}

		res.Write([]byte("Username dan Password Tidak Sesuai"))
	})
}

// Data Dummy Temporary. If stop running this server, then the data not saved
func Products() []Product {
	prod := []Product{
		{1, "Television", 10, 1000},
		{2, "Icebox", 20, 2000},
		{3, "Microwave", 5, 4000},
		{4, "Fan", 3, 500},
	}
	return prod
}

var products = Products()

// Mencegah Race Condition
var mu sync.Mutex

// Fetch Data
func GetProduct(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {

		// Casting and Parsing. Mengubah data Go Values menjadi tipe json
		productCast, err := json.Marshal(products)

		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
		}

		// Memberikan informasi bahwa data tersebut json type
		res.Header().Set("Content-Type", "application/json")
		// Message Status jika berhasil dibaca/ditulis
		res.WriteHeader(http.StatusOK)
		// Menampilkan/Menulisan data yang sudah berhasil diubah ke json
		res.Write(productCast)

		return
	}

	// Error
	http.Error(res, "Error", http.StatusNotFound)
}

// Post Data
func PostProduct(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var prodPost Product

	if req.Method == "POST" {
		if req.Header.Get("Content-Type") == "application/json" {

			// Parsing JSON
			DecodeJSON := json.NewDecoder(req.Body)
			err := DecodeJSON.Decode(&prodPost)

			if err != nil {
				log.Fatal(err)
			}

			prodPost.ID = len(products) + 1

			products = append(products, prodPost)
		}

		// Casting Go Values to JSON
		CastJSON, err := json.Marshal(&prodPost)
		if err != nil {
			log.Fatal(err)
		}
		res.WriteHeader(http.StatusCreated)
		res.Write(CastJSON)
		return
	}

	http.Error(res, "Bad Req", http.StatusBadRequest)
}

// Update | PUT Data
func PutProduct(res http.ResponseWriter, req *http.Request) {
	var prodPUT Product

	if req.Method == "PUT" {

		// Decode JSON dari Request Body
		err := json.NewDecoder(req.Body).Decode(&prodPUT)

		if err != nil {
			http.Error(res, "Bad Request", http.StatusBadRequest)
			return
		}

		mu.Lock()
		defer mu.Unlock()

		found := false
		for i := range products {
			if products[i].ID == prodPUT.ID {
				// Update Product
				products[i] = prodPUT
				found = true
				break
			}
		}

		// Jika Product tidak ditemukan
		if !found {
			http.Error(res, "Product Not Found", http.StatusNotFound)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(prodPUT)
		return
	}

	http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// PATCH | Update Field Data
func PatchProduct(res http.ResponseWriter, req *http.Request) {
	if req.Method == "PATCH" {
		idStr := strings.TrimPrefix(req.URL.Path, "/updateproduct/")

		// Ambil data berdasarkan ID
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(res, "Bad Request", http.StatusBadRequest)
			return
		}

		var found *Product
		for i := range products {
			if products[i].ID == id {
				found = &products[i]
			}
		}

		// Decode JSON
		var prodPatch map[string]interface{}
		err = json.NewDecoder(req.Body).Decode(&prodPatch)
		if err != nil {
			http.Error(res, "Invalid Request Body", http.StatusBadRequest)
			return
		}

		// Update Field
		name, ok := prodPatch["Name"].(string)
		if ok {
			found.Name = name
		}

/* 
		weight, ok := prodPatch["Weight"].(int)
		if ok {
			found.Weight = weight
		}

		price, ok := prodPatch["Price"].(int)
		if ok {
			found.Price = price
		}
 */

		// Response Success
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(&found)
		return
	}

	http.Error(res, "Method Not Found", http.StatusNotFound)
}

// Delete Data
func DeleteProduct(res http.ResponseWriter, req *http.Request) {
	if req.Method == "DELETE" {
		idStr := strings.TrimPrefix(req.URL.Path, "/deleteproduct/")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(res, "Bad Request", http.StatusBadRequest)
			return
		}

		// cari Index
		index := -1
		for i := range products {
			if products[i].ID == id {
				index = i
				break
			}
		}

		// Jika Product Tidak Ditemukan
		if index == -1 {
			http.Error(res, "Bad Request", http.StatusBadRequest)
			return
		}

		// Hapus produk dari slice
		products = append(products[:index], products[index+1:]...)

		// Respon sukses
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(map[string]string{"message": "Product deleted successfully"})
		return
	}

	http.Error(res, "Method Not Found", http.StatusNotFound)
} 
