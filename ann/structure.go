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
	Incomes   []*Neuron
	Input     float64
	Targets   []*Neuron
	Output    float64
}

// 	n.GenerateInput(x) 
// 	n.GenerateInput(nil) 
func (n *Neuron) GenerateInput (x,w []float64) {
	if x == nil {
		for _,neu := range n.Incomes {
			x = append(x, neu.Output)
		}
	}
	n.Input = n.Aggregate(x, w)
}

// 	n.GenerateOutput() 
func (n *Neuron) GenerateOutput () {
	n.Output = n.Activate[0](n.Input)
}



type Network struct {
	Neurons []Neuron
	Layers  []int
	Links   [][]int
	Weights [][]float64
}

// n1 := ann.Neuron{Input: ann.SU(1), Activation: ann.Linear(1)}
// n2 := ann.Neuron{Input: ann.SU(3), Activation: ann.Sigmoid(2)}
// q := Neuron{Aggregate: Zero, Activate: Bias()}
func NetBuild(neurons []Neuron, layers []int, links [][]int) Network {
	net := Network{Neurons: neurons, Layers: layers, Links: links}
	net.Linkify()
	return net
}

func (net *Network) Linkify () {
	for i,_ := range net.Links {
		net.Neurons[i].Incomes = []*Neuron{}
		net.Neurons[i].Targets = []*Neuron{}
	}
	for i, row := range net.Links {
		for j, col := range row {
			if col > 0 {
				isExist := false
				for _,out := range net.Neurons[j].Incomes {
					if &net.Neurons[i] == out {
						isExist = true
					}
				}
				if !isExist {
					net.Neurons[j].Incomes = append(net.Neurons[j].Incomes, &net.Neurons[i])
					net.Neurons[i].Targets = append(net.Neurons[i].Targets, &net.Neurons[j])
				}
			}
		}
	}
}

func (net *Network) Initialize (weights [][]float64) {
	if weights == nil {
		weights = transLayerToWeight(net.Layers)
	}
	net.Weights = weights
}


// The last element of the parameter must be fasle. 
func Augment (neurons []Neuron, layers []int, links [][]int, augment []bool) ([]Neuron, []int, [][]int) {
	q := Neuron{Aggregate: Zero, Activate: Bias()}
	start := 0
	for key, status := range augment {
		end := start + layers[key]
		if status {
			neurons = append(neurons[:end], append([]Neuron{q}, neurons[end:]...)...)
			for i,_ := range links {
				links[i] = append(links[i][:end], append([]int{0}, links[i][end:]...)...)
			}
			auglinkrow := make([]int, len(neurons))
			links = append(links[:end], append([][]int{auglinkrow}, links[end:]...)...)
			for i := 1; i <= layers[key+1]; i++ {
				links[end][end+i] = 1
			}
			layers[key] += 1
		}
		start += layers[key]
	}
	return neurons, layers, links
}

// type Dataset struct {
// }
