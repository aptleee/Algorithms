package Trie

type Trie struct {
	arr [26]*Trie
	end bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{arr: [26]*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	head := this
	for i := 0; i < len(word); i++ {
		if head.arr[word[i]-'a'] == nil {
			head.arr[word[i]-'a'] = &Trie{}
		}
		head = head.arr[word[i]-'a']
	}
	head.end = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	head := this
	for i := 0; i < len(word); i++ {
		if head.arr[word[i]-'a'] == nil {
			return false
		}
		head = head.arr[word[i]-'a']
	}
	return head.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	head := this
	for i := 0; i < len(prefix); i++ {
		if head.arr[prefix[i]-'a'] == nil {
			return false
		}
		head = head.arr[prefix[i]-'a']
	}
	return true
}