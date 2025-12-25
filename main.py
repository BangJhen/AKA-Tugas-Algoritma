from node import flatten_tree_to_stack
from dfs_recursive import dfs_recursive
from dfs_iterative import dfs_iterative
from sample_filesystems import (
    create_sample_file_system,
    create_small_file_system,
    create_large_file_system,
    create_flat_file_system,
    create_balanced_binary_tree,
    create_perfect_binary_tree,
    print_file_system_stats
)


def print_file_system(node, level=0):
    if node is None:
        return
    
    print("  " * level, end="")
    
    if node.is_dir:
        print(f"[{node.name}/]")
    else:
        print(node.name)
    
    for child in node.children:
        print_file_system(child, level + 1)


def main():
    user_root = None
    pilih = 0
    
    print("=== PILIH FILE SYSTEM ===")
    print("1. PacOS (Default - Medium, ~50 nodes)")
    print("2. SmallOS (Small - ~20 nodes)")
    print("3. LargeOS (Large - ~100+ nodes)")
    print("4. FlatOS (2 Level - Root -> Files)")
    print("5. BinaryTreeOS (Perfect Binary Tree - 31 nodes)")
    print("6. Custom Binary Tree (Depth 3-7)")
    
    try:
        fs_choice = int(input("Pilih: "))
    except:
        fs_choice = 1
    
    if fs_choice == 2:
        user_root = create_small_file_system()
    elif fs_choice == 3:
        user_root = create_large_file_system()
    elif fs_choice == 4:
        user_root = create_flat_file_system()
    elif fs_choice == 5:
        user_root = create_balanced_binary_tree()
    elif fs_choice == 6:
        try:
            depth = int(input("Masukkan kedalaman tree (3-7): "))
            if depth < 3:
                depth = 3
            elif depth > 7:
                depth = 7
        except:
            depth = 3
        user_root = create_perfect_binary_tree(depth)
    else:
        user_root = create_sample_file_system()
    
    print_file_system_stats(user_root)
    
    flat_stack = flatten_tree_to_stack(user_root)
    
    while pilih != 4:
        print("=== SISTEM PENCARIAN FILE SYSTEM DENGAN DFS ===")
        print("1. Tampilkan Diretory")
        print("2. Pencarian dengan DFS Rekursif (Linear Recursion)")
        print("3. Pencarian dengan DFS Iteratif (Linear Loop)")
        print("4. Exit")
        
        try:
            pilih = int(input("Pilih: "))
        except:
            pilih = 0
        
        if pilih == 1:
            print("Struktur File System:")
            print_file_system(user_root, 0)
            print()
        
        elif pilih == 2:
            target = input("Masukkan nama file/direktori yang dicari: ")
            path, found = dfs_recursive(flat_stack, target, 0)
            
            if found:
                print(f"File/Directory Ditemukan! Path: {path}")
            else:
                print(f"File {target} tidak ditemukan.")
            print()
        
        elif pilih == 3:
            target = input("Masukkan nama file/direktori yang dicari: ")
            path, found = dfs_iterative(flat_stack, target)
            
            if found:
                print(f"File/Directory Ditemukan! Path: {path}")
            else:
                print(f"File {target} tidak ditemukan.")
            print()
        
        elif pilih == 4:
            print("Terima kasih!")
        
        else:
            print("Pilihan tidak valid!")


if __name__ == "__main__":
    main()
