from node import Node


def create_sample_file_system():
    root = Node("PacOS", True)
    
    system = Node("System", True)
    users = Node("Users", True)
    programs = Node("Programs", True)
    temp = Node("Temp", True)
    
    root.children.extend([system, users, programs, temp])
    
    kernel = Node("Kernel", True)
    drivers = Node("Drivers", True)
    logs = Node("Logs", True)
    system.children.extend([kernel, drivers, logs])
    
    admin = Node("Admin", True)
    guest = Node("Guest", True)
    student = Node("Student", True)
    users.children.extend([admin, guest, student])
    
    games = Node("Games", True)
    office = Node("Office", True)
    development = Node("Development", True)
    programs.children.extend([games, office, development])
    
    kernel.children.extend([
        Node("bootloader", False),
        Node("scheduler", False),
        Node("memory_manager", False)
    ])
    
    drivers.children.extend([
        Node("graphics_driver", False),
        Node("audio_driver", False),
        Node("network_driver", False)
    ])
    
    logs.children.extend([
        Node("system_log", False),
        Node("error_log", False),
        Node("access_log", False)
    ])
    
    desktop = Node("Desktop", True)
    documents = Node("Documents", True)
    settings = Node("Settings", True)
    admin.children.extend([desktop, documents, settings])
    
    homework = Node("Homework", True)
    projects = Node("Projects", True)
    downloads = Node("Downloads", True)
    student.children.extend([homework, projects, downloads])
    
    games.children.extend([
        Node("puzzle_game", False),
        Node("arcade_game", False),
        Node("strategy_game", False)
    ])
    
    office.children.extend([
        Node("text_editor", False),
        Node("calculator", False),
        Node("calendar", False)
    ])
    
    development.children.extend([
        Node("compiler", False),
        Node("debugger", False),
        Node("code_editor", False)
    ])
    
    desktop.children.extend([
        Node("system_monitor", False),
        Node("file_manager", False),
        Node("readme_file", False)
    ])
    
    documents.children.extend([
        Node("system_report", False),
        Node("backup_config", False),
        Node("user_config", False)
    ])
    
    homework.children.extend([
        Node("math_assignment", False),
        Node("science_project", False),
        Node("history_essay", False)
    ])
    
    projects.children.extend([
        Node("website_project", False),
        Node("game_project", False),
        Node("ai_project", False)
    ])
    
    downloads.children.extend([
        Node("programming_tutorial", False),
        Node("development_tools", False),
        Node("study_materials", False)
    ])
    
    temp.children.extend([
        Node("cache_data", False),
        Node("temp_session", False),
        Node("temp_backup", False)
    ])
    
    return root


def create_small_file_system():
    root = Node("SmallOS", True)
    
    docs = Node("Documents", True)
    apps = Node("Applications", True)
    root.children.extend([docs, apps])
    
    work = Node("Work", True)
    personal = Node("Personal", True)
    docs.children.extend([work, personal])
    
    utilities = Node("Utilities", True)
    games = Node("Games", True)
    apps.children.extend([utilities, games])
    
    work.children.extend([
        Node("quarterly_report", False),
        Node("annual_report", False),
        Node("client_presentation", False),
        Node("budget_sheet", False),
        Node("meeting_notes", False)
    ])
    
    personal.children.extend([
        Node("vacation_photo", False),
        Node("personal_diary", False),
        Node("favorite_recipes", False)
    ])
    
    utilities.children.extend([
        Node("calculator_app", False),
        Node("text_editor", False),
        Node("file_manager", False)
    ])
    
    games.children.extend([
        Node("puzzle_game", False),
        Node("chess_game", False),
        Node("solitaire_game", False),
        Node("tetris_game", False)
    ])
    
    return root


def create_large_file_system():
    root = Node("LargeOS", True)
    
    system = Node("System", True)
    users = Node("Users", True)
    programs = Node("Programs", True)
    data = Node("Data", True)
    root.children.extend([system, users, programs, data])
    
    kernel = Node("Kernel", True)
    drivers = Node("Drivers", True)
    config = Node("Config", True)
    system.children.extend([kernel, drivers, config])
    
    admin = Node("Admin", True)
    user1 = Node("User1", True)
    user2 = Node("User2", True)
    guest = Node("Guest", True)
    users.children.extend([admin, user1, user2, guest])
    
    office = Node("Office", True)
    development = Node("Development", True)
    multimedia = Node("Multimedia", True)
    internet = Node("Internet", True)
    utilities = Node("Utilities", True)
    programs.children.extend([office, development, multimedia, internet, utilities])
    
    databases = Node("Databases", True)
    backups = Node("Backups", True)
    logs = Node("Logs", True)
    data.children.extend([databases, backups, logs])
    
    kernel.children.extend([
        Node("bootloader", False),
        Node("scheduler", False),
        Node("memory_manager", False),
        Node("process_manager", False),
        Node("interrupt_handler", False),
        Node("syscall_handler", False)
    ])
    
    drivers.children.extend([
        Node("graphics_driver", False),
        Node("audio_driver", False),
        Node("network_driver", False),
        Node("usb_driver", False),
        Node("disk_driver", False),
        Node("keyboard_driver", False),
        Node("mouse_driver", False),
        Node("printer_driver", False)
    ])
    
    config.children.extend([
        Node("system_config", False),
        Node("network_config", False),
        Node("user_config", False),
        Node("security_config", False),
        Node("display_config", False)
    ])
    
    admin_docs = Node("Documents", True)
    admin_projects = Node("Projects", True)
    admin_settings = Node("Settings", True)
    admin.children.extend([admin_docs, admin_projects, admin_settings])
    
    admin_docs.children.extend([
        Node("monthly_report", False),
        Node("quarterly_report", False),
        Node("annual_report", False),
        Node("system_manual", False),
        Node("security_policy", False),
        Node("it_budget", False),
        Node("audit_log", False)
    ])
    
    project1 = Node("WebApp", True)
    project2 = Node("MobileApp", True)
    project3 = Node("Database", True)
    project4 = Node("API", True)
    project5 = Node("Infrastructure", True)
    admin_projects.children.extend([project1, project2, project3, project4, project5])
    
    project1.children.extend([
        Node("frontend_code", False),
        Node("backend_code", False),
        Node("database_schema", False),
        Node("deployment_config", False)
    ])
    
    project2.children.extend([
        Node("android_app", False),
        Node("ios_app", False),
        Node("shared_library", False),
        Node("test_suite", False)
    ])
    
    project3.children.extend([
        Node("schema_design", False),
        Node("migration_scripts", False),
        Node("stored_procedures", False),
        Node("backup_strategy", False)
    ])
    
    office.children.extend([
        Node("word_processor", False),
        Node("spreadsheet", False),
        Node("presentation", False),
        Node("email_client", False),
        Node("calendar_app", False),
        Node("notes_app", False)
    ])
    
    development.children.extend([
        Node("code_editor", False),
        Node("compiler", False),
        Node("debugger", False),
        Node("version_control", False),
        Node("container_tool", False),
        Node("terminal_emulator", False),
        Node("database_client", False)
    ])
    
    databases.children.extend([
        Node("users_database", False),
        Node("products_database", False),
        Node("orders_database", False),
        Node("analytics_database", False),
        Node("cache_database", False)
    ])
    
    backups.children.extend([
        Node("daily_backup", False),
        Node("weekly_backup", False),
        Node("monthly_backup", False),
        Node("system_backup", False),
        Node("user_backup", False),
        Node("database_backup", False)
    ])
    
    logs.children.extend([
        Node("system_log", False),
        Node("error_log", False),
        Node("access_log", False),
        Node("security_log", False),
        Node("application_log", False),
        Node("database_log", False),
        Node("network_log", False),
        Node("performance_log", False)
    ])
    
    return root


def create_flat_file_system():
    root = Node("FlatOS", True)
    
    for i in range(1, 21):
        root.children.append(Node(f"file_{i:03d}", False))
    
    return root


def create_balanced_binary_tree():
    root = Node("BinaryTreeOS", True)
    
    left1 = Node("LeftBranch", True)
    right1 = Node("RightBranch", True)
    root.children.extend([left1, right1])
    
    left2_1 = Node("LeftLeft", True)
    left2_2 = Node("LeftRight", True)
    left1.children.extend([left2_1, left2_2])
    
    right2_1 = Node("RightLeft", True)
    right2_2 = Node("RightRight", True)
    right1.children.extend([right2_1, right2_2])
    
    ll_1 = Node("LL_Left", True)
    ll_2 = Node("LL_Right", True)
    left2_1.children.extend([ll_1, ll_2])
    
    lr_1 = Node("LR_Left", True)
    lr_2 = Node("LR_Right", True)
    left2_2.children.extend([lr_1, lr_2])
    
    rl_1 = Node("RL_Left", True)
    rl_2 = Node("RL_Right", True)
    right2_1.children.extend([rl_1, rl_2])
    
    rr_1 = Node("RR_Left", True)
    rr_2 = Node("RR_Right", True)
    right2_2.children.extend([rr_1, rr_2])
    
    ll_1.children.extend([Node("data_001", False), Node("data_002", False)])
    ll_2.children.extend([Node("data_003", False), Node("data_004", False)])
    lr_1.children.extend([Node("data_005", False), Node("data_006", False)])
    lr_2.children.extend([Node("data_007", False), Node("data_008", False)])
    rl_1.children.extend([Node("data_009", False), Node("data_010", False)])
    rl_2.children.extend([Node("data_011", False), Node("data_012", False)])
    rr_1.children.extend([Node("data_013", False), Node("data_014", False)])
    rr_2.children.extend([Node("data_015", False), Node("data_016", False)])
    
    return root


def create_perfect_binary_tree(depth):
    if depth <= 0:
        return None
    
    root = Node("Root", True)
    node_counter = [0]
    
    def build_tree(parent, current_depth):
        if current_depth >= depth:
            return
        
        for i in range(2):
            node_counter[0] += 1
            
            if current_depth == depth - 1:
                child = Node(f"File_{node_counter[0]}", False)
            else:
                child = Node(f"Dir_{node_counter[0]}", True)
            
            parent.children.append(child)
            
            if child.is_dir:
                build_tree(child, current_depth + 1)
    
    build_tree(root, 0)
    return root


def get_file_system_stats(root):
    if root is None:
        return 0, 0, 0, 0
    
    total_nodes = [0]
    total_dirs = [0]
    total_files = [0]
    
    def count_nodes(node, depth):
        if node is None:
            return depth
        
        total_nodes[0] += 1
        if node.is_dir:
            total_dirs[0] += 1
        else:
            total_files[0] += 1
        
        current_max_depth = depth
        for child in node.children:
            child_depth = count_nodes(child, depth + 1)
            if child_depth > current_max_depth:
                current_max_depth = child_depth
        
        return current_max_depth
    
    max_depth = count_nodes(root, 1)
    return total_nodes[0], total_dirs[0], total_files[0], max_depth


def is_perfect_binary_tree(root):
    if root is None:
        return True
    
    _, _, _, max_depth = get_file_system_stats(root)
    
    def check_perfect(node, depth, level):
        if node is None:
            return True
        
        if len(node.children) == 0:
            return depth == level
        
        if len(node.children) != 2:
            return False
        
        return (check_perfect(node.children[0], depth, level + 1) and
                check_perfect(node.children[1], depth, level + 1))
    
    return check_perfect(root, max_depth, 1)


def print_file_system_stats(root):
    total_nodes, total_dirs, total_files, max_depth = get_file_system_stats(root)
    
    print("\n=== STATISTIK FILE SYSTEM ===")
    print(f"Nama Root       : {root.name}")
    print(f"Total Nodes     : {total_nodes}")
    print(f"Total Direktori : {total_dirs}")
    print(f"Total Files     : {total_files}")
    print(f"Kedalaman Max   : {max_depth} level")
    
    if is_perfect_binary_tree(root):
        print("Struktur        : Perfect Binary Tree")
        expected_nodes = (1 << max_depth) - 1
        print(f"Nodes (teoritis): {expected_nodes} (2^{max_depth} - 1)")
    
    print()
