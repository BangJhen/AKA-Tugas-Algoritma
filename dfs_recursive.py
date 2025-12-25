def dfs_recursive(flat_stack, target, index=0):
    if index >= len(flat_stack):
        return "", False
    
    if flat_stack[index].name == target:
        return flat_stack[index].path, True
    else:
        return dfs_recursive(flat_stack, target, index + 1)
