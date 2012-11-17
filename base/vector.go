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

// Package base provides basic interfaces and algorithms.
package base

import (
	"math"
	// "fmt"
)

// vector.go defines algorithms for slice vector manipulations.

type Vector []float64


// Sorting vector into non-decreasing or non-increasingorder with Insertion-Sort.
func (A Vector) InsertionSort (order bool) {
	if order {
		for j, key := range A {
			i := j - 1
			for i >= 0 && A[i] > key {
				A[i+1] = A[i]
				i = i - 1
			}
			A[i+1] = key
		}
	} else {
		for j, key := range A {
			i := j - 1
			for i >= 0 && A[i] < key {
				A[i+1] = A[i]
				i = i - 1
			}
			A[i+1] = key
		}

	}
}

// Sorting vector into non-decreasing or non-increasingorder with Merge-Sort.
func (A Vector) MergeSort (p int, r int, order bool) {
	if p < r - 1 {
		q := (p+r)/2
		A.MergeSort(p,q, order)
		A.MergeSort(q,r, order)
		A.merge(order, p,q,r)
	}
}

func (A Vector) merge (order bool, p, q, r int) {
    L := make(Vector, q-p)
    R := make(Vector, r-q)
    _ = copy(L, A[p:q])
    _ = copy(R, A[q:r])
	var i,j int
	if order {
		L = append(L, math.Inf(1))
		R = append(R, math.Inf(1))
		for k := p; k < r; k++ {
			if L[i] <= R[j] {
				A[k] = L[i]
				i++
			} else {
				A[k] = R[j]
				j++
			}
		}
	} else {
		L = append(L, math.Inf(-1))
		R = append(R, math.Inf(-1))
		for k := p; k < r; k++ {
			if L[i] >= R[j] {
				A[k] = L[i]
				i++
			} else {
				A[k] = R[j]
				j++
			}
		}
	}
}







// Suger functions

// Check two vectors' equality.
func EqualVector (a, b Vector) bool {
	if len(a) != len(b) {
		return false
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
	}
	return true
}