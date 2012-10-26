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
	// "math"
	"fmt"
	"github.com/wwguo/ai.go/matrix"
)

// var A, b, c *matrix.Matrix
// A := matrix.MakeMatrix(3, 7, []float64{1, 1, 1, 1, 1, 0, 0, 7, 5, 3, 2, 0, 1, 0, 3, 5, 10, 15, 0, 0, 1})
// b := matrix.MakeMatrix(3, 1, []float64{15, 120, 100})
// c := matrix.MakeMatrix(1, 7, []float64{4, 5, 9, 11, 0, 0, 0})

func Simplex(c, A, b *matrix.Matrix) (z float64, x []float64) {
	// Check matrix rows and columns. 
	m1, n1 := A.Size() // numRows, numCols
	m2, _ := b.Size()
	_, n2 := c.Size()

	var m, n int

	if m1 == m2 && n1 == n2 {
		if m1 < n1 {
			m = m1
			n = n1
		} else {
			fmt.Printf("Restricts are more than variables. The Simplex isn't suitable for the problem.")
			return
		}			
	} else {
		fmt.Printf("Row numbers or column numbers doesn't fit.")
		return
	}

	// Prepare index list for creating matrices in each iteration.
	index := make([]int, n, n)
	for i := 1; i <= n; i++ {
		index[i-1] = i
	}
	iB := index[n-m : n]
	iN := index[0 : n-m]

	// Define variables specified in Simplex algorithm.
	var X, Z, cB, cN, B, N *matrix.Matrix
	var k, l int
	var σmax, βmin float64

	// For starting the iteration.
	k = -1

	for k != 0 {
		// Initialize matrices.
		X = matrix.MakeZero(n, 1)

		cB = c.GetCols(iB)
		cN = c.GetCols(iN)
		B  = A.GetCols(iB)
		N  = A.GetCols(iN)

		// Compute intermidiate matrices for algorithm.
		Binv, _ := B.Inverse()
		T1, _ := matrix.Product(cB, Binv, N)
		Sigma, _ := matrix.Subtract(cN, T1)
		σ := Sigma.Vector()

		// Find the biggest σ for judgement. 
		k = 0
		σmax = 0
		for i, s := range σ {
			if s > σmax {
				σmax = s
				k = i + 1
			}
		}

		if k == 0 {
			// If there is no σ positive, the iteration stop and optimization reaches.
			d, _ := matrix.Multiply(Binv, b)
			X.SetRows(iB, d)
			Z, _ = matrix.Product(cB, Binv, b)
			x = X.Vector()
			z = Z.Get(1, 1)
			continue
		} else {
			// If there is at least one σ positive, looking for varialbe to exchange.
			ak := N.GetCol(k)
			a, _ := matrix.Multiply(Binv, ak)
			d, _ := matrix.Multiply(Binv, b)

			// This is for find a positive β for later comparison.
			avec := a.Vector()
			dvec := d.Vector()
			for i, ai := range avec {
				if ai > 0 {
					βmin = dvec[i] / ai
					l = i + 1
					break
				}
			}

			// Find the smallest β which corresponding to the variable to exchange.
			for i := 0; i < m; i++ {
				β := dvec[i] / avec[i]
				if avec[i] > 0 && β < βmin {
					βmin = β
					l = i + 1
				}
			}

			// Exchange varialbes, controlled by indices.
			iB[l-1], iN[k-1] = iN[k-1], iB[l-1]
		}
	}

	return
}
