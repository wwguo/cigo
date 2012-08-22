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

// Package ann provides interfaces and types for artifical neural networks.
package ann

// utils.go defines utilty functions.

// import (
// 	"math/rand"
// )


func sumFloat(a []float64) (s float64) {
    for _, v := range a {
        s += v
    }
    return
}

func sumInt(a []int) (s int) {
    for _, v := range a {
        s += v
    }
    return
}

func transLayerToLink (layer []int) (matrix [][]int) {
	length := sumInt(layer)
	for i := 0; i < length; i++ {
		matrix = append(matrix, make([]int, length))
	}
	return
}

func transLayerToWeight (layer []int) (matrix [][]float64) {
	length := sumInt(layer)
	for i := 0; i < length; i++ {
		matrix = append(matrix, make([]float64, length))
	}
	return
}
