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
)

// vector.go defines algorithms for slice vector manipulations.

// TODO: Change Vector to "map[int]float64" or add another set of Vector.

type Vector []float64

type Index  []int


func (A Vector) Extrame(order bool) float64 {
	ext := A[0]
	if order {
		for _, val := range A {
			if ext < val {
				ext = val
			}
		}
	} else {
		for _, val := range A {
			if ext > val {
				ext = val
			}
		}
	}
	return ext
}

func (A Vector) Select(p,r,i int) (v float64) {
	if p == r {
		v = A[p-1]
	}
	if p < r{
		q := A.randomPartition(p, r)
		k := q - p + 1
		switch {
		case i == k:
			v = A[q-1]
		case i < k:
			v = A.Select(p, q-1, i)
		case i > k:
			v = A.Select(q+1, r, i-k)
		}
	}
	return
}


// Sorting methods

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

// Sorting vector into non-decreasing or non-increasingorder with QuickSort.
// TODO: Add another sorting order.
func (A Vector) QuickSort(p, r int) {
	if p < r {
		q := A.partition(p,r)
		A.QuickSort(p,q-1)
		A.QuickSort(q+1,r)
	}
}

func (A Vector) partition(p, r int) int {
	x := A[r-1]
	i := p - 1
	for j := p; j <= r-1; j++ {
		if A[j-1] <= x {
			i++
			A[i-1], A[j-1] = A[j-1], A[i-1]
		}
	}
	A[i], A[r-1] = A[r-1], A[i]
	return i + 1
}

func (A Vector) randomPartition(p, r int) int {
	i := random(p, r)
	A[r-1], A[i-1] = A[i-1], A[r-1]
	return A.partition(p,r)
}

func (A Vector) RandomQuickSort(p, r int) {
	if p < r {
		q := A.randomPartition(p,r)
		A.RandomQuickSort(p,q-1)
		A.RandomQuickSort(q+1,r)
	}
}

// Sorting vector into non-decreasing or non-increasingorder with HeapSort.
func (A Vector) HeapSort (order bool) {
	H := BuildOrderHeap(A, order)
	for i := H.size; i >= 2; i-- {
		H.Vector[0], H.Vector[i-1] = H.Vector[i-1], H.Vector[0]
		H.size = H.size - 1
		H.OrderHeapify(1, order)
	}
}

// Sorting vector into non-decreasing with BucketSort.
func (A Vector) BucketSort () {
	var B [][]float64
	n := len(A)
	for i := 0; i < n; i++ {
		B = append(B, nil)
	}
	for j := 0; j < len(A); j++ {
		intpart, _ := math.Modf(A[j]*float64(n))
		index := int(intpart)
		B[index] = append(B[index], A[j])
	}
	for i := 0; i < n; i++ {
		for j, key := range B[i] {
			k := j - 1
			for k >= 0 && B[i][k] > key {
				B[i][k+1] = B[i][k]
				k = k - 1
			}
			B[i][k+1] = key
		}
	}
	var C []float64
	for i := 0; i < n; i++ {
		C = append(C, B[i]...)
	}
	for i := 0; i < n; i++ {
		A[i] = C[i]
	}
}

// Sorting int vector into non-decreasing or non-increasingorder with CountingSort.
func (I Index) CountingSort (k int) {
	var B, C []int
	for i := 0; i <= k; i++ {
		C = append(C, 0)
	}
	for j := 0; j < len(I); j++ {
		C[I[j]]++
		B = append(B, 0)
	}
	for i := 1; i <= k; i++ {
		C[i] = C[i] + C[i-1]
	}
	for j := len(I)-1; j >= 0; j-- {
		B[C[I[j]]-1] = I[j]
		C[I[j]]--
	}
	for j := 0; j < len(I); j++ {
		I[j] = B[j]
	}
}

// Sorting int vector into non-decreasing or non-increasingorder with RadixSort.
func (I Index) RadixSort (d int) {
	var L, C, B []int
	for _, val := range I {
		L = append(L, IntDigit(val,d))
	}
	for k := 0; k <= 9; k++ {
		C = append(C, 0)
	}
	for j := 0; j < len(I); j++ {
		C[L[j]]++
		B = append(B, 0)
	}
	for k := 1; k <= 9; k++ {
		C[k] = C[k] + C[k-1]
	}
	for j := len(I)-1; j >= 0; j-- {
		B[C[L[j]]-1] = I[j]
		C[L[j]]--
	}
	for j := 0; j < len(I); j++ {
		I[j] = B[j]
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
