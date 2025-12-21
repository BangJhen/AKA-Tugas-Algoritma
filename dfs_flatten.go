package main

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

// dfsRecursiveOnFlatStack melakukan pencarian pada flat stack secara linear recursion
// Ini adalah linear recursion: T(n) = T(n-1) + 1
func dfsRecursiveOnFlatStack(flatStack []FlatNode, target string, index int) (string, bool) {
	// Base case: sudah sampai akhir stack
	if index >= len(flatStack) {
		return "", false
	}

	// Operasi dasar: cek node saat ini
	if flatStack[index].Name == target {
		return flatStack[index].Path, true
	}

	// Linear recursion: lanjut ke index berikutnya
	// T(n) = T(n-1) + 1
	return dfsRecursiveOnFlatStack(flatStack, target, index+1)
}

// dfsWithFlatten adalah fungsi utama yang menggabungkan flatten + search
func dfsWithFlatten(root *Node, target string) (string, bool) {
	// Step 1: Flatten tree menjadi linear stack
	flatStack := flattenTreeToStack(root)

	// Step 2: Search pada flat stack dengan linear recursion
	return dfsRecursiveOnFlatStack(flatStack, target, 0)
}

// printFlatStack untuk debugging - menampilkan hasil flatten
func printFlatStack(flatStack []FlatNode) {
	var i int
	for i = 0; i < len(flatStack); i++ {
		node := flatStack[i]
		nodeType := "FILE"
		if node.IsDir {
			nodeType = "DIR "
		}
		// Menggunakan Printf dengan format yang benar
		var format string
		format = "[%d] %s | %s | %s\n"
		// Menggunakan variabel terpisah untuk setiap argumen
		var idx int
		var typ string
		var name string
		var path string
		idx = i
		typ = nodeType
		name = node.Name
		path = node.Path
		// Memanggil Printf dengan semua argumen
		printFormattedNode(format, idx, typ, name, path)
	}
}

// printFormattedNode adalah helper untuk print dengan format
func printFormattedNode(format string, idx int, nodeType string, name string, path string) {
	// Implementasi sederhana tanpa Printf
	print("[")
	print(idx)
	print("] ")
	print(nodeType)
	print(" | ")
	print(name)
	print(" | ")
	print(path)
	print("\n")
}
