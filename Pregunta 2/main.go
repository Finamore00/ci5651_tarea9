package main

import (
	"fmt"
	"math/rand"
)

/*
Implementación del algoritmo de verificación de Freivalds para
productos de matrices. Se asume que las matrices A, B y C tienen
las dimensiones correctas.
*/
func freivaldsMatrixMult(A [][]int, B [][]int, C [][]int) bool {
	n := len(A[0])
	x := make([]int, n)

	for i := 0; i < n; i += 1 {
		x[i] = rand.Int() % 2
	}

	bx := make([]int, n)
	cx := make([]int, n)
	for i := 0; i < n; i += 1 {
		for j := 0; j < n; j += 1 {
			bx[i] += B[i][j] * x[j]
			cx[i] += C[i][j] * x[j]
		}
	}

	abx := make([]int, n)
	for i := 0; i < n; i += 1 {
		for j := 0; j < n; j += 1 {
			abx[i] += A[i][j] * bx[j]
		}
	}

	for i := 0; i < n; i += 1 {
		if abx[i]-cx[i] != 0 {
			return false
		}
	}

	return true
}

/*
Función que dadas dos matrices A y B, indica si B es la inversa de A utilizando
el algoritmo de Freivalds. Se asume que las dimensiones de A y B son correctas
*/
func determineIfInverse(A [][]int, B [][]int, epsilon float64) bool {
	k := 1
	tmp := 1.0
	for tmp > epsilon {
		k += 1
		tmp /= 2.0
	}
	n := len(A[0])
	I := make([][]int, n)
	for i := 0; i < n; i += 1 {
		I[i] = make([]int, n)
		I[i][i] = 1
	}

	resAccum := true
	for i := 0; i < k; i += 1 {
		resAccum = resAccum && freivaldsMatrixMult(A, B, I)
	}
	return resAccum
}

func main() {
	A := [][]int{
		{
			2, 3, 2,
		},
		{
			4, 2, 3,
		},
		{
			9, 6, 7,
		},
	}

	inverseA := [][]int{
		{
			-4, -9, 5,
		},
		{
			-1, -4, 2,
		},
		{
			6, 15, -8,
		},
	}

	notInverseA := [][]int{
		{
			-4, -9, 5,
		},
		{
			-1, -5, 2,
		},
		{
			6, 15, -8,
		},
	}

	fmt.Println(determineIfInverse(A, inverseA, 0.00390625))
	fmt.Println(determineIfInverse(A, notInverseA, 0.00390625))
}
