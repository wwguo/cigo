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

// Package matrix provides interfaces and algorithms for matrix manipulation.
package base

import (
	// "math"
	"fmt"
)


type Matrix struct {
	Vector
	rows int
	cols int
}

func (M *Matrix) Size() (r, c int) {
	r = M.rows
	c = M.cols
	return
}

func (M *Matrix) Row() (r int) {
	r = M.rows
	return
}

func (M *Matrix) Col() (c int) {
	c = M.cols
	return
}

// Functions for defining matrices.

func MakeMatrix (r int, c int, E Vector) *Matrix {
	M := new(Matrix)
	if len(E) == r*c {
		M.Vector = E
		M.rows = r
		M.cols = c
	} else {
		fmt.Printf("The lenght of elements doesn't compitable with the specified matrix.")
	}
	return M
}

func MakeZero(r, c int) *Matrix {
	M := new(Matrix)
	M.Vector = make(Vector, r*c)
	M.rows = r
	M.cols = c
	return M
}

func MakeUnit(r, c int) *Matrix {
	M := new(Matrix)
	M.Vector = make(Vector, r*c)
	for i, _ := range M.Vector {
		M.Vector[i] = 1.0
	}
	M.rows = r
	M.cols = c
	return M
}

func MakeIdentity(s int) *Matrix {
	M := new(Matrix)
	M.Vector = make(Vector, s*s)
	for i := 0; i < s; i++ {
		M.Vector[s*i+i] = 1.0
	}
	M.rows = s
	M.cols = s
	return M
}

// Methods for manipulating parts of a matrix.

// Operation on single element at i,j. 
func (M *Matrix) Get(i,j int) (e float64) {
	e = M.Vector[(i-1)*M.cols+(j-1)]
	return e
}

func (M *Matrix) Set(i,j int, e float64) () {
	M.Vector[(i-1)*M.cols+(j-1)] = e
	return
}

// Operation on single column at c. 
func (M *Matrix) GetCol(c int) *Matrix {
	B := new(Matrix)
	for i := 1; i <= M.rows; i++ {
		B.Vector = append(B.Vector, M.Get(i,c))
	}
	B.rows = M.rows
	B.cols = 1
	return B
}

func (M *Matrix) SetCol(c int, A *Matrix) (err error) {
	if A.cols > 1 || M.rows != A.rows {
		err = ErrorDimensionMismatch
		return
	}
	for i := 1; i <= M.rows; i++ {
		M.Set(i, c, A.Get(i,1))
	}
	return
}

// Operation on single row at r. 
func (M *Matrix) GetRow(r int) *Matrix {
	B := new(Matrix)
	B.Vector = M.Vector[(r-1)*M.cols : r*M.cols]
	B.rows = 1
	B.cols = M.cols
	return B
}

func (M *Matrix) SetRow(r int, A *Matrix) (err error) {
	if A.rows > 1 || M.cols != A.cols {
		err = ErrorDimensionMismatch
		return
	}
	for i := 1; i <= M.cols; i++ {
		M.Set(r,i, A.Get(1,i))
	}
	return
}

// Operation on mutiple columns at column list c. 
func (M *Matrix) GetCols(c []int) *Matrix {
	B := new(Matrix)
	for i := 1; i <= M.rows; i++ {
		for _, j := range c {
			e := M.Get(i,j)
			B.Vector = append(B.Vector, e)
		}
	}
	B.rows = M.rows
	B.cols = len(c)
	return B
}

func (M *Matrix) SetCols(c []int, A *Matrix) (err error) {
	for j, col := range c {
		M.SetCol(col, A.GetCol(j+1))
	}
	return
}

// Operation on mutiple rows at row list r. 
func (M *Matrix) GetRows(r []int) *Matrix {
	B := new(Matrix)
	for _, i := range r {
		row := M.GetRow(i)
		B.Vector = append(B.Vector, row.Vector...)
	}
	B.rows = len(r)
	B.cols = M.cols
	return B
}

func (M *Matrix) SetRows(r []int, A *Matrix) (err error) {
	for i, row := range r {
		M.SetRow(row, A.GetRow(i+1))
	}
	return
}


// Get submatirx by sequences of rows and cols. 
func (M *Matrix) GetSub(r, c []int) *Matrix {
	A := M.GetRows(r)
	B := A.GetCols(c)
	B.rows = len(r)
	B.cols = len(c)
	return B
}

// Get submatrix by the upper left corner (i,j) and bottom right corner (m,n).
func (M *Matrix) GetMatrix(i, j, m, n int) *Matrix {
	var r, c []int
	for k := i; k <= m; k++ {
		r = append(r, k)
	}
	for k := j; k <= n; k++ {
		c = append(c, k)
	}
	A := M.GetRows(r)
	B := A.GetCols(c)
	B.rows = len(r)
	B.cols = len(c)
	return B
}

// Matrix computation functions

// Exchange tow rows.
func (M *Matrix) SwapRows(r1, r2 int) {
	r10 := (r1-1)*M.cols
	r20 := (r2-1)*M.cols
	for j := 0; j < M.cols; j++ {
		M.Vector[r10+j], M.Vector[r20+j] = M.Vector[r20+j], M.Vector[r10+j]
	}
}

// Exchange two cols.
func (M *Matrix) SwapCols(c1, c2 int) {
	for i := 0; i < M.rows; i++ {
		M.Vector[i*M.cols+c1], M.Vector[i*M.cols+c2] = M.Vector[i*M.cols+c2], M.Vector[i*M.cols+c1]
	}
}

// Combine two matrix and get a new matrix [A B].
func Combine(order bool, A, B *Matrix) (M *Matrix, err error) {
	if order {
		if A.rows != B.rows {
			err = ErrorDimensionMismatch
			return
		}
		M = MakeZero(A.rows, A.cols + B.cols)
		for i := 1; i<=A.cols; i++ {
			M.SetCol(i, A.GetCol(i))
		}
		for i := 1; i<=B.cols; i++ {
			M.SetCol(A.cols+i, B.GetCol(i))
		}
	} else {
		if A.cols != B.cols {
			err = ErrorDimensionMismatch
			return
		}
		M = MakeZero(A.rows+B.rows, A.cols)
		for i := 1; i<=A.rows; i++ {
			M.SetRow(i, A.GetRow(i))
		}
		for i := 1; i<=B.rows; i++ {
			M.SetRow(A.rows+i, B.GetRow(i))
		}
	}
	return
}

// Matrix inversion.
func (M *Matrix) Inverse() (*Matrix, error) {
	if M.rows != M.cols {
		return nil, ErrorDimensionMismatch
	}
	s := M.rows
	I := MakeIdentity(s)
	A, _ := Combine(true, M, I)
	// Transforing first matrix to identity matrix one line each time.
	for i := 1; i <= s; i++ {
		// // Looking for the largest absolute value of the corresponding column. 
		// j := i
		// for k := i; k <= s; k++ {
		// 	if math.Abs(A.Get(k, i)) > math.Abs(A.Get(j, i)) {
		// 		j = k
		// 	}
		// }
		// // Moving the row with the largest value of the column to current top.
		// if j != i {
		// 	A.SwapRows(i, j)
		// }

		// If the largest value is 0, the matrix is not inversable.
		rmax := A.Get(i, i)
		if rmax == 0 {
			return nil, ExceptionSingular
		}
		// Transforing the diagnal value of the row to 1. 
		for t := 0; t < 2*s; t++ {
			A.Vector[2*s*(i-1)+t] *= 1.0/rmax
		}
		// Transforing other values in the column to 0.
		for l := 1; l <= s; l++ {
			if l == i {
				continue
			}
			// Find the value to be subtracted out the line. 
			cval := A.Vector[2*s*(l-1)+(i-1)]
			// Each value of the line subtracts the corresponding value in i line multipling cval.
			for c := 0; c < 2*s; c++ {
				A.Vector[2*s*(l-1)+c] -= A.Vector[2*s*(i-1)+c]*cval
			}
		}
	}
	Minv := A.GetMatrix(1, s+1, s, s*2)
	// Minv := A.GetMatrix(1, s+1, s, s)
	return Minv, nil
}

// Matrix addition.
func Add(A, B *Matrix) (M *Matrix, err error) {
	if A.rows != B.rows || A.cols != B.cols {
		return nil, ErrorDimensionMismatch
	}
	M = MakeZero(A.rows, A.cols)
	for i := 1; i <= A.rows; i++ {
		for j := 1; j <= A.cols; j++ {
			M.Set(i, j, A.Get(i,j)+B.Get(i, j))
		}
	}
	return M, nil
}

// Matrix subtract (actually addition).
func Subtract(A, B *Matrix) (M *Matrix, err error) {
	if A.rows != B.rows || A.cols != B.cols {
		return nil, ErrorDimensionMismatch
	}
	M = MakeZero(A.rows, A.cols)
	for i := 1; i <= A.rows; i++ {
		for j := 1; j <= A.cols; j++ {
			M.Set(i, j, A.Get(i,j)-B.Get(i, j))
		}
	}
	return M, nil
}

// Matrix multiplication, element by element.
func Times(A, B *Matrix) (M *Matrix, err error) {
	if A.rows != B.rows || A.cols != B.cols {
		return nil, ErrorDimensionMismatch
	}
	M = MakeZero(A.rows, A.cols)
	for i := 1; i <= A.rows; i++ {
		for j := 1; j <= A.cols; j++ {
			M.Set(i, j, A.Get(i,j)*B.Get(i, j))
		}
	}
	return M, nil
}

// Matrix element devidient.
func Divide(A, B *Matrix) (M *Matrix, err error) {
	if A.rows != B.rows || A.cols != B.cols {
		return nil, ErrorDimensionMismatch
	}
	M = MakeZero(A.rows, A.cols)
	for i := 1; i <= A.rows; i++ {
		for j := 1; j <= A.cols; j++ {
			M.Set(i, j, A.Get(i,j)/B.Get(i, j))
		}
	}
	return M, nil
}

// Matrix multiplication, the common one.
func Multiply(A, B *Matrix) (M *Matrix, err error) {
	if A.cols != B.rows {
		return nil, ErrorDimensionMismatch
	}
	M = MakeZero(A.rows, B.cols)
	// Use Strassen's algorithm for square matrix multiplication.
	if A.cols == A.rows && B.cols == B.rows {
		M = squareMatrixMultiply(A, B)
	} else {
		for i := 1; i <= A.rows; i++ {
			for j := 1; j <= B.cols; j++ {
				sum := float64(0)
				for k := 1; k <= A.cols; k++ {
					sum += A.Get(i,k) * B.Get(k, j)
				}
				M.Set(i,j,sum)
			}
		}
	}
	return M, nil
}

func squareMatrixMultiply (A, B *Matrix) (M *Matrix) {
	n := A.rows
	M = MakeZero(n,n)
	if n == 1 {
		M.Set(1, 1, A.Get(1,1) * B.Get(1,1))
	} else {
		A11 := A.GetMatrix(1,     1,     n/2,   n/2 )
		A12 := A.GetMatrix(1,     n/2+1, n/2,   n   )
		A21 := A.GetMatrix(n/2+1, 1,     n,     n/2 )
		A22 := A.GetMatrix(n/2+1, n/2+1, n,     n   )
		B11 := B.GetMatrix(1,     1,     n/2,   n/2 )
		B12 := B.GetMatrix(1,     n/2+1, n/2,   n   )
		B21 := B.GetMatrix(n/2+1, 1,     n,     n/2 )
		B22 := B.GetMatrix(n/2+1, n/2+1, n,     n   )
		M11, _ := Add(squareMatrixMultiply(A11, B11), squareMatrixMultiply(A12, B21))
		M12, _ := Add(squareMatrixMultiply(A11, B12), squareMatrixMultiply(A12, B22))
		M21, _ := Add(squareMatrixMultiply(A21, B11), squareMatrixMultiply(A22, B21))
		M22, _ := Add(squareMatrixMultiply(A21, B12), squareMatrixMultiply(A22, B22))
		M1, _ := Combine(true, M11, M12)
		M2, _ := Combine(true, M21, M22)
		M, _ = Combine(false, M1, M2)
	}
	return
}

// Suger functions.

// Production of multiple matrices.
func Product(Mlist ...*Matrix) (M *Matrix, err error) {
	if len(Mlist) < 2 {
		err = ErrorNotEnoughMatrix
		return
	}
	A := Mlist[0]
	for _, B := range Mlist[1:] {
		A, err = Multiply(A, B)
	}
	M = A
	return
}

// Tequality of the two matrices
func EqualMatrix (A, B *Matrix) bool {
	if A.rows != B.rows || A.cols != B.cols {
		return false
	}
	for i := 1; i <= A.rows; i++ {
		for j := 1; j <= A.cols; j++ {
			if A.Get(i, j) != B.Get(i, j) {
				return false
			}
		}
	}
	return true
}
