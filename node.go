package main

// Node merepresentasikan sebuah file atau direktori dalam file system
type Node struct {
	Name     string  // Nama file atau direktori
	IsDir    bool    // True jika ini adalah direktori, false jika file
	Children []*Node // Daftar children (hanya untuk direktori)
}

// Stack untuk implementasi DFS iteratif
type Stack struct {
	items []StackItem
}

// StackItem menyimpan node dan path-nya
type StackItem struct {
	node *Node
	path string
}

// pushStack menambahkan item ke stack
func pushStack(s *Stack, item StackItem) {
	s.items = append(s.items, item)
}

// popStack mengambil dan menghapus item teratas dari stack
func popStack(s *Stack) (StackItem, bool) {
	if len(s.items) == 0 {
		return StackItem{}, false
	}
	
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// isStackEmpty mengecek apakah stack kosong
func isStackEmpty(s *Stack) bool {
	return len(s.items) == 0
}
