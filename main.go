package main

import (
    "fmt"
	"net/http"
	"html/template"
	"strconv"
	"os"
	"encoding/json"
	"io/ioutil"
)

type Arr interface {
	ReadJson()
	Read()
	Write()
	Create(int64)
}

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

type LUStr struct {
	M Matrix
	V Vector
	X []float64
}
	const n = 3
	var k Matrix
	var v Vector
	var x []float64
	var lustr LUStr

func indexHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseGlob("templates/*.html"))
	
	t.ExecuteTemplate(w, "index.html", lustr)
}

func calculate(w http.ResponseWriter, r *http.Request) {
	
	for i:=0; i<n; i++{
		for j:=0; j<n; j++{
			s:=r.FormValue("mtrx"+strconv.Itoa(i)+strconv.Itoa(j))
			lustr.M.MTRX[i].Vec[j], _ = strconv.ParseFloat(s, 64)
		}
		q:=r.FormValue("vec"+strconv.Itoa(i))
		lustr.V.Vec[i],_=strconv.ParseFloat(q,64)
	}
	lustr.X = LUStruct(lustr.M, lustr.V)
	
	t := template.Must(template.ParseGlob("templates/*.html"))
	t.ExecuteTemplate(w, "index.html", lustr)
}

func clear(w http.ResponseWriter, r *http.Request) {
	
	k.Create(n)
	v.Create(n)
	x = make([]float64,n)
	lustr = LUStr{k,v,x}
	
	t := template.Must(template.ParseGlob("templates/*.html"))
	t.ExecuteTemplate(w, "index.html", lustr)
}



func main() {
	k.Create(n)
	v.Create(n)
	x = make([]float64,n)
	lustr = LUStr{k,v,x}
	
	
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/Calculate", calculate)
	http.HandleFunc("/Clear", clear)
	
	fmt.Println("Listening on port :3000")
	http.ListenAndServe(":3000", nil)
		
}


