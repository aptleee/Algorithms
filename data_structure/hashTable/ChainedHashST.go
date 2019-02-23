package hashTable

// INSERT, SEARCH, DELETE
// simpler notion of an ordinary array.
// directly addressing into an ordinary array
// we can take directly addressing
// when we can afford to allocate an array that has one position for every possible key
// the number of keys actually stored is small relative to the total number of
// possible keys, hash table become an effective alternative to directly addressing
// an array
// chaining and open addressing
// collision

// h(k) hash value
// random mixing and chopping
// collision resolution technique

// CHAINED-HASH-INSERT(T,x)
// insert x at the head of list T[h(x.key)]

// CHAINED-HASH-SEARCH(T,k)
// search for an element with key k in list T[h(k)]

// CHAINED-HASH-DELETE(T, x)
// delete x from the list T[h(x.key)]

// insertion worst case O(1)
// search worst case running time is proportional to the length of the list
// delete worst case O(1) for doubly linked list

// m slots that stores n elements, define load factor a for n/m
// simple uniform hashing: any given elements is equally likely to hash into any of the m
// slots, independently of where any other elements has hashed to.
// successful search and unsuccessful search

// theorem 11.1
// in a hash table in which collision are resolved by chaining, an unsuccessful search takes
// average-case time big theta of (1+a), under the assumption of simple uniform hashing

// theorem 11.2
// in a hash table in which collisions are resolved by chaining, a successful search takes
// average-case time big theta of (1+a), under the assumption of simple uniform hashing

// hash function
// A good hash function satisfies (approximately) the assumption of simple uniform hashing
// a good approach derives the hash value in a way that we expect to be independent of any
// patterns that might exist in the data
// the data or the key may have a distribution
// division method use prime numbers that is unrelated to any patterns in the distribution
// of keys
// Finally, we note that some applications of hash functions might require stronger properties
// that are provided by simple uniform hashing.
// For example, we might want keys that are “close” in some sense to yield hash values that are far apart.

// most hash functions assume that the universe of keys in the set N = {0, 1, 2,...}
// of natural numbers. Thus, if the keys are not natural numbers, we find a way to
// interpret them as natural numbers.

// division method
// h(k) = k % m
// usually avoid certain values of m, such as a power of 2
// since if m = 2^p, then h(k) is just the p lowest-order bits of k.

// multiplication method
// h(k) = floor(m(kA%1))  m=2^p, s = A * 2^w, w is the word size of bits

// universal hashing
// choosing hash functions randomly
// in universal hashing, at the beginning of execution we select the hash func
// at random from a carefully designed class of func.

// universal means randomly choose a hash function from H, the chance of a collision between
// distinct keys k and l is no more than 1/m.

// designing a universal class of hash funcs
// let p be a large prime number
// Zp = {0, 1, 2,..., p-1}
// Zp* = {1, 2,..., p-1}
// the number of universal keys is greater than number of slots in th hash table
// p > m
// for ant a in Zp* and b in Zp, using a linear transformation
// hab(k) = ((ak+b)%p)%m

// k != l
// r = (ak+b)%p
// s = (al+b)%p

// r-s = a(k-l)%p != 0

// hash := 0
// for i:=0; i < len(s); i++ {
//	hash = (R * hash + int(s[i])) % M
// }
// hash = (((day * R + month) % M) * R + year) % M
// M = 31
import (
	"AlgorithmsGo-master/data_structure/linked_list"
)

type chainedHashTable struct {
	N, M int
	Table []linkedlist.DoublyLinkedList
}

func New(M int) *chainedHashTable {
	return &chainedHashTable{
		N: 0,
		M: M,
		Table: make([]linkedlist.DoublyLinkedList, 31),
	}
}



func (ta *chainedHashTable) resize(m int) {
	t := New(m)
	for i := 0; i < m; i++ {
		for n := ta.Table[i].Front(); n != nil; n = n.Next() {
			t.Put(n.Key, ta.Table[i].Search(n.Key))
		}
	}
	ta.M = t.M
	ta.N = t.N
	ta.Table = t.Table
}

func (ta *chainedHashTable) Get(key hash) interface{} {
	a := ta.Table[key.hash(ta.M)].Search(key)
	if a == nil {
		return -1
	} else {
		return a.Val
	}
}

func (ta *chainedHashTable) Put(key hash, v int) {
	if ta.N > 8*ta.M {
		ta.resize(2*ta.M)
	}
	a := ta.Table[key.hash(ta.M)].Search(key)
	if a == nil {
		ta.Table[key.hash(ta.M)].AddAtBeg(key, v)
		ta.N++
	} else {
			a.Val = v
	}

}

func (ta *chainedHashTable) Delete(key hash) interface{} {
	a := ta.Table[key.hash(ta.M)].Search(key)
	if  a == nil{
		return -1
	} else {
		ta.Table[key.hash(ta.M)].Delete(key)
		ta.N--
		if ta.N <= 2*ta.M {
			ta.resize(ta.M/2)
		}
		return a
	}

}





