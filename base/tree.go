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
	// number []int
	size   int
}

// Functions for defining heap tree.

func MakeHeap (e Vector, s int) *Heap {
	H := new(Heap)
	H.Vector = e
	// for i, v := range e {
	// 	H.Vector = append(H.Vector, v)
	// 	H.number = append(H.number, i + 1)
	// }
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
	// n := H.number[i-1]
	// e = H.Vector[n-1]
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


// TODO: Might change the function to one that transfer a slice/Vector to a heap. 
func BuildOrderHeap (e Vector, order bool) *Heap {
	H := MakeHeap(e, len(e))
	for i := H.size/2; i >= 1; i-- {
		H.OrderHeapify(i, order)
	}
	return H
}

func (H *Heap) OrderHeapify(i int, order bool) {
	l, lv := H.Left(i)
	r, rv := H.Right(i)
	ov := H.Get(i)
	var o int
	if order {
		if l <= H.size && lv > ov {
			o = l
			ov = lv
		} else {
			o = i
		}
		if r <= H.size && rv > ov {
			o = r
		}
	} else {
		if l <= H.size && lv < ov {
			o = l
			ov = lv
		} else {
			o = i
		}
		if r <= H.size && rv < ov {
			o = r
		}
	}
	if o != i {
		H.Vector[o-1], H.Vector[i-1] = H.Vector[i-1], H.Vector[o-1]
		// H.number[o-1], H.number[i-1] = H.number[i-1], H.number[o-1]
		H.OrderHeapify(o, order)
	}
}

//==== Priority Queue ====

func (H *Heap) Maximum() float64 {
	return H.Get(1)
}

func (H *Heap) ExtractMax() (max flaot64, err error) {
	if H.size < 1 {
		err = ErrorIllegalIndex
	}
	max = H.Get(1)
	H.Vector[0] = H.Vector[H.size-1]
	H.size = H.size - 1
	H.OrderHeapify(1,true)
	return
}

func (H *Heap) IncreaseKey(i int, key flaot64) {
	if key < H.Vector[i-1] {
		err = ErrorIllegalIndex
	}
	H.Vector[i-1] = key
	for i > 1 
}
	