package main

import (
	"fmt"
	"time"
)

func main() {
	// Membuat struktur file system untuk testing
	var userRoot *Node
	var target, path string
	var pilih, fsChoice int
	var found bool

	// Menu pemilihan file system
	fmt.Println("=== PILIH FILE SYSTEM ===")
	fmt.Println("1. PacOS (Default - Medium, ~50 nodes)")
	fmt.Println("2. SmallOS (Small - ~20 nodes)")
	fmt.Println("3. LargeOS (Large - ~100+ nodes)")
	fmt.Println("4. FlatOS (2 Level - Root → Files)")
	fmt.Println("5. BinaryTreeOS (Perfect Binary Tree - 31 nodes)")
	fmt.Println("6. Custom Binary Tree (Depth 3-7)")
	fmt.Print("Pilih: ")
	fmt.Scanln(&fsChoice)

	switch fsChoice {
	case 2:
		userRoot = createSmallFileSystem()
	case 3:
		userRoot = createLargeFileSystem()
	case 4:
		userRoot = createFlatFileSystem()
	case 5:
		userRoot = createBalancedBinaryTree()
	case 6:
		var depth int
		fmt.Print("Masukkan kedalaman tree (3-7): ")
		fmt.Scanln(&depth)
		if depth < 3 {
			depth = 3
		} else if depth > 7 {
			depth = 7
		}
		userRoot = createPerfectBinaryTree(depth)
	default:
		userRoot = createSampleFileSystem()
	}

	printFileSystemStats(userRoot)

	// Flatten tree menjadi stack SEKALI di awal (di luar loop)
	// Sehingga overhead flatten tidak mempengaruhi perbandingan rekursif vs iteratif
	flatStack := flattenTreeToStack(userRoot)

	for pilih != 4 {
		fmt.Println("=== SISTEM PENCARIAN FILE SYSTEM DENGAN DFS ===")
		fmt.Println("1. Tampilkan Diretory")
		fmt.Println("2. Pencarian dengan DFS Rekursif (Linear Recursion)")
		fmt.Println("3. Pencarian dengan DFS Iteratif (Linear Loop)")
		fmt.Println("4. Exit")
		fmt.Print("Pilih: ")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			fmt.Println("Struktur File System:")
			printFileSystem(userRoot, 0)
			fmt.Println()
		case 2:
			fmt.Print("Masukkan nama file/direktori yang dicari: ")
			fmt.Scanln(&target)

			start := time.Now()
			path, found = dfsRecursive(flatStack, target, 0)
			elapsed := time.Since(start)

			if found {
				fmt.Printf("File/Directory Ditemukan! Path: %s\n", path)
			} else {
				fmt.Printf("File %s tidak ditemukan.\n", target)
			}
			fmt.Printf("Waktu pencarian: %v (%d ns)\n", elapsed, elapsed.Nanoseconds())
			fmt.Println()
		case 3:
			fmt.Print("Masukkan nama file/direktori yang dicari: ")
			fmt.Scanln(&target)

			start := time.Now()
			path, found = dfsIterative(flatStack, target)
			elapsed := time.Since(start)

			if found {
				fmt.Printf("File/Directory Ditemukan! Path: %s\n", path)
			} else {
				fmt.Printf("File %s tidak ditemukan.\n", target)
			}
			fmt.Printf("Waktu pencarian: %v (%d ns)\n", elapsed, elapsed.Nanoseconds())
			fmt.Println()
		case 4:
			fmt.Println("Terima kasih!")
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
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
		child := node.Children[i]
		printFileSystem(child, level+1)
	}
}
