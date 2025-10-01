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
	targetFile = "scheduler"
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
	targetFile = "math_assignment"
	fmt.Printf("Mencari file: %s\n", targetFile)
	
	path, found = dfsIterative(root, targetFile)
	if found {
		fmt.Printf("File ditemukan! Path: %s\n", path)
	} else {
		fmt.Printf("File %s tidak ditemukan.\n", targetFile)
	}
	fmt.Println()
	
	// Test pencarian file lainnya
	fmt.Println("=== TEST PENCARIAN FILE LAIN ===")
	targetFile = "compiler"
	fmt.Printf("Mencari file: %s\n", targetFile)
	
	path, found = dfsRecursive(root, targetFile)
	if found {
		fmt.Printf("File ditemukan! Path: %s\n", path)
	} else {
		fmt.Printf("File %s tidak ditemukan.\n", targetFile)
	}
	fmt.Println()
	
	// Test pencarian file yang tidak ada
	fmt.Println("=== TEST FILE TIDAK ADA ===")
	targetFile = "virus_scanner"
	fmt.Printf("Mencari file: %s\n", targetFile)
	
	path, found = dfsIterative(root, targetFile)
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
	root = &Node{Name: "PacOS", IsDir: true}
	
	// Level 1 - Direktori utama sistem
	var system *Node
	var users *Node
	var programs *Node
	var temp *Node
	system = &Node{Name: "System", IsDir: true}
	users = &Node{Name: "Users", IsDir: true}
	programs = &Node{Name: "Programs", IsDir: true}
	temp = &Node{Name: "Temp", IsDir: true}
	
	root.Children = append(root.Children, system, users, programs, temp)
	
	// Level 2 - System
	var kernel *Node
	var drivers *Node
	var logs *Node
	kernel = &Node{Name: "Kernel", IsDir: true}
	drivers = &Node{Name: "Drivers", IsDir: true}
	logs = &Node{Name: "Logs", IsDir: true}
	system.Children = append(system.Children, kernel, drivers, logs)
	
	// Level 2 - Users
	var admin *Node
	var guest *Node
	var student *Node
	admin = &Node{Name: "Admin", IsDir: true}
	guest = &Node{Name: "Guest", IsDir: true}
	student = &Node{Name: "Student", IsDir: true}
	users.Children = append(users.Children, admin, guest, student)
	
	// Level 2 - Programs
	var games *Node
	var office *Node
	var development *Node
	games = &Node{Name: "Games", IsDir: true}
	office = &Node{Name: "Office", IsDir: true}
	development = &Node{Name: "Development", IsDir: true}
	programs.Children = append(programs.Children, games, office, development)
	
	// Level 3 - System/Kernel
	var bootloader *Node
	var scheduler *Node
	var memory *Node
	bootloader = &Node{Name: "bootloader", IsDir: false}
	scheduler = &Node{Name: "scheduler", IsDir: false}
	memory = &Node{Name: "memory_manager", IsDir: false}
	kernel.Children = append(kernel.Children, bootloader, scheduler, memory)
	
	// Level 3 - System/Drivers
	var graphics *Node
	var audio *Node
	var network *Node
	graphics = &Node{Name: "graphics_driver", IsDir: false}
	audio = &Node{Name: "audio_driver", IsDir: false}
	network = &Node{Name: "network_driver", IsDir: false}
	drivers.Children = append(drivers.Children, graphics, audio, network)
	
	// Level 3 - System/Logs
	var systemLog *Node
	var errorLog *Node
	var accessLog *Node
	systemLog = &Node{Name: "system_log", IsDir: false}
	errorLog = &Node{Name: "error_log", IsDir: false}
	accessLog = &Node{Name: "access_log", IsDir: false}
	logs.Children = append(logs.Children, systemLog, errorLog, accessLog)
	
	// Level 3 - Users/Admin
	var desktop *Node
	var documents *Node
	var settings *Node
	desktop = &Node{Name: "Desktop", IsDir: true}
	documents = &Node{Name: "Documents", IsDir: true}
	settings = &Node{Name: "Settings", IsDir: true}
	admin.Children = append(admin.Children, desktop, documents, settings)
	
	// Level 3 - Users/Student
	var homework *Node
	var projects *Node
	var downloads *Node
	homework = &Node{Name: "Homework", IsDir: true}
	projects = &Node{Name: "Projects", IsDir: true}
	downloads = &Node{Name: "Downloads", IsDir: true}
	student.Children = append(student.Children, homework, projects, downloads)
	
	// Level 3 - Programs/Games
	var puzzle *Node
	var arcade *Node
	var strategy *Node
	puzzle = &Node{Name: "puzzle_game", IsDir: false}
	arcade = &Node{Name: "arcade_game", IsDir: false}
	strategy = &Node{Name: "strategy_game", IsDir: false}
	games.Children = append(games.Children, puzzle, arcade, strategy)
	
	// Level 3 - Programs/Office
	var textEditor *Node
	var calculator *Node
	var calendar *Node
	textEditor = &Node{Name: "text_editor", IsDir: false}
	calculator = &Node{Name: "calculator", IsDir: false}
	calendar = &Node{Name: "calendar", IsDir: false}
	office.Children = append(office.Children, textEditor, calculator, calendar)
	
	// Level 3 - Programs/Development
	var compiler *Node
	var debugger *Node
	var editor *Node
	compiler = &Node{Name: "compiler", IsDir: false}
	debugger = &Node{Name: "debugger", IsDir: false}
	editor = &Node{Name: "code_editor", IsDir: false}
	development.Children = append(development.Children, compiler, debugger, editor)
	
	// Level 4 - Users/Admin/Desktop
	var shortcut1 *Node
	var shortcut2 *Node
	var readme *Node
	shortcut1 = &Node{Name: "system_monitor", IsDir: false}
	shortcut2 = &Node{Name: "file_manager", IsDir: false}
	readme = &Node{Name: "readme_file", IsDir: false}
	desktop.Children = append(desktop.Children, shortcut1, shortcut2, readme)
	
	// Level 4 - Users/Admin/Documents
	var report *Node
	var backup *Node
	var config *Node
	report = &Node{Name: "system_report", IsDir: false}
	backup = &Node{Name: "backup_config", IsDir: false}
	config = &Node{Name: "user_config", IsDir: false}
	documents.Children = append(documents.Children, report, backup, config)
	
	// Level 4 - Users/Student/Homework
	var math *Node
	var science *Node
	var history *Node
	math = &Node{Name: "math_assignment", IsDir: false}
	science = &Node{Name: "science_project", IsDir: false}
	history = &Node{Name: "history_essay", IsDir: false}
	homework.Children = append(homework.Children, math, science, history)
	
	// Level 4 - Users/Student/Projects
	var webProject *Node
	var gameProject *Node
	var aiProject *Node
	webProject = &Node{Name: "website_project", IsDir: false}
	gameProject = &Node{Name: "game_project", IsDir: false}
	aiProject = &Node{Name: "ai_project", IsDir: false}
	projects.Children = append(projects.Children, webProject, gameProject, aiProject)
	
	// Level 4 - Users/Student/Downloads
	var tutorial *Node
	var software *Node
	var media *Node
	tutorial = &Node{Name: "programming_tutorial", IsDir: false}
	software = &Node{Name: "development_tools", IsDir: false}
	media = &Node{Name: "study_materials", IsDir: false}
	downloads.Children = append(downloads.Children, tutorial, software, media)
	
	// Level 2 - Temp (file temporary)
	var cache *Node
	var tempFile1 *Node
	var tempFile2 *Node
	cache = &Node{Name: "cache_data", IsDir: false}
	tempFile1 = &Node{Name: "temp_session", IsDir: false}
	tempFile2 = &Node{Name: "temp_backup", IsDir: false}
	temp.Children = append(temp.Children, cache, tempFile1, tempFile2)
	
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
