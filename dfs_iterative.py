def dfs_iterative(flat_stack, target):
    found = False
    i = 0
    while (i < len(flat_stack) and not found):
        if flat_stack[i].name == target:
            found = True
        i += 1
    
    return flat_stack[i].path, found