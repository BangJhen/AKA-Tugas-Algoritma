class Node:
    def __init__(self, name, is_dir=False):
        self.name = name
        self.is_dir = is_dir
        self.children = []


class FlatNode:
    def __init__(self, name, path, is_dir):
        self.name = name
        self.path = path
        self.is_dir = is_dir


def flatten_tree_to_stack(root):
    if root is None:
        return []
    
    flat_list = []
    _flatten_helper(root, root.name, flat_list)
    return flat_list


def _flatten_helper(node, current_path, flat_list):
    if node is None:
        return
    
    flat_list.append(FlatNode(node.name, current_path, node.is_dir))
    
    if node.is_dir:
        for child in node.children:
            new_path = current_path + "/" + child.name
            _flatten_helper(child, new_path, flat_list)
