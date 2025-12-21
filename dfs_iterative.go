package main

// dfsIterative melakukan pencarian pada flat stack dengan iterasi linear
// Kompleksitas: O(n) - linear search
func dfsIterative(flatStack []FlatNode, target string) (string, bool) {
	// Iterasi linear melalui flat stack
	var i int

	// Linear search (1 + 1 + ... + 1) = O(n) = Sum of One
	for i = 0; i < len(flatStack); i++ {
		// Operasi dasar: comparison
		if flatStack[i].Name == target {
			return flatStack[i].Path, true
		}
	}

	// File/direktori tidak ditemukan
	return "", false
}
