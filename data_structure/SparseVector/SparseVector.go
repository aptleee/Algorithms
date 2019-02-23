package SparseVector


type SparseVector struct {
	Sv map[int]float64
}


func (sv *SparseVector) Put(i int, x float64) {
	sv.Sv[i] = x
}

func (sv *SparseVector) Get(i int) float64 {
	v, ok := sv.Sv[i]
	if !ok {
		return 0.0
	}
	return v
}

func (sv *SparseVector) Dot(that []float64) float64 {
	var sum float64 = 0
	for k := range sv.Sv{
		sum += that[k] * sv.Sv[k]
	}
	return sum
}

