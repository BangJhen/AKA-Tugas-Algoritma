package main

// Node merepresentasikan sebuah file atau direktori dalam file system
type Node struct {
	Name     string
	IsDir    bool
	Children []*Node
}

// FlatNode merepresentasikan node yang sudah di-flatten dengan path lengkapnya
type FlatNode struct {
	Name  string
	Path  string
	IsDir bool
}

// flattenTreeToStack mengubah struktur tree menjadi array/stack secara linear
// Menggunakan DFS untuk flatten tree
func flattenTreeToStack(root *Node) []FlatNode {
	if root == nil {
		return []FlatNode{}
	}

	var flatList []FlatNode
	flattenHelper(root, root.Name, &flatList)
	return flatList
}

// flattenHelper adalah helper function untuk flatten tree secara rekursif
func flattenHelper(node *Node, currentPath string, flatList *[]FlatNode) {
	if node == nil {
		return
	}

	// Tambahkan node saat ini ke flat list
	*flatList = append(*flatList, FlatNode{
		Name:  node.Name,
		Path:  currentPath,
		IsDir: node.IsDir,
	})

	// Rekursi untuk semua children
	if node.IsDir {
		var i int
		for i = 0; i < len(node.Children); i++ {
			child := node.Children[i]
			newPath := currentPath + "/" + child.Name
			flattenHelper(child, newPath, flatList)
		}
	}
}
