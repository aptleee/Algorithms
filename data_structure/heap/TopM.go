package main

func TopM(M int, input []int) *tmp{
	MinQ := &tmp{}

	buildHeap(MinQ)
	for _, x := range input {
		if MinQ.Len() < M {
			Push(MinQ, x)
		} else if x > MinQ.Front().(int) {
			Pop(MinQ)
			Push(MinQ, x)
		}
	}
	return MinQ
}
