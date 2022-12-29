package composite

type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, other)
}

type NeuronLayer struct {
	Neurons []Neuron
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}

// composite Solution
type NeuronInterface interface {
	Iter() []*Neuron
}

func (n *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)
	for i := range n.Neurons {
		result = append(result, &n.Neurons[i])
	}
	return result
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

// Iter is defined in any tipe that implements NeuronInterface so this can make the trick to connect both NeruonLayer and Neuron
func Connect(left, right NeuronInterface) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}
