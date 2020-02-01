package dom

// // NewNode return a new node
// func NewNode() Node {
// 	return &BaseNode{
// 		childNodes: NodeList{
// 			Length: 0,
// 			list:   []*Node{},
// 		},
// 		firstChild:      nil,
// 		lastChild:       nil,
// 		isConnected:     false,
// 		nodeType:        0,
// 		parentNode:      nil,
// 		parentElement:   nil,
// 		previousSibling: nil,
// 		nextSibling:     nil,
// 	}
// }

// // AppendChild adds the specified childNode argument
// // as the last child to the current node.
// func (n *BaseNode) AppendChild(childNode Node) (*Node, error) {
// 	cn := &childNode

// 	// Avoid loop
// 	if childNode.Contains(n) {
// 		return nil, errLoop
// 	}

// 	// Setting parent and next sibling node
// 	cn.previousSibling = n.lastChild
// 	cn.parentNode = n

// 	// Setting first and last child
// 	n.lastChild = cn
// 	if n.childNodes.Length == 0 {
// 		n.firstChild = cn
// 	}

// 	// Appending to the child node
// 	n.childNodes.list = append(n.childNodes.list, cn)
// 	n.childNodes.Length++

// 	return cn, nil
// }

// // CloneNode return a clone of the current node.
// func (n *BaseNode) CloneNode(deep bool) *Node {
// 	clone := NewNode()

// 	clone.nodeType = n.nodeType

// 	if deep {
// 		n.childNodes.forEach(func(index int, child *Node) {
// 			c := child.CloneNode(deep)
// 			clone.AppendChild(*c)
// 		})
// 	}

// 	return clone
// }

// // CompareDocumentPosition method compares the
// // position of the given node against another node in any document.
// func (n *BaseNode) CompareDocumentPosition(otherNode Node) {
// 	// TODO func (n *BaseNode) CompareDocumentPosition()
// }

// // Contains return a bool value wether or
// // not a node is a descendant of the calling.
// func (n *BaseNode) Contains(otherNode *Node) (contain bool) {
// 	for _, childNode := range n.childNodes.list {
// 		if childNode.IsEqualNode(otherNode) {
// 			contain = true
// 			break
// 		}
// 	}

// 	if !contain {
// 		for _, childNode := range n.childNodes.list {
// 			if childNode.Contains(otherNode) {
// 				contain = true
// 				break
// 			}
// 		}
// 	}

// 	return contain
// }

// // GetRootNode return the context object's root
// func (n *BaseNode) GetRootNode() (root *Node) {
// 	// TODO shadow root
// 	if n.parentNode != nil {
// 		root = n.parentNode.GetRootNode()
// 	} else {
// 		root = n
// 	}

// 	return
// }

// // HasChildNodes return wether or not
// // the element has any child node
// func (n *BaseNode) HasChildNodes() bool {
// 	return n.childNodes.Length > 0
// }

// // InsertBefore the given node before the reference node
// func (n *BaseNode) InsertBefore(newNode Node, referenceNode *Node) (*Node, error) {
// 	index := n.childNodes.IndexOf(referenceNode)

// 	// Error reference node is not a child
// 	if index == -1 {
// 		return nil, errNotChild
// 	}

// 	// Inserting the child
// 	childs1 := make([]*Node, 0)
// 	if index != 0 {
// 		childs1 = n.childNodes.list[:index-1]
// 	}

// 	childs2 := n.childNodes.list[index:]

// 	newList := append(childs1, &newNode)
// 	newList = append(childs1, childs2...)
// 	n.childNodes.list = newList
// 	n.childNodes.Length++

// 	// Updating siblings
// 	if prev := n.childNodes.Item(index - 1); prev != nil {
// 		prev.nextSibling = &newNode
// 	} else {
// 		// Updating first child
// 		n.firstChild = &newNode
// 	}
// 	referenceNode.previousSibling = &newNode

// 	return &newNode, nil
// }

// // IsEqualNode return wether or note the given
// // node are equal
// func (n *BaseNode) IsEqualNode(otherNode *Node) (equal bool) {
// 	// Checking the list of childrens length
// 	if n.childNodes.Length != otherNode.childNodes.Length {
// 		goto notEqual
// 	}

// 	// Check children
// 	for i, child := range n.childNodes.list {
// 		if otherChild := otherNode.childNodes.Item(i); otherChild != nil {
// 			if !otherChild.IsEqualNode(child) {
// 				goto notEqual
// 			}
// 		}
// 	}

// 	// Check nodeType
// 	if n.nodeType != otherNode.nodeType {
// 		goto notEqual
// 	}

// 	return true

// notEqual:
// 	return false
// }

// // IsSameNode return a value indicating whether or
// // not the two nodes are the same (same reference)
// func (n *BaseNode) IsSameNode(otherNode *Node) bool {
// 	return n == otherNode
// }

// // RemoveChild removes a child node
// //  and returns the removed node.
// func (n *BaseNode) RemoveChild(childNode *Node) (cn *Node, err error) {

// 	// Avoid loop
// 	if childNode.Contains(n) {
// 		return nil, errLoop
// 	}

// 	// Removing parent and siblings node
// 	childNode.parentNode = nil
// 	childNode.previousSibling = nil
// 	childNode.nextSibling = nil

// 	// Updating siblings
// 	index := n.childNodes.IndexOf(childNode)
// 	prev := n.childNodes.Item(index - 1)
// 	next := n.childNodes.Item(index + 1)

// 	// Removing child node
// 	// Children not found
// 	if index == -1 {
// 		return childNode, errNotChild
// 	}

// 	childs1 := make([]*Node, 0)
// 	childs2 := make([]*Node, 0)

// 	if index != 0 {
// 		childs1 = n.childNodes.list[:index-1]
// 	}
// 	if n.childNodes.Length > (index + 1) {
// 		childs2 = n.childNodes.list[index+1:]
// 	}

// 	n.childNodes.list = append(childs1, childs2...)
// 	n.childNodes.Length--

// 	// Updating last & first child
// 	if prev != nil {
// 		prev.nextSibling = next
// 	} else {
// 		n.firstChild = next
// 	}

// 	if next != nil {
// 		next.previousSibling = prev
// 	} else {
// 		n.lastChild = prev
// 	}

// 	return childNode, nil
// }

// // ReplaceChild method replaces a child node with the given node.
// func (n *BaseNode) ReplaceChild(newChild, oldChild *Node) {}

// // ERRORS
// var (
// 	errLoop     = errors.New("The child element contains the parent")
// 	errNotChild = errors.New("The given node is not a child of this node")
// )
