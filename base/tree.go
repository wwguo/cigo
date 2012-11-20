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


type Heap struct {
	Vector
	number []int
	size   int
}

// Functions for defining trees.

func MakeHeap (e []float64, s int) *Heap {
	H := new(Heap)
	for i, v := range e {
		H.Vector = append(H.Vector, v)
		H.number = append(H.number, i + 1)
	}
	if s < 0 || s > len(e) {
		H.size = len(e)
	} else {
		H.size = s
	}
	return H
}

// TODO: need some operations when return a NaN. 

func (H *Heap) Get(i int) (e float64) {
	e = H.Vector[i-1]
	return
}

func (H *Heap) Parent(i int) (n int, e float64) {
	n = i/2
	if n - 1 < 0 || i > H.size {
		e = math.NaN()
	} else {
		e = H.Get(n)
	}
	return
}

func (H *Heap) Left(i int) (n int, e float64) {
	n = 2*i
	if n > H.size || i > H.size {
		e = math.NaN()
	} else {
		e = H.Get(n)
	}
	return
}

func (H *Heap) Right(i int) (n int, e float64) {
	n = 2*i + 1
	if n > H.size || i > H.size {
		e = math.NaN()
	} else {
		e = H.Get(n)
	}
	return
}


func (H *Heap) MaxHeapify(i int) {
	l, lv := H.Left(i)
	r, rv := H.Right(i)
	iv := H.Get(i)
	if l <= H.size && lv > iv {
		largest := l
	} else {
		largest := i
	} 
}




// Sorting vector into non-decreasing or non-increasingorder with HeapSort.
func (H Heap) HeapSort () {
	
}
