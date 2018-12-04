package lupac

import (
	"sync"
)

func sum(L,U [][]float64, nn,i,j int) float64 {
	if i == nn{
		return 0.0
	}else{
		return ((L[j][i] * U[i][nn]) + sum(L, U, nn, i + 1, j))
	}
}

func sum1(F [][]float64, x []float64, nn,i int) float64 {
	if i == nn {
		return 0.0
	}else{
		return ((x[i] * F[nn][i]) + sum1(F, x, nn, i + 1))
	}
}


func sum2(F [][]float64, x []float64, nn,i,j int) float64 {
	if i == nn {
		return 0.0
	}else {
		return ((x[i] * F[j][i]) + sum2(F, x, nn, i + 1, j))
	}
}

func LU(A [][]float64, b []float64, n int)[]float64{
   	var wg sync.WaitGroup
	L := make([][]float64,n)
	U := make([][]float64,n)
	x := make([]float64,n)
	y := make([]float64,n)
    for i := 0; i < n; i++{
		L[i] = make([]float64,n)
		U[i] = make([]float64,n)
		for j := 0; j < n; j++{
			U[i][j] = 0
            L[i][j] = 0
            if i == j {U[i][j] = 1}
        }
	}
	
	for i := 0; i < n; i++{
		L[i][i] = A[i][i] - sum(L, U, i, 0, i)
			for j := i+1; j < n; j++{
				sum1:=0.0
				sum2:=0.0
				wg.Add(1)
				go func(i,j int){
					for s:=0; s<=i; s++{
						sum1+=(L[j][s])*(U[s][i])
					}
					defer wg.Done()
				}(i,j)
				wg.Add(1)
				go func(i,j int){
					for s:=0;s<=j;s++{
						sum2+=L[i][s]*U[s][j]
					}
					defer wg.Done()
				}(i,j)
				wg.Wait()
				L[j][i] = A[j][i] - sum1
				U[i][j] = (A[i][j] - sum2) / L[i][i]
			}
    }
	
	for i := 0; i < n; i++{
		y[i] = (b[i] - sum1(L, y, i, 0)) / L[i][i];
    }

	
	for i := n - 1; i >= 0; i--{
		h := sum2(U, x, n, i + 1, i);
        x[i] = (y[i] - h) / U[i][i];
    
	}
	
	return x
}

func Gaus(A [][]float64, B []float64, n int)[]float64{
	a := make([][]float64,n)
	b := make([]float64,n)
	copy(a, A)
	copy(b, B)
	x := make([]float64,n)
    for k := 0; k < n - 1; k++{
		for i := k + 1; i < n; i++{
			m := -a[i][k] / a[k][k]
			b[i] += m * b[k]
            for j := k+1; j < n; j++{
				a[i][j] += m * a[k][j]
			}
        }
	}
    x[n-1]=b[n-1]/a[n-1][n-1]
    for k := n - 2; k >= 0; k--{
		sum := 0.0
        for j := k + 1; j < n; j++{
			sum += a[k][j] * x[j]
		}
        x[k] = (b[k] - sum) / a[k][k]
    }
	return x
}




func LU001(A [][]float64, b []float64, n int)[]float64{
	L := make([][]float64,n)
	U := make([][]float64,n)
	x := make([]float64,n)
	y := make([]float64,n)
    for i := 0; i < n; i++{
		L[i] = make([]float64,n)
		U[i] = make([]float64,n)
		for j := 0; j < n; j++{
			U[i][j] = 0
            L[i][j] = 0
            if i == j {U[i][j] = 1}
        }
	}
	
	for i := 0; i < n; i++{
		L[i][i] = A[i][i] - sum(L, U, i, 0, i)
			for j := i+1; j < n; j++{
				
				
				L[j][i] = A[j][i] - sum(L, U, i, 0, j)
				U[i][j] = (A[i][j] - sum(L, U, j, 0, i)) / L[i][i]
			}
					
    }
		for i := 0; i < n; i++{
			y[i] = (b[i] - sum1(L, y, i, 0)) / L[i][i];
        }
    
	
		for i := n - 1; i >= 0; i--{
			h := sum2(U, x, n, i + 1, i);
            x[i] = (y[i] - h) / U[i][i];
        
		}
		
	return x
}
