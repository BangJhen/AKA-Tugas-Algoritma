def dfs_iterative(flat_stack, target):
    for i in range(len(flat_stack)):
        if flat_stack[i].name == target:
            return flat_stack[i].path, True
    
    return "", False
