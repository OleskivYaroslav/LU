package main

import (
    "fmt"
	"./lupac"
	"./arr"
	"net/http"
	"html/template"
	"strconv"
)

type LUStr struct {
	M arr.Matrix
	V arr.Vector
	X []float64
}
	const n = 3
	var k arr.Matrix
	var v arr.Vector
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
	lustr.X = lupac.LUStruct(lustr.M, lustr.V)
	
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
