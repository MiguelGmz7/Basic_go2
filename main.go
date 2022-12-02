
package main 

import ( 
	"fmt"
	"encoding/json"
)

type Manager struct {
	Fullname string
	Position string
	Age int32
	YearsInCompany int32
}

func EncodeManager(manager *Manager) {
	e, err := json.Marshal(manager)
	if e, err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))
}
func main () {
	man := &Manager{Fullname: "Rocky", Position: "Tier1", Age: 51, YearsInCompany: 5}
    EncodeManager(man)
}

/*
package main

import (
    "fmt"
    "encoding/json"
)

type Employee struct {
    Name string `json:"empname"`
	Number int  `json:"empid"`
}

func main() {
    emp := &Employee{Name: "Rocky",Number: 5454}
    e, err := json.Marshal(emp)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(e))
}
*/