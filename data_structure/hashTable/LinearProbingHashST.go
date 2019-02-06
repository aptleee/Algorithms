package hashTable

import "fmt"

// M > N open addressing

type LinearProbingHashST struct {
	N, M int
	Key []hash
	Val []interface{}
}

func Newv1() *LinearProbingHashST {
	return &LinearProbingHashST{
		N:0,
		M:31,
		Key:make([]hash, 31),
		Val:make([]interface{}, 31),
	}
}

func Newv2(M int) *LinearProbingHashST {
	return &LinearProbingHashST{
		N:0,
		M:M,
		Key:make([]hash, M),
		Val:make([]interface{}, M),
	}
}

type hash interface {
	hash(M int) int
}

type myint int

func (i myint) hash(M int) int {
	return int(i) % M
}


func main() {
	a := Newv1()
	for i:=0; i < 11; i++ {
		a.Put(myint(i), i)
	}
	a.Put(myint(41), 41)
	fmt.Println(a)
	for i := 11; i < 31; i++ {
		a.Put(myint(i), i)
	}
	fmt.Println(a)
	fmt.Println(a.Get(myint(41)))
}


func (lst *LinearProbingHashST) resize(M int) {
	t := Newv2(M)
	for i:=0; i < lst.M; i++ {
		if lst.Key[i] != nil {
			t.Put(lst.Key[i], lst.Val[i]) // put will not call resize again, for t.N will always be less than t.M/4
		}
	}
	lst.Key = t.Key
	lst.Val = t.Val
	lst.M = t.M
	// lst.N is still lst.N
}

func (lst *LinearProbingHashST) Put(k hash, v interface{}) {
	if lst.N >= lst.M/2 {
		lst.resize(2*lst.M) //
	}
	var i int
	for i = k.hash(lst.M); lst.Key[i] != nil; i = (i+1)%lst.M {
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
	var i int
	for i = k.hash(lst.M); lst.Key[i] != nil; i = (i+1)%lst.M {
		if lst.Key[i] == k {
			return lst.Val[i]
		}
	}
	return nil
}

func (lst *LinearProbingHashST) contain(k hash) bool {
	for i := 0; i < lst.M; i++ {
		if lst.Key[i] == k {
			return true
		}
	}
	return false
}
func (lst *LinearProbingHashST) Delete(k hash) { // cannot just set the lst.Key[k.hash()] to nil, which will make the
// elements after it unavailable, so we need to reinsert into the table all of the keys
// in the cluster to the right of the deleted key
	if lst.contain(k) == false {
		return
	}
	i := k.hash(lst.M)
	for k != lst.Key[i] {
		i = (i+1)%lst.M
	}

	lst.Key[i] = nil
	lst.Val[i] = nil
	i = (i+1)%lst.M
	for lst.Key[i] != nil {
		keyToRedo := lst.Key[i]
		ValToRedo := lst.Val[i]
		lst.Key[i] = nil
		lst.Val[i] = nil
		lst.N--
		lst.Put(keyToRedo, ValToRedo)
		i = (i+1)%lst.M
	}
	lst.N--
	if lst.N > 0 && lst.N == lst.M / 8 {
		lst.resize(lst.M/2)
	}
}

