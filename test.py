from node import flatten_tree_to_stack
from dfs_recursive import dfs_recursive
from dfs_iterative import dfs_iterative
from sample_filesystems import create_sample_file_system, print_file_system_stats


print("Test Program File System PacOS")
print("=" * 40)

root = create_sample_file_system()
print_file_system_stats(root)

flat_stack = flatten_tree_to_stack(root)
print(f"Total nodes di flat stack: {len(flat_stack)}")
print()

test_files = ["bootloader", "calculator", "ai_project", "tidak_ada"]

print("Test DFS Recursive:")
print("-" * 40)
for target in test_files:
    path, found = dfs_recursive(flat_stack, target, 0)
    if found:
        print(f"[OK] {target} -> {path}")
    else:
        print(f"[X]  {target} -> Tidak ditemukan")

print()
print("Test DFS Iterative:")
print("-" * 40)
for target in test_files:
    path, found = dfs_iterative(flat_stack, target)
    if found:
        print(f"[OK] {target} -> {path}")
    else:
        print(f"[X]  {target} -> Tidak ditemukan")

print()
print("Test selesai!")
