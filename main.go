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