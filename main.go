package main

// Package controller Dash Testrail API
//
// The purpose of this application is to store and retrieve Employee records
//
//
//
//     BasePath: /api/testrail/v1
//     Version: 0.0.1
//     License: Copy Right © 2019 GG Lab Private Limited

//		All rights reserved.
//		This material is confidential and proprietary to GS Lab Private Limited.
//		No part of this material should be reproduced, published in any form by any means,
//		electronic or mechanical including photocopy or any information storage or
//		retrieval system nor shougo:generate swagger generate spec -m -o ./swagger.json
//ld the material be disclosed to third parties without
//		the express written authorization of GS Lab GS LabPrivate Limited.Licensed Materials -

//
//     Contact: Himansu Gupta <himansu.gupta@gslab.com>
//
//     Consumes:
//       - application/json
//
//     Produces:
//       - application/json
//
//     Security:
//       - token:
//
//     SecurityDefinitions:
//       token:
//         type: apiKey
//         in: header
//         name: Authorization
//
//
// 		swagger:meta

//		go:generate swagger generate spec -m -o ./swagger.json

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	dbconnection "github.com/himansu/restapi/apidatabaseconnection"
	decodejson "github.com/himansu/restapi/apidecodejson"
)

// GetEmployees from db
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET Employees  Implement")
	dbconnection.GetConnection()
	var employees []decodejson.Employee
	employees, error := dbconnection.FindAllEmployees()
	if error == nil {
		json.NewEncoder(w).Encode(&employees)
	} else {
		fmt.Println("Error :- ", error)
	}

}

// GetEmployeeByID from db
func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET Employees  Implement")
	var employee decodejson.Employee
	dbconnection.GetConnection()
	params := mux.Vars(r)
	fmt.Println(len(params))
	fmt.Println(params["id"])
	employee, error := dbconnection.FindByID(params["id"])
	if error == nil {
		json.NewEncoder(w).Encode(&employee)
	} else {
		fmt.Println("Error :- ", error)
	}
}

// AddNewEmployee from db
func AddNewEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POSt -- Add Employees Implement")
	var employee decodejson.Employee
	error := json.NewDecoder(r.Body).Decode(&employee)
	if error == nil {
		dbconnection.GetConnection()
		employee, error := dbconnection.AddEmployee(employee)
		if error == nil {
			fmt.Println("Employee Added successfuly ", employee)
		} else {
			fmt.Println("Error :- ", error)
		}
		json.NewEncoder(w).Encode(employee)
	}
}

//UpdateEmployeeByID update Employe
func UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POSt -- Update Employees Implement")
	var employee decodejson.Employee
	error := json.NewDecoder(r.Body).Decode(&employee)
	if error == nil {
		dbconnection.GetConnection()
		employee, error := dbconnection.UpdateByID(employee)
		if error == nil {
			fmt.Println("Employee Updated successfuly ", employee)
		} else {
			fmt.Println("Error :- ", error)
		}
		json.NewEncoder(w).Encode(employee)
	}
}

// DeleteEmployee based on ID
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET Employees  Implement")
	var employee decodejson.Employee
	dbconnection.GetConnection()
	params := mux.Vars(r)
	fmt.Println(len(params))
	fmt.Println(params["id"])
	error := dbconnection.DeleteByID(params["id"])
	if error == nil {
		json.NewEncoder(w).Encode(&employee)
	} else {
		fmt.Println("Error :- ", error)
	}
}

//main entry define all api uri and methods
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employees", GetEmployees).Methods("GET")
	router.HandleFunc("/employee", AddNewEmployee).Methods("POST")
	router.HandleFunc("/employee/{id}", GetEmployeeByID).Methods("GET")
	router.HandleFunc("/updateemployee", UpdateEmployeeByID).Methods("POST")
	router.HandleFunc("/employee/{id}", DeleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
