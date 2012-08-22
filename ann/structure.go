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

// structure.go defines structure elements of a neurual network.

// import (
// 	"fmt"
// )


// Elements:
//     activate function: bias, linear, step, ramp, sigmoid, hyperbolic, or gaussian
//     aggregate function: Zeor, SU, or PU
//     output functiona: [-1,1] or [0,1]
type Neuron struct {
	Aggregate func([]float64, []float64) float64
	Activate  [2]func(float64) float64
	Inputs    []*Neuron
	Weights   []float64
	Output    float64
}

type Network struct {
	Neurons []Neuron
	Layers  []int
	Links   [][]int
}

// type Dataset struct {
// }

// 	n.GenerateOutput(x) 
// 	n.GenerateOutput(nil) 
func (n *Neuron) GenerateOutput (x []float64) {
	if x == nil {
		for _,point := range n.Inputs {
			x = append(x, point.Output)
		}
	}
	input := n.Aggregate(x, n.Weights)
	n.Output = n.Activate[0](input)
}

// The last element of the parameter must be fasle. 
// net.Augment([]bool{..., false}) 
func (net *Network) Augment(layer []bool) {
	q := Neuron{Aggregate: Zero, Activate: Bias()}
	start := 0
	for key, status := range layer {
		end := start + net.Layers[key]
		if status {
			net.Neurons = append(net.Neurons[:end], append([]Neuron{q}, net.Neurons[end:]...)...)
			for i,_ := range net.Links {
				net.Links[i] = append(net.Links[i][:end], append([]int{0}, net.Links[i][end:]...)...)
			}
			auglinkrow := make([]int, len(net.Neurons))
			net.Links = append(net.Links[:end], append([][]int{auglinkrow}, net.Links[end:]...)...)
			for i := 1; i <= net.Layers[key+1]; i++ {
				net.Neurons[end+i].Inputs = append(net.Neurons[end+i].Inputs, &net.Neurons[end])
				net.Neurons[end+i].Weights = append(net.Neurons[end+i].Weights, 0)
				net.Links[end][end+i] = 1
			}
			net.Layers[key] += 1
		}
		start += net.Layers[key]
	}
}

// n1 := ann.Neuron{Input: ann.SU(1), Activation: ann.Linear(1)}
// n2 := ann.Neuron{Input: ann.SU(3), Activation: ann.Sigmoid(2)}
// q := Neuron{Aggregate: Zero, Activate: Bias()}
func NetIntialize(neurons []Neuron, layers []int, links [][]int, weights [][]float64) Network {
	net := Network{Neurons: neurons, Layers: layers, Links: links}
	if weights == nil {
		weights = transLayerToWeight(layers)
	}
	for i, row := range links {
		for j, col := range row {
			if col == 1 {
				isExist := false
				for _,out := range net.Neurons[j].Inputs {
					if &net.Neurons[i] == out {
						isExist = true
					}
				}
				if !isExist {
					net.Neurons[j].Inputs = append(net.Neurons[j].Inputs, &net.Neurons[i])
					net.Neurons[j].Weights = append(net.Neurons[j].Weights, weights[i][j])
				}
			}
		}
	}
	return net
}
