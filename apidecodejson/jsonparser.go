package decodejson

import "fmt"
import "encoding/json"

//Employee Details Structure
type Employee struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Practice string `json:"practice"`
	Team     string `json:"team"`
	Manager  string `json:"manager"`
}

// JSONTOStruct to Predefine Structure
func JSONTOStruct(jsonString string) (*Employee, error) {
	var employee Employee
	error := json.Unmarshal([]byte(jsonString), &employee)
	if error == nil {
		fmt.Printf("%+v\n", employee)
	} else {
		fmt.Println(error)
	}
	return &employee, error
}

// fmt.Printf("%+v\n", employee.ID)
// fmt.Printf("%+v\n", employee.Name)
// fmt.Printf("%+v\n", employee.Address)
// fmt.Printf("%+v\n", employee.Practice)
// fmt.Printf("%+v\n", employee.Team)
// fmt.Printf("%+v\n", employee.Manager)
