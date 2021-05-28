//208. Implement Trie (Prefix Tree)
//https://leetcode.com/problems/implement-trie-prefix-tree/
type Trie struct {
	Child map[rune]*Trie
	End   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	link := make(map[rune]*Trie)
	return Trie{End: false, Child: link}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for _, v := range word {
		if child, ok := cur.Child[v]; ok {
			cur = child
		} else {
			newChild := &Trie{End: false, Child: make(map[rune]*Trie)}
			cur.Child[v] = newChild
			cur = newChild
		}
	}
	cur.End = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for _, v := range word {
		if child, ok := cur.Child[v]; ok {
			cur = child
			continue
		}
		return false
	}
	return cur.End
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, v := range prefix {
		if child, ok := cur.Child[v]; ok {
			cur = child
			continue
		}
		return false
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */