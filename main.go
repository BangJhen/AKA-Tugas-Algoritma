package main

import (
	"fmt"
)

func main() {
	// Membuat struktur file system untuk testing
	var root *Node
	root = createSampleFileSystem()
	
	fmt.Println("=== SISTEM PENCARIAN FILE SYSTEM DENGAN DFS ===")
	fmt.Println()
	
	// Menampilkan struktur file system
	fmt.Println("Struktur File System:")
	printFileSystem(root, 0)
	fmt.Println()
	
	// Test pencarian dengan DFS Rekursif
	fmt.Println("=== PENCARIAN DENGAN DFS REKURSIF ===")
	var targetFile string
	targetFile = "config.txt"
	fmt.Printf("Mencari file: %s\n", targetFile)
	
	var path string
	var found bool
	path, found = dfsRecursive(root, targetFile)
	if found {
		fmt.Printf("File ditemukan! Path: %s\n", path)
	} else {
		fmt.Printf("File %s tidak ditemukan.\n", targetFile)
	}
	fmt.Println()
	
	// Test pencarian dengan DFS Iteratif
	fmt.Println("=== PENCARIAN DENGAN DFS ITERATIF ===")
	targetFile = "data.csv"
	fmt.Printf("Mencari file: %s\n", targetFile)
	
	path, found = dfsIterative(root, targetFile)
	if found {
		fmt.Printf("File ditemukan! Path: %s\n", path)
	} else {
		fmt.Printf("File %s tidak ditemukan.\n", targetFile)
	}
	fmt.Println()
	
	// Test pencarian file yang tidak ada
	fmt.Println("=== TEST FILE TIDAK ADA ===")
	targetFile = "notfound.txt"
	fmt.Printf("Mencari file: %s\n", targetFile)
	
	path, found = dfsRecursive(root, targetFile)
	if found {
		fmt.Printf("File ditemukan! Path: %s\n", path)
	} else {
		fmt.Printf("File %s tidak ditemukan.\n", targetFile)
	}
}

// Fungsi untuk membuat sample file system
func createSampleFileSystem() *Node {
	// Membuat struktur direktori
	var root *Node
	root = &Node{Name: "root", IsDir: true}
	
	// Level 1
	var home *Node
	var usr *Node
	var etc *Node
	home = &Node{Name: "home", IsDir: true}
	usr = &Node{Name: "usr", IsDir: true}
	etc = &Node{Name: "etc", IsDir: true}
	
	root.Children = append(root.Children, home, usr, etc)
	
	// Level 2 - home
	var documents *Node
	var downloads *Node
	documents = &Node{Name: "documents", IsDir: true}
	downloads = &Node{Name: "downloads", IsDir: true}
	home.Children = append(home.Children, documents, downloads)
	
	// Level 2 - usr
	var bin *Node
	var lib *Node
	bin = &Node{Name: "bin", IsDir: true}
	lib = &Node{Name: "lib", IsDir: true}
	usr.Children = append(usr.Children, bin, lib)
	
	// Level 2 - etc
	var config *Node
	var settings *Node
	config = &Node{Name: "config.txt", IsDir: false}
	settings = &Node{Name: "settings.ini", IsDir: false}
	etc.Children = append(etc.Children, config, settings)
	
	// Level 3 - documents
	var report *Node
	var data *Node
	report = &Node{Name: "report.pdf", IsDir: false}
	data = &Node{Name: "data.csv", IsDir: false}
	documents.Children = append(documents.Children, report, data)
	
	// Level 3 - downloads
	var image *Node
	var video *Node
	image = &Node{Name: "image.jpg", IsDir: false}
	video = &Node{Name: "video.mp4", IsDir: false}
	downloads.Children = append(downloads.Children, image, video)
	
	// Level 3 - bin
	var app1 *Node
	var app2 *Node
	app1 = &Node{Name: "app1.exe", IsDir: false}
	app2 = &Node{Name: "app2.exe", IsDir: false}
	bin.Children = append(bin.Children, app1, app2)
	
	return root
}

// Fungsi untuk print struktur file system
func printFileSystem(node *Node, level int) {
	if node == nil {
		return
	}
	
	// Print indentasi
	var i int
	for i = 0; i < level; i++ {
		fmt.Print("  ")
	}
	
	// Print nama node
	if node.IsDir {
		fmt.Printf("[%s/]\n", node.Name)
	} else {
		fmt.Printf("%s\n", node.Name)
	}
	
	// Print children
	for i = 0; i < len(node.Children); i++ {
		var child *Node
		child = node.Children[i]
		printFileSystem(child, level+1)
	}
}
