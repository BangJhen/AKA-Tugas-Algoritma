package main

// dfsRecursive melakukan pencarian pada flat stack dengan linear recursion
// Relasi rekurensi: T(n) = T(n-1) + 1
// Kompleksitas: O(n)
func dfsRecursive(flatStack []FlatNode, target string, index int) (string, bool) {
	// Operasi dasar: comparison (seperti contoh di modul)
	if flatStack[index].Name == target {
		return flatStack[index].Path, true
	} else {
		// Linear recursion: T(n) = T(n-1) + 1
		// Sama seperti factorial/power di modul
		return dfsRecursive(flatStack, target, index+1)
	}
}
