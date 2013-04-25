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
	"github.com/wwguo/ai.go/base"
	"math"
	// "fmt"
	// "reflect"
)

// matrix.go defines algorithms for matrix manipulations.

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
			// Create new slice is necessory, or will cause line append error. 
			a := make([]float64, n, n)
			copy(a, v[i*n:(i+1)*n])
			M = append(M, a)
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

//======OPERATIONS======

func (M matrix) RowSwap(i,j int) {
	M[i], M[j] = M[j], M[i]
}

func (M matrix) ColSwap(i,j int) {
	for r, _ := range M {
		M[r][i], M[r][j] = M[r][j], M[r][i]
	}
}

func (M matrix) ColAbsMax(c int) (r int, max float64) {
	n, _ := M.Size()
	for i:=0;i<n;i++ {
		if math.Abs(M[i][c]) > max {
			max = math.Abs(M[i][c])
			r = i
		}
	}
	return
}

func (M matrix) RowScalTimes(r int, s float64) {
	for i, _ := range M[r] {
		M[r][i] *= s
	}
}

func (M matrix) ColScalTimes(c int, s float64) {
	n, _ := M.Size()
	for i:=0;i<n;i++ {
		M[i][c] *= s
	}
}

//------row and column: getter and setter-------

// Get submatrix by the upper left corner (i,j) and bottom right corner (m,n).
func (M matrix) GetMatrix(i, j, m, n int) (C matrix) {
	for r:=i-1;r<m;r++ {
		C = append(C, M[r][j-1:n])
	}
	return
}

// Combine two matrix and get a new matrix [A B].
func CombineMatrix(order string, A, B matrix) (M matrix, err error) {
	ma, na := A.Size()
	mb, nb := B.Size()
	if order == "r" {
		if ma != mb {
			err = base.ErrorDimensionMismatch
			return nil, err
		}
		for i:=0; i<ma; i++ {
			M = append(M, append(A[i], B[i]...))
		}
	} else if order == "c" {
		if na != nb {
			err = base.ErrorDimensionMismatch
			return nil, err
		}
		M = append(A, B...)
	} else {
		err = base.ErrorNilMatrix
	    return nil, err
	}
	return M, nil
}

//------Matrix operations-----------

// Matrix Addition.
func MatrixAdd(A, B matrix) (M matrix, err error) {
	m, n := A.Size()
	mb, nb := B.Size()
	if m != mb || n != nb {
		return nil, base.ErrorDimensionMismatch
	}
	M = Matrix(m,n,nil)
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			M[i][j] = A[i][j] + B[i][j]
		}
	}
	return M, nil
}

// Matrix Multiplication
func MatrixMultiply(A, B matrix) (M matrix, err error) {
	ma, na := A.Size()
	mb, nb := B.Size()
	if na != mb {
		return nil, base.ErrorDimensionMismatch
	}
	//M = Matrix(ma, nb, nil)
	if ma == na && mb == nb {
		M = squareMatrixMultiply(A, B)
	} else {
		M = Matrix(ma,nb,nil)
		for i := 0; i < ma; i++ {
			for j := 0; j < nb; j++ {
				for k := 0; k < na; k++ {
					M[i][j] += A[i][k] * B[k][j]
				}
			}
		}
	}
	return M, nil
}

// Strassen's algorithm for square matrix multiplication.
func squareMatrixMultiply(A, B matrix) (M matrix) {
	n,_ := A.Size()
	M = Matrix(n,n,nil)
	if n == 1 {
		M[0][0]= A[0][0] * B[0][0]
	} else {
		A11 := A.GetMatrix(1,     1,     n/2,   n/2 )
		A12 := A.GetMatrix(1,     n/2+1, n/2,   n   )
		A21 := A.GetMatrix(n/2+1, 1,     n,     n/2 )
		A22 := A.GetMatrix(n/2+1, n/2+1, n,     n   )
		B11 := B.GetMatrix(1,     1,     n/2,   n/2 )
		B12 := B.GetMatrix(1,     n/2+1, n/2,   n   )
		B21 := B.GetMatrix(n/2+1, 1,     n,     n/2 )
		B22 := B.GetMatrix(n/2+1, n/2+1, n,     n   )
		M11, _ := MatrixAdd(squareMatrixMultiply(A11, B11), squareMatrixMultiply(A12, B21))
		M12, _ := MatrixAdd(squareMatrixMultiply(A11, B12), squareMatrixMultiply(A12, B22))
		M21, _ := MatrixAdd(squareMatrixMultiply(A21, B11), squareMatrixMultiply(A22, B21))
		M22, _ := MatrixAdd(squareMatrixMultiply(A21, B12), squareMatrixMultiply(A22, B22))
		M1, _ := CombineMatrix("r", M11, M12)
		M2, _ := CombineMatrix("r", M21, M22)
		M, _ = CombineMatrix("c", M1, M2)
	}
	return
}

