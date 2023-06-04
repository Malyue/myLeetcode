package main

// longest Common Prefix

type Node struct {
	Value    string
	Children map[string]*Node
}

type Trim struct {
	root *Node
}

func NewPrefixTrie() *Trim {
	return &Trim{
		root: &Node{
			Value:    "",
			Children: make(map[string]*Node),
		},
	}
}

func (t *Trim) Insert(str string) {
	// get the root
	node := t.root
	// for range the str
	for _, el := range str {
		char := string(el)
		// if exists,to the childNode,otherwise create a new node
		if _, ok := node.Children[char]; !ok {
			node.Children[char] = &Node{
				Value:    char,
				Children: make(map[string]*Node),
			}
		}
		node = node.Children[char]
	}
}

func getMinStr(strs []string) string {
	min := strs[0]
	for _, str := range strs {
		if len(str) < len(min) {
			min = str
		}
	}
	return min
}

func longestCommonPrefix(strs []string) string {
	// find the longest Commpn Prefix
	// the easiest is to range all the strs,while the value is not common,return the number
	// p := strs[0]
	// for _,s := range strs {
	//     i := 0
	//     for ;i < len(s) && i < len(p) && p[i] == s[i];i ++ {}
	//     p = p[:i]
	// }
	// return p

	// use the trim tree
	if len(strs) == 1 {
		return strs[0]
	}

	var result string
	prefixTrie := NewPrefixTrie()

	// insert the strs into the tree
	for _, str := range strs {
		prefixTrie.Insert(str)
	}

	node := prefixTrie.root

	// get the minstr,because the longest common prefix must be less than the minstr
	minstr := getMinStr(strs)
	// find the longest common prefix
	for _, el := range minstr {
		char := string(el)
		// if the length of the node.Children is 1,it is the common prefix
		if len(node.Children) == 1 {
			result += char
			node = node.Children[char]
		} else {
			return result
		}
	}

	return result

}
