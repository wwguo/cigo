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

// Package opt provides interfaces and algorithms for optimization.
package opt

// simplex.go defines the simplex algorithm for linear programming.

import (
	"math"
	"github.com/bobhancock/gomatrix/matrix"
	// "code.google.com/p/gomatrix/matrix"
	"fmt"
)


// var A, x, b, c matrix.DenseMatrix

func SubMatrix 



func Simplex (A, b, c matrix.DenseMatrix) (x matrix.DenseMatrix) {
	m1, n1 := A.GetSize()  // numRows, numCols
	m2, _ := b.GetSize()
	_, n2 := c.GetSize()

	if m1 == m2 && n1 == n2 {
		if m < n {
			m := m1
			n := n1
		} else {
			fmt.Printf("Restricts are more than variables. The Simplex isn't suitable for the problem.")
			break
		}			
	} else {
		fmt.Printf("Row numbers or column numbers doesn't fit.")
		break
	}

	x := matrix.Zeros(n, 1)

	xB := x.GetMatrix(1, 1, m, 1)
	cB := c.GetMatrix(1, 1, 1, m)
	B  := A.GetMatrix(1, 1, m, m)
	xN := x.GetMatrix(m+1, 1, n, 1)
	cN := c.GetMatrix(1, m+1, 1, n)
	N  := A.GetMatrix(m+1, m+1, n, n)

	// Todo: 查替换矩阵的行或列（是否要用指针类型？） 
	//       GetRowVector(i int)
	//       SetRowVector(src, row int)
	//       GetColVector(j int)
	//       SetMatrix(1, j, B)
	//       Get(i,j)
	//       Set(i,j)

	Binv, _ := B.Inverse()
	xB, _ = Binv.Times(b)
	z, _ := cB.Times(xB)

	

}