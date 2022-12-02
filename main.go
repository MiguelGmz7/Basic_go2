
package main

import (
    "fmt"
    "encoding/json"
)

type Manager struct {
	FullName string 'json:"full_name"'
	Position string 'json: "position"'
	Age int32 'json: "age"'
	YearsInCompany 'json: "years_in_company"'
}

func main() {
    emp := &Manager{FullName: "Jack Oliver", Position: "CEO", Age: "44", YearsInCompany: "8"}
    e, err := json.Marshal(emp)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(e))
}
