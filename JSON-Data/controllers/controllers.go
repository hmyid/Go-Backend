package controllers

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

/*
Kasus ini merupakan Decode (deserialization/unmarshal) karena kita mengubah format JSON (string)
menjadi struktur data Go yang bisa digunakan dalam kode.
*/

func EmployeeControllerDecode() {
	// Casting
	var ObjectEmployees = `{"name" : "Hamid", "age": 20}`
	var formatByteJSON = []byte(ObjectEmployees)
	
	// Data Pertama Menggunakan Struct
	var data Employee
	json.Unmarshal(formatByteJSON, &data)

	// Data Kedua Menggunakan Map Interface/Any
	var data1 map[string]interface{}
	json.Unmarshal(formatByteJSON, &data1)

	// Data Ketiga Mengguakan Interface/Any
	var data2 interface{}
	json.Unmarshal(formatByteJSON, &data2)

	// Type assertion = pengecekan
	decodedData, ok := data2.(map[string]interface{})
	if !ok {
		fmt.Println("Failed to convert data2 to map[string]interface{}")
		return
	}

	fmt.Println("============================================================")

	fmt.Println("Data Pertama - Menggunakan Struct")
	fmt.Println("Name : ", data.Name)
	fmt.Println("Age : ", data.Age)

	fmt.Println()

	fmt.Println("Data Kedua - Menggunakan Map")
	fmt.Println("Name : ", data1["name"])
	fmt.Println("Age : ", data1["age"])

	fmt.Println()

	fmt.Println("Data Kedua - Menggunakan Any")
	fmt.Println("Name : ", decodedData["name"])
	fmt.Println("Age : ", decodedData["age"])
}

/* 
Kasus ini merupakan Endode (deserialization/unmarshal) karena kita mengubah struktur data Go (string)
menjadi format json untuk menampilkan disisi Client.
*/

func EmployeeControllerEncode() {
	var ObjectEmployees1 = []Employee{{ "Hamid", 20 }, { "Abdurrahman", 18 }}
	var dataJSON, _ = json.Marshal(ObjectEmployees1)

	var JSON = string(dataJSON)
	fmt.Println(JSON)
}

