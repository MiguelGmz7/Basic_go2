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

func EncodeManager(manager *Manager) (io.Reader, error) {
	err := json.Marshal(manager)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main () {
	man := &Manager{Fullname: "Rocky", Position: "Tier1", Age: 51, YearsInCompany: 5}
	EncodeManager(man)
}