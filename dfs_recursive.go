package main

// dfsRecursive melakukan pencarian file menggunakan DFS secara rekursif
func dfsRecursive(node *Node, target string) (string, bool) {
	if node == nil {
		return "", false
	}
	
	// Panggil helper function dengan path awal
	return dfsRecursiveHelper(node, target, node.Name)
}

// dfsRecursiveHelper adalah fungsi helper untuk DFS rekursif
func dfsRecursiveHelper(node *Node, target string, currentPath string) (string, bool) {
	// Base case: jika node nil
	if node == nil {
		return "", false
	}
	
	// Cek apakah node saat ini adalah target
	if !node.IsDir && node.Name == target {
		return currentPath, true
	}
	
	// Jika node adalah direktori, cari di children-nya
	if node.IsDir {
		var i int
		for i = 0; i < len(node.Children); i++ {
			child := node.Children[i]
			
			// Buat path baru
			newPath := currentPath + "/" + child.Name
			
			// Rekursi ke child
			var path string
			var found bool
			path, found = dfsRecursiveHelper(child, target, newPath)
			if found {
				return path, true
			}
		}
	}
	
	// File tidak ditemukan
	return "", false
}
