package minigin

import (
	"fmt"
)

type node struct {
	part     string
	isValid  bool
	children []*node
	handler  map[string]HandlerFunc
}

func newNode(method string, values ...interface{}) *node {
	n := &node{}
	n.children = make([]*node, 0)
	n.handler = make(map[string]HandlerFunc)
	for _, value := range values {
		switch v := value.(type) {
		case string:
			n.part = v
		case bool:
			n.isValid = v
		case HandlerFunc:
			n.handler[method] = v
		default:
			fmt.Println("Unknown value type:", v)
		}
	}
	return n
}

func updateNode(nd *node, method string, isEnd bool, hdl HandlerFunc) *node {
	nd.isValid = isEnd
	nd.handler[method] = hdl
	return nd
}

func updateTrie(root *node, method string, parts []string, handler HandlerFunc) {
	l := len(parts)
	curNode := root
	if l == 1 { // parts = ["/"]
		updateNode(root, method, true, handler)
		return
	}
	for i := 1; i < l; i++ {
		part := parts[i]
		isEnd := false
		if i == l-1 {
			isEnd = true
		}
		sr := findMatch(curNode, part)
		if sr == nil {
			temp := newNode(method, part, isEnd, handler)
			curNode.children = append(curNode.children, temp)
			curNode = temp
		} else {
			if isEnd {
				sr = updateNode(sr, method, isEnd, handler)
			}
			curNode = sr
		}
	}
}

func findMatch(root *node, target string) *node {
	for _, ptr := range root.children {
		if ptr != nil && ptr.part == target {
			return ptr
		}
	}
	return nil
}

func searchTrie(index int, nd *node, method string, parts []string, paramsMap map[string]string) HandlerFunc {
	if matchPart(parts[index], nd.part, paramsMap) {
		if index == len(parts)-1 {
			if nd.handler[method] != nil && nd.isValid {
				return nd.handler[method]
			} else {
				return nil
			}
		} else {
			for _, child := range nd.children {
				ret := searchTrie(index+1, child, method, parts, paramsMap)
				if ret != nil {
					return ret
				}
			}
		}
	}
	return nil
}

func matchPart(cand string, obj string, m map[string]string) bool {
	if cand == obj {
		return true
	} else if obj[0] == '*' {
		m[obj] = cand
		return true
	} else {
		return false
	}
}
