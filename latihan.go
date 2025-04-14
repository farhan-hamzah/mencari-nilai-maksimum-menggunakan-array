package main
import "fmt"

func main(){
	const NMAX int = 100
	var A[NMAX]int
	var i, n, puncak int

	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
		puncak = A[i]
	}
	for i = 0; i < n; i++{
		if puncak < A[i]{
			puncak = A[i]
		}
	}
	fmt.Print(puncak)
}
