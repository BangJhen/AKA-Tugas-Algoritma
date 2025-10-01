package main

// dfsIterative melakukan pencarian file menggunakan DFS secara iteratif dengan stack
func dfsIterative(root *Node, target string) (string, bool) {
	if root == nil {
		return "", false
	}
	
	// Inisialisasi stack
	stack := &Stack{}
	
	// Push root ke stack
	rootItem := StackItem{
		node: root,
		path: root.Name,
	}
	pushStack(stack, rootItem)
	
	// Proses selama stack tidak kosong
	for !isStackEmpty(stack) {
		// Pop item dari stack
		item, ok := popStack(stack)
		if !ok {
			break
		}
		
		var currentNode *Node
		var currentPath string
		currentNode = item.node
		currentPath = item.path
		
		// Skip jika node nil
		if currentNode == nil {
			continue
		}
		
		// Cek apakah node saat ini adalah target
		if !currentNode.IsDir && currentNode.Name == target {
			return currentPath, true
		}
		
		// Jika node adalah direktori, push semua children ke stack
		if currentNode.IsDir {
			// Push children dalam urutan terbalik agar diproses dari kiri ke kanan
			var i int
			for i = len(currentNode.Children) - 1; i >= 0; i-- {
				child := currentNode.Children[i]
				childItem := StackItem{
					node: child,
					path: currentPath + "/" + child.Name,
				}
				pushStack(stack, childItem)
			}
		}
	}
	
	// File tidak ditemukan
	return "", false
}
