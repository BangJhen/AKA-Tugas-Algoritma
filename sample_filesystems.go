package main

import "fmt"

// createSampleFileSystem membuat file system default PacOS dengan 4 level dan ~50 nodes
// Cocok untuk demonstrasi standar dan testing umum
func createSampleFileSystem() *Node {
	root := &Node{Name: "PacOS", IsDir: true}

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

// createSmallFileSystem membuat file system kecil dengan 3 level dan ~20 nodes
// Cocok untuk testing cepat dan debugging
func createSmallFileSystem() *Node {
	root := &Node{Name: "SmallOS", IsDir: true}

	// Level 1 - 2 direktori utama
	var docs *Node
	var apps *Node
	docs = &Node{Name: "Documents", IsDir: true}
	apps = &Node{Name: "Applications", IsDir: true}

	root.Children = append(root.Children, docs, apps)

	// Level 2 - Documents
	var work *Node
	var personal *Node
	work = &Node{Name: "Work", IsDir: true}
	personal = &Node{Name: "Personal", IsDir: true}
	docs.Children = append(docs.Children, work, personal)

	// Level 2 - Applications
	var utilities *Node
	var games *Node
	utilities = &Node{Name: "Utilities", IsDir: true}
	games = &Node{Name: "Games", IsDir: true}
	apps.Children = append(apps.Children, utilities, games)

	// Level 3 - Work files (5 files)
	var report1 *Node
	var report2 *Node
	var presentation *Node
	var spreadsheet *Node
	var notes *Node
	report1 = &Node{Name: "quarterly_report", IsDir: false}
	report2 = &Node{Name: "annual_report", IsDir: false}
	presentation = &Node{Name: "client_presentation", IsDir: false}
	spreadsheet = &Node{Name: "budget_sheet", IsDir: false}
	notes = &Node{Name: "meeting_notes", IsDir: false}
	work.Children = append(work.Children, report1, report2, presentation, spreadsheet, notes)

	// Level 3 - Personal files (3 files)
	var photo *Node
	var diary *Node
	var recipe *Node
	photo = &Node{Name: "vacation_photo", IsDir: false}
	diary = &Node{Name: "personal_diary", IsDir: false}
	recipe = &Node{Name: "favorite_recipes", IsDir: false}
	personal.Children = append(personal.Children, photo, diary, recipe)

	// Level 3 - Utilities files (3 files)
	var calculator *Node
	var textEditor *Node
	var fileManager *Node
	calculator = &Node{Name: "calculator_app", IsDir: false}
	textEditor = &Node{Name: "text_editor", IsDir: false}
	fileManager = &Node{Name: "file_manager", IsDir: false}
	utilities.Children = append(utilities.Children, calculator, textEditor, fileManager)

	// Level 3 - Games files (4 files)
	var puzzle *Node
	var chess *Node
	var solitaire *Node
	var tetris *Node
	puzzle = &Node{Name: "puzzle_game", IsDir: false}
	chess = &Node{Name: "chess_game", IsDir: false}
	solitaire = &Node{Name: "solitaire_game", IsDir: false}
	tetris = &Node{Name: "tetris_game", IsDir: false}
	games.Children = append(games.Children, puzzle, chess, solitaire, tetris)

	return root
}

// createLargeFileSystem membuat file system besar dengan 5 level dan ~100+ nodes
// Cocok untuk stress testing dan benchmark performa
func createLargeFileSystem() *Node {
	root := &Node{Name: "LargeOS", IsDir: true}

	// Level 1 - 4 direktori utama
	var system *Node
	var users *Node
	var programs *Node
	var data *Node
	system = &Node{Name: "System", IsDir: true}
	users = &Node{Name: "Users", IsDir: true}
	programs = &Node{Name: "Programs", IsDir: true}
	data = &Node{Name: "Data", IsDir: true}

	root.Children = append(root.Children, system, users, programs, data)

	// Level 2 - System (3 subdirs)
	var kernel *Node
	var drivers *Node
	var config *Node
	kernel = &Node{Name: "Kernel", IsDir: true}
	drivers = &Node{Name: "Drivers", IsDir: true}
	config = &Node{Name: "Config", IsDir: true}
	system.Children = append(system.Children, kernel, drivers, config)

	// Level 2 - Users (4 subdirs)
	var admin *Node
	var user1 *Node
	var user2 *Node
	var guest *Node
	admin = &Node{Name: "Admin", IsDir: true}
	user1 = &Node{Name: "User1", IsDir: true}
	user2 = &Node{Name: "User2", IsDir: true}
	guest = &Node{Name: "Guest", IsDir: true}
	users.Children = append(users.Children, admin, user1, user2, guest)

	// Level 2 - Programs (5 subdirs)
	var office *Node
	var development *Node
	var multimedia *Node
	var internet *Node
	var utilities *Node
	office = &Node{Name: "Office", IsDir: true}
	development = &Node{Name: "Development", IsDir: true}
	multimedia = &Node{Name: "Multimedia", IsDir: true}
	internet = &Node{Name: "Internet", IsDir: true}
	utilities = &Node{Name: "Utilities", IsDir: true}
	programs.Children = append(programs.Children, office, development, multimedia, internet, utilities)

	// Level 2 - Data (3 subdirs)
	var databases *Node
	var backups *Node
	var logs *Node
	databases = &Node{Name: "Databases", IsDir: true}
	backups = &Node{Name: "Backups", IsDir: true}
	logs = &Node{Name: "Logs", IsDir: true}
	data.Children = append(data.Children, databases, backups, logs)

	// Level 3 - System/Kernel (6 files)
	var boot *Node
	var scheduler *Node
	var memory *Node
	var process *Node
	var interrupt *Node
	var syscall *Node
	boot = &Node{Name: "bootloader", IsDir: false}
	scheduler = &Node{Name: "scheduler", IsDir: false}
	memory = &Node{Name: "memory_manager", IsDir: false}
	process = &Node{Name: "process_manager", IsDir: false}
	interrupt = &Node{Name: "interrupt_handler", IsDir: false}
	syscall = &Node{Name: "syscall_handler", IsDir: false}
	kernel.Children = append(kernel.Children, boot, scheduler, memory, process, interrupt, syscall)

	// Level 3 - System/Drivers (8 files)
	var graphics *Node
	var audio *Node
	var network *Node
	var usb *Node
	var disk *Node
	var keyboard *Node
	var mouse *Node
	var printer *Node
	graphics = &Node{Name: "graphics_driver", IsDir: false}
	audio = &Node{Name: "audio_driver", IsDir: false}
	network = &Node{Name: "network_driver", IsDir: false}
	usb = &Node{Name: "usb_driver", IsDir: false}
	disk = &Node{Name: "disk_driver", IsDir: false}
	keyboard = &Node{Name: "keyboard_driver", IsDir: false}
	mouse = &Node{Name: "mouse_driver", IsDir: false}
	printer = &Node{Name: "printer_driver", IsDir: false}
	drivers.Children = append(drivers.Children, graphics, audio, network, usb, disk, keyboard, mouse, printer)

	// Level 3 - System/Config (5 files)
	var systemConf *Node
	var networkConf *Node
	var userConf *Node
	var securityConf *Node
	var displayConf *Node
	systemConf = &Node{Name: "system_config", IsDir: false}
	networkConf = &Node{Name: "network_config", IsDir: false}
	userConf = &Node{Name: "user_config", IsDir: false}
	securityConf = &Node{Name: "security_config", IsDir: false}
	displayConf = &Node{Name: "display_config", IsDir: false}
	config.Children = append(config.Children, systemConf, networkConf, userConf, securityConf, displayConf)

	// Level 3 - Users/Admin (3 subdirs)
	var adminDocs *Node
	var adminProjects *Node
	var adminSettings *Node
	adminDocs = &Node{Name: "Documents", IsDir: true}
	adminProjects = &Node{Name: "Projects", IsDir: true}
	adminSettings = &Node{Name: "Settings", IsDir: true}
	admin.Children = append(admin.Children, adminDocs, adminProjects, adminSettings)

	// Level 4 - Admin/Documents (7 files)
	var report1 *Node
	var report2 *Node
	var report3 *Node
	var manual *Node
	var policy *Node
	var budget *Node
	var audit *Node
	report1 = &Node{Name: "monthly_report", IsDir: false}
	report2 = &Node{Name: "quarterly_report", IsDir: false}
	report3 = &Node{Name: "annual_report", IsDir: false}
	manual = &Node{Name: "system_manual", IsDir: false}
	policy = &Node{Name: "security_policy", IsDir: false}
	budget = &Node{Name: "it_budget", IsDir: false}
	audit = &Node{Name: "audit_log", IsDir: false}
	adminDocs.Children = append(adminDocs.Children, report1, report2, report3, manual, policy, budget, audit)

	// Level 4 - Admin/Projects (5 subdirs untuk depth)
	var project1 *Node
	var project2 *Node
	var project3 *Node
	var project4 *Node
	var project5 *Node
	project1 = &Node{Name: "WebApp", IsDir: true}
	project2 = &Node{Name: "MobileApp", IsDir: true}
	project3 = &Node{Name: "Database", IsDir: true}
	project4 = &Node{Name: "API", IsDir: true}
	project5 = &Node{Name: "Infrastructure", IsDir: true}
	adminProjects.Children = append(adminProjects.Children, project1, project2, project3, project4, project5)

	// Level 5 - Project files (4 files each)
	var web1, web2, web3, web4 *Node
	web1 = &Node{Name: "frontend_code", IsDir: false}
	web2 = &Node{Name: "backend_code", IsDir: false}
	web3 = &Node{Name: "database_schema", IsDir: false}
	web4 = &Node{Name: "deployment_config", IsDir: false}
	project1.Children = append(project1.Children, web1, web2, web3, web4)

	var mobile1, mobile2, mobile3, mobile4 *Node
	mobile1 = &Node{Name: "android_app", IsDir: false}
	mobile2 = &Node{Name: "ios_app", IsDir: false}
	mobile3 = &Node{Name: "shared_library", IsDir: false}
	mobile4 = &Node{Name: "test_suite", IsDir: false}
	project2.Children = append(project2.Children, mobile1, mobile2, mobile3, mobile4)

	var db1, db2, db3, db4 *Node
	db1 = &Node{Name: "schema_design", IsDir: false}
	db2 = &Node{Name: "migration_scripts", IsDir: false}
	db3 = &Node{Name: "stored_procedures", IsDir: false}
	db4 = &Node{Name: "backup_strategy", IsDir: false}
	project3.Children = append(project3.Children, db1, db2, db3, db4)

	// Level 3 - Programs/Office (6 files)
	var word *Node
	var excel *Node
	var powerpoint *Node
	var email *Node
	var calendar *Node
	var notes *Node
	word = &Node{Name: "word_processor", IsDir: false}
	excel = &Node{Name: "spreadsheet", IsDir: false}
	powerpoint = &Node{Name: "presentation", IsDir: false}
	email = &Node{Name: "email_client", IsDir: false}
	calendar = &Node{Name: "calendar_app", IsDir: false}
	notes = &Node{Name: "notes_app", IsDir: false}
	office.Children = append(office.Children, word, excel, powerpoint, email, calendar, notes)

	// Level 3 - Programs/Development (7 files)
	var ide *Node
	var compiler *Node
	var debugger *Node
	var git *Node
	var docker *Node
	var terminal *Node
	var database *Node
	ide = &Node{Name: "code_editor", IsDir: false}
	compiler = &Node{Name: "compiler", IsDir: false}
	debugger = &Node{Name: "debugger", IsDir: false}
	git = &Node{Name: "version_control", IsDir: false}
	docker = &Node{Name: "container_tool", IsDir: false}
	terminal = &Node{Name: "terminal_emulator", IsDir: false}
	database = &Node{Name: "database_client", IsDir: false}
	development.Children = append(development.Children, ide, compiler, debugger, git, docker, terminal, database)

	// Level 3 - Data/Databases (5 files)
	var userDB *Node
	var productDB *Node
	var orderDB *Node
	var analyticsDB *Node
	var cacheDB *Node
	userDB = &Node{Name: "users_database", IsDir: false}
	productDB = &Node{Name: "products_database", IsDir: false}
	orderDB = &Node{Name: "orders_database", IsDir: false}
	analyticsDB = &Node{Name: "analytics_database", IsDir: false}
	cacheDB = &Node{Name: "cache_database", IsDir: false}
	databases.Children = append(databases.Children, userDB, productDB, orderDB, analyticsDB, cacheDB)

	// Level 3 - Data/Backups (6 files)
	var backup1 *Node
	var backup2 *Node
	var backup3 *Node
	var backup4 *Node
	var backup5 *Node
	var backup6 *Node
	backup1 = &Node{Name: "daily_backup", IsDir: false}
	backup2 = &Node{Name: "weekly_backup", IsDir: false}
	backup3 = &Node{Name: "monthly_backup", IsDir: false}
	backup4 = &Node{Name: "system_backup", IsDir: false}
	backup5 = &Node{Name: "user_backup", IsDir: false}
	backup6 = &Node{Name: "database_backup", IsDir: false}
	backups.Children = append(backups.Children, backup1, backup2, backup3, backup4, backup5, backup6)

	// Level 3 - Data/Logs (8 files)
	var sysLog *Node
	var errLog *Node
	var accessLog *Node
	var secLog *Node
	var appLog *Node
	var dbLog *Node
	var netLog *Node
	var perfLog *Node
	sysLog = &Node{Name: "system_log", IsDir: false}
	errLog = &Node{Name: "error_log", IsDir: false}
	accessLog = &Node{Name: "access_log", IsDir: false}
	secLog = &Node{Name: "security_log", IsDir: false}
	appLog = &Node{Name: "application_log", IsDir: false}
	dbLog = &Node{Name: "database_log", IsDir: false}
	netLog = &Node{Name: "network_log", IsDir: false}
	perfLog = &Node{Name: "performance_log", IsDir: false}
	logs.Children = append(logs.Children, sysLog, errLog, accessLog, secLog, appLog, dbLog, netLog, perfLog)

	return root
}

// createPerfectBinaryTree membuat file system dengan struktur perfect binary tree
func createPerfectBinaryTree(depth int) *Node {
	if depth <= 0 {
		return nil
	}

	root := &Node{Name: "Root", IsDir: true}

	// Helper function untuk membuat tree secara rekursif
	var buildTree func(*Node, int, int) int
	buildTree = func(parent *Node, currentDepth int, nodeCounter int) int {
		if currentDepth >= depth {
			return nodeCounter
		}

		// Buat 2 children untuk setiap parent
		var i int
		for i = 0; i < 2; i++ {
			nodeCounter++
			var child *Node

			// Level terakhir adalah file, level lainnya adalah direktori
			if currentDepth == depth-1 {
				child = &Node{
					Name:  fmt.Sprintf("File_%d", nodeCounter),
					IsDir: false,
				}
			} else {
				child = &Node{
					Name:     fmt.Sprintf("Dir_%d", nodeCounter),
					IsDir:    true,
					Children: []*Node{},
				}
			}

			parent.Children = append(parent.Children, child)

			// Rekursi untuk membuat subtree
			if child.IsDir {
				nodeCounter = buildTree(child, currentDepth+1, nodeCounter)
			}
		}

		return nodeCounter
	}

	buildTree(root, 0, 0)
	return root
}

// getFileSystemStats menghitung statistik file system
func getFileSystemStats(root *Node) (totalNodes, totalDirs, totalFiles, maxDepth int) {
	if root == nil {
		return 0, 0, 0, 0
	}

	var countNodes func(*Node, int) int
	countNodes = func(node *Node, depth int) int {
		if node == nil {
			return depth
		}

		totalNodes++
		if node.IsDir {
			totalDirs++
		} else {
			totalFiles++
		}

		currentMaxDepth := depth
		var i int
		for i = 0; i < len(node.Children); i++ {
			childDepth := countNodes(node.Children[i], depth+1)
			if childDepth > currentMaxDepth {
				currentMaxDepth = childDepth
			}
		}

		return currentMaxDepth
	}

	maxDepth = countNodes(root, 1)
	return
}

// createFlatFileSystem membuat file system dengan 2 level (root → files)
// Tidak ada subdirektori, semua file langsung di bawah root
func createFlatFileSystem() *Node {
	root := &Node{Name: "FlatOS", IsDir: true}

	// Langsung tambahkan files ke root (total 20 files)
	var i int
	for i = 1; i <= 20; i++ {
		file := &Node{
			Name:  fmt.Sprintf("file_%03d", i),
			IsDir: false,
		}
		root.Children = append(root.Children, file)
	}

	return root
}

// createBalancedBinaryTree membuat binary tree dengan nama yang lebih deskriptif
func createBalancedBinaryTree() *Node {
	root := &Node{Name: "BinaryTreeOS", IsDir: true}

	// Level 1: 2 direktori
	var left1 *Node
	var right1 *Node
	left1 = &Node{Name: "LeftBranch", IsDir: true}
	right1 = &Node{Name: "RightBranch", IsDir: true}
	root.Children = append(root.Children, left1, right1)

	// Level 2: 4 direktori (2 per parent)
	var left2_1, left2_2 *Node
	left2_1 = &Node{Name: "LeftLeft", IsDir: true}
	left2_2 = &Node{Name: "LeftRight", IsDir: true}
	left1.Children = append(left1.Children, left2_1, left2_2)

	var right2_1, right2_2 *Node
	right2_1 = &Node{Name: "RightLeft", IsDir: true}
	right2_2 = &Node{Name: "RightRight", IsDir: true}
	right1.Children = append(right1.Children, right2_1, right2_2)

	// Level 3: 8 direktori (2 per parent)
	var ll_1, ll_2 *Node
	ll_1 = &Node{Name: "LL_Left", IsDir: true}
	ll_2 = &Node{Name: "LL_Right", IsDir: true}
	left2_1.Children = append(left2_1.Children, ll_1, ll_2)

	var lr_1, lr_2 *Node
	lr_1 = &Node{Name: "LR_Left", IsDir: true}
	lr_2 = &Node{Name: "LR_Right", IsDir: true}
	left2_2.Children = append(left2_2.Children, lr_1, lr_2)

	var rl_1, rl_2 *Node
	rl_1 = &Node{Name: "RL_Left", IsDir: true}
	rl_2 = &Node{Name: "RL_Right", IsDir: true}
	right2_1.Children = append(right2_1.Children, rl_1, rl_2)

	var rr_1, rr_2 *Node
	rr_1 = &Node{Name: "RR_Left", IsDir: true}
	rr_2 = &Node{Name: "RR_Right", IsDir: true}
	right2_2.Children = append(right2_2.Children, rr_1, rr_2)

	// Level 4: 16 files (2 per parent)
	// LL_Left
	ll_1.Children = append(ll_1.Children,
		&Node{Name: "data_001", IsDir: false},
		&Node{Name: "data_002", IsDir: false})

	// LL_Right
	ll_2.Children = append(ll_2.Children,
		&Node{Name: "data_003", IsDir: false},
		&Node{Name: "data_004", IsDir: false})

	// LR_Left
	lr_1.Children = append(lr_1.Children,
		&Node{Name: "data_005", IsDir: false},
		&Node{Name: "data_006", IsDir: false})

	// LR_Right
	lr_2.Children = append(lr_2.Children,
		&Node{Name: "data_007", IsDir: false},
		&Node{Name: "data_008", IsDir: false})

	// RL_Left
	rl_1.Children = append(rl_1.Children,
		&Node{Name: "data_009", IsDir: false},
		&Node{Name: "data_010", IsDir: false})

	// RL_Right
	rl_2.Children = append(rl_2.Children,
		&Node{Name: "data_011", IsDir: false},
		&Node{Name: "data_012", IsDir: false})

	// RR_Left
	rr_1.Children = append(rr_1.Children,
		&Node{Name: "data_013", IsDir: false},
		&Node{Name: "data_014", IsDir: false})

	// RR_Right
	rr_2.Children = append(rr_2.Children,
		&Node{Name: "data_015", IsDir: false},
		&Node{Name: "data_016", IsDir: false})

	return root
}

// printFileSystemStats menampilkan statistik file system
func printFileSystemStats(root *Node) {
	totalNodes, totalDirs, totalFiles, maxDepth := getFileSystemStats(root)

	fmt.Println("\n=== STATISTIK FILE SYSTEM ===")
	fmt.Printf("Nama Root       : %s\n", root.Name)
	fmt.Printf("Total Nodes     : %d\n", totalNodes)
	fmt.Printf("Total Direktori : %d\n", totalDirs)
	fmt.Printf("Total Files     : %d\n", totalFiles)
	fmt.Printf("Kedalaman Max   : %d level\n", maxDepth)

	// Cek apakah perfect binary tree
	if isPerfectBinaryTree(root) {
		fmt.Println("Struktur        : Perfect Binary Tree ")
		expectedNodes := (1 << maxDepth) - 1 // 2^depth - 1
		fmt.Printf("Nodes (teoritis): %d (2^%d - 1)\n", expectedNodes, maxDepth)
	}

	fmt.Println()
}

// isPerfectBinaryTree mengecek apakah tree adalah perfect binary tree
func isPerfectBinaryTree(root *Node) bool {
	if root == nil {
		return true
	}

	// Helper function untuk cek perfect binary tree
	var checkPerfect func(*Node, int, int) bool
	checkPerfect = func(node *Node, depth int, level int) bool {
		if node == nil {
			return true
		}

		// Jika leaf node, cek apakah di level yang tepat
		if len(node.Children) == 0 {
			return depth == level
		}

		// Jika bukan leaf, harus punya tepat 2 children
		if len(node.Children) != 2 {
			return false
		}

		// Cek kedua subtree
		return checkPerfect(node.Children[0], depth, level+1) &&
			checkPerfect(node.Children[1], depth, level+1)
	}

	// Hitung depth
	_, _, _, maxDepth := getFileSystemStats(root)

	return checkPerfect(root, maxDepth, 1)
}
