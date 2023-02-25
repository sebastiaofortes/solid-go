package liskovsubstitution

type Notas struct {
	lis []float32
}

func (n *Notas) FiltrarAcima(f float32) []float32 {
	var result []float32
	for i := range n.lis {
		if n.lis[i] > f {
			result = append(result, n.lis[i])
		}
	}
	return result
}

func (n *Notas) FiltrarAbaixo(f float32) []float32 {
	var result []float32
	for i := range n.lis {
		if n.lis[i] < f {
			result = append(result, n.lis[i])
		}
	}
	return result
}

type SubNotas struct {
	Notas
}

// overide
func (n *SubNotas) FiltrarAcima(f float32) []float32 {
	var result []float32
	for i := range n.lis {
		if n.lis[i] > f {
			result = append(result, n.lis[i])
		}
	}
	return result
}

// overide
func (n *SubNotas) FiltrarAbaixo(f float32) []float32 {
	var result []float32
	for i := range n.lis {
		if n.lis[i] < f {
			result = append(result, n.lis[i])
		}
	}
	return result
}

func (n *SubNotas) SetList(l []float32){
	n.lis = l
}