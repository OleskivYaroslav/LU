package arr

import (
    "fmt"
	"os"
	"encoding/json"
	"io/ioutil"
)

type Matrix struct {
	MTRX []Vector `json:"MTRX"`
}

func (m *Matrix) Read() {
	
}

func (m Matrix) Write() {
	fmt.Println(m)
}

func (m *Matrix) ReadJson() {
	file, err := os.Open("mtrx.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &m)
}

func (m *Matrix) Create(n int64) {
	
	m.MTRX = make([]Vector, n)
	for i := range m.MTRX{
		var v Vector
		v.Create(n)
		m.MTRX[i] = v
	}
}