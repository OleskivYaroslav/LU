package arr

import (
    "fmt"
	"os"
	"encoding/json"
	"io/ioutil"
)

type Vector struct {
	Vec []float64 `json:"Vec"`
}

func (v *Vector) Read() {
}

func (v Vector) Write() {
	fmt.Println(v)
}

func (v *Vector) ReadJson() {
	file, err := os.Open("vect.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &v)
}

func (v *Vector) Create(n int64) {
	v.Vec = make([]float64, n)
}