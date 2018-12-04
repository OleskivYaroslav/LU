package lupac

import (
	"../arr"
)

func LUStruct(A arr.Matrix, b arr.Vector)[]float64{
	n := len(b.Vec) 
	mtr := make([][]float64,n)
	for i, v := range A.MTRX{
		mtr[i] = v.Vec
	}
	f := make([]float64,n)
	f = b.Vec
	 
	return LU(mtr, f, n)
}

func GausStruct(A arr.Matrix, b arr.Vector)[]float64{
	n := len(b.Vec) 
	mtr := make([][]float64,n)
	for i, v := range A.MTRX{
		mtr[i] = v.Vec
	}
	f := make([]float64,n)
	f = b.Vec
	 
	return Gaus(mtr, f, n)
}