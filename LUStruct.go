package main

import (
	
)

func LUStruct(A Matrix, b Vector)[]float64{
	n := len(b.Vec) 
	mtr := make([][]float64,n)
	for i, v := range A.MTRX{
		mtr[i] = v.Vec
	}
	f := make([]float64,n)
	f = b.Vec
	 
	return LU(mtr, f, n)
}

func GausStruct(A Matrix, b Vector)[]float64{
	n := len(b.Vec) 
	mtr := make([][]float64,n)
	for i, v := range A.MTRX{
		mtr[i] = v.Vec
	}
	f := make([]float64,n)
	f = b.Vec
	 
	return Gaus(mtr, f, n)
}
