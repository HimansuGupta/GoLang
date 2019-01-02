package databaseconnection

import (
	"fmt"
	"log"

	decodejson "github.com/himansu/restapi/apidecodejson"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//EmployeeDB for contain Server,DB and Collection name
type EmployeeDB struct {
	Server     string
	Database   string
	Collection string
}

var empdb EmployeeDB

// GetConnection database name and collection name
func GetConnection() {
	empdb = EmployeeDB{
		Server:     "localhost",
		Database:   "employeedb",
		Collection: "employeedetails",
	}
	Connect(&empdb)
}

//mgo.Connect.Database
var db *mgo.Database

// Connect to database using EmployeeDB Stuct
func Connect(empdbstruct *EmployeeDB) {
	session, err := mgo.Dial(empdbstruct.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(empdbstruct.Database)
	fmt.Println(db.Name)
	fmt.Println(db.CollectionNames())
}

// FindAllEmployees details
func FindAllEmployees() ([]decodejson.Employee, error) {
	fmt.Println("Start Find All")
	var employees []decodejson.Employee
	error := db.C(empdb.Collection).Find(bson.M{}).All(&employees)
	fmt.Println("After get data ", employees)
	return employees, error
}

// FindByID details
func FindByID(id string) (decodejson.Employee, error) {
	var employee decodejson.Employee
	err := db.C(empdb.Collection).Find(bson.M{"id": id}).One(&employee)
	fmt.Println("in side FindBYID", employee)
	return employee, err
}

//AddEmployee new Employee in employeedb
func AddEmployee(employee decodejson.Employee) (decodejson.Employee, error) {
	err := db.C(empdb.Collection).Insert(&employee)
	return employee, err
}

//UpdateByID Employee can edit/update his detail
func UpdateByID(employee decodejson.Employee) (decodejson.Employee, error) {
	err := db.C(empdb.Collection).Update(bson.M{"id": employee.ID}, &employee)
	return employee, err
}

// DeleteByID employee details
func DeleteByID(id string) error {
	err := db.C(empdb.Collection).Remove(bson.M{"id": id})
	return err
}

// func main() {

// 	fmt.Println(empdb)
// 	Connect(&empdb)
// 	FindAllEmployees()
// 	FindByID("GS-2080")

// 	//employee create for test Add Employee
// 	// employee := decodejson.Employee{
// 	// 	"GS-2081",
// 	// 	"Rakesh Dave",
// 	// 	"Viman nagar",
// 	// 	"IBM",
// 	// 	"Cloud-Orchestration",
// 	// 	"Gautam",
// 	// }

// 	// AddEmployee(employee)

// 	//UpdateByID(employee)

// 	DeleteByID("GS-2081")
// }
