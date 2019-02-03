package hashTable

// M > N open addressing

type LinearProbingHashST struct {
	N, M int
	Key []hash
	Val []interface{}
}

func Newv2() *LinearProbingHashST {
	return &LinearProbingHashST{
		N:0,
		M:31,
		Key:make([]hash, 31),
		Val:make([]interface{}, 31),
	}
}

func (lst *LinearProbingHashST) resize() {

}

func (lst *LinearProbingHashST) Put(k hash, v interface{}) {
	if lst.N >= lst.M/2 {
		lst.resize()
	}
	var i int
	for i = k.hash(); lst.Key[i] != nil; i = (i+1)%lst.M {
		if lst.Key[i] == k {
			lst.Val[i] = v
			return
		}
	}
	lst.Key[i] = k
	lst.Val[i] = v
	lst.N++
}

func (lst *LinearProbingHashST) Get(k hash) interface{} {

}

func (lst *LinearProbingHashST) Delete(k hash) {

}

