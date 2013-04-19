// Copyright ©2012 Wei-Wei Guo <wwguocn@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package num provides interfaces and algorithms for numeric analysis.
package num

import (
	// "github.com/wwguo/ai.go/base"
	// "math"
	// "fmt"
)

type matrix [][]float64

// Assign a matrix by row.
//   m : rows of the matrix
//   n : columns of the matrix, if zero, square matrix of mxm
//   e : `nil' for zeor vector, same effect as []float64{0}
//       slice for recursive assignment
func Matrix(m int, n int, e []float64) (M matrix) {
	if m !=0 && n != 0 {
		l := m*n
		v := Vector(l, e)
		for i := 0; i < m; i++ {
			M = append(M, v[i*n:(i+1)*n])
		}
	} else if m != 0 && n == 0 {
		v := Vector(m, e)
		for i := 0; i < m; i++ {
			z := make([]float64, m, m)
			M = append(M, z)
			M[i][i] = v[i]
		}
	} 
	return
}

func (M matrix) Size() (r, c int) {
	r = len(M)
	c = len(M[0])
	return
}

// //======OPERATIONS======

// func (M matrix) RowSwap(i,j int) {
// 	M[i], M[j] = M[j], M[i]
// }

// func (M matrix) ColSwap(i,j int) {
// 	for r, _ := range M {
// 		M[r][i], M[r][j] = M[r][j], M[r][i]
// 	}
// }

// func (M matrix) ScalarMultiply(r int, s float64) {
// 	for i, _ := range M[r] {
// 		M[r][i] *= s
// 	}
// }



// func GaussJordan(A, B matrix) error {
// 	n := A.NRows()
// 	m := B.NCols()
// 	var irow, icol int
// 	var indxr, indxc, pivot Index
// 	for j := 0; j < n; j++ {
// 		pivot = append(pivot, 0)
// 		indxr = append(indxr, j)
// 		indxc = append(indxc, j)
// 	}
// 	for i := 0; i < n; i++ {
// 		var big float64
// 		// Find the index of the bigest element in the non-pivoting part.
// 		for j := 0; j < n; j++ {
// 			if pivot[j] != 1 {
// 				for k := 0; k < n; k++ {
// 					if pivot[k] == 0 {
// 						if math.Abs(A[j][k]) >= big {
// 							big = math.Abs(A[j][k])
// 							irow = j
// 							icol = k
// 						}
// 					}
// 				}
// 			}
// 		}
// 		pivot[icol]++
// 		if irow != icol {
// 			// interchange columns and rows. 
// 			A.RowSwap(irow, icol)
// 			B.RowSwap(irow, icol)
// 		}
// 		indxr[i] = irow
// 		indxc[i] = icol
// 		// Pay attention. The index is (icol, icol).
// 		if A[icol][icol] == 0 {
// 			return base.ExceptionSingular
// 		}
// 		pivint := 1/A[icol][icol]
// 		// A[icol][icol] = 1
// 		// pivoting columns. 
// 		A.ScalarMultiply(icol, pivint)
// 		B.ScalarMultiply(icol, pivint)
// 		for ll := 0; ll < n; ll++ {
// 			if ll != icol {
// 				dum := A[ll][icol]
// 				for l := 0; l < n; l++ {
// 					A[ll][l] -= A[icol][l]*dum
// 				}
// 				for l := 0; l < m; l++ {
// 					B[ll][l] -= B[icol][l]*dum
// 				}
// 			}
// 		}
// 	}
// 	fmt.Printf("Index: %v, %v\n", indxr, indxc)
// 	// Change row order back to return X matrix in original order. 
// 	for l := n-1; l >=0; l-- {
// 		if indxr[l] != indxc[l] {
// 			A.RowSwap(indxr[l], indxc[l])
// 			B.RowSwap(indxr[l], indxc[l])
// 		}
// 	}
// 	return nil
// }

// func GaussJordanInvert(A matrix) (matrix, error) {
// 	n := A.NRows()
// 	if !A.VerifySquare(n) {
// 		return nil, base.ErrorDimensionMismatch
// 	}
// 	B := MakeIdentity(n)
// 	// B := MakeConstant(0,n,0)
// 	err := GaussJordan(A, B)
// 	return B, err
// }

