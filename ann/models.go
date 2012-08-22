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

// models.go defines frequently used neurual networks.

func FFNN (neurons []Neuron, layers []int, weights [][]float64) Network {
	links := transLayerToLink(layers)
	augment := make([]bool, len(layers))
	for i,_ := range layers {
		if i < len(layers)-1 {
			outstart := sumInt(layers[:i])
			outend := sumInt(layers[:i+1]) - 1
			instart := sumInt(layers[:i+1])
			inend := sumInt(layers[:i+2]) - 1
			for k := outstart; k <= outend; k++ {
				for l := instart; l <= inend; l++ {
					links[k][l] = 1
				}
			}
			augment[i] = true
		} else {
			augment[i] = false
		}
	}
	net := NetIntialize(neurons, layers, links, weights)
	net.Augment(augment)
	return net
}

// func PUNN () {
// }

// func SRNN () {
// }

// func TDNN () {
// }

// func CNN () {
// }
