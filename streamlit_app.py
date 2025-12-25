import streamlit as st
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
    get_file_system_stats
)


def get_file_icon(filename):
    """Dapatkan icon berdasarkan ekstensi file"""
    if '.' not in filename:
        return "📄"
    
    ext = filename.split('.')[-1].lower()
    
    icons = {
        'txt': '📝',
        'pdf': '📕',
        'doc': '📘',
        'docx': '📘',
        'xls': '📊',
        'xlsx': '📊',
        'ppt': '📙',
        'pptx': '📙',
        'jpg': '🖼',
        'jpeg': '🖼',
        'png': '🖼',
        'gif': '🖼',
        'mp3': '🎵',
        'wav': '🎵',
        'mp4': '🎬',
        'avi': '🎬',
        'zip': '🗜',
        'rar': '🗜',
        'py': '🐍',
        'js': '📜',
        'html': '🌐',
        'css': '🎨',
        'json': '📋',
        'xml': '📋',
        'log': '📊',
        'md': '📝',
        'go': '🔵',
        'java': '☕',
        'cpp': '⚙',
        'c': '⚙',
        'sh': '💻',
        'exe': '⚙',
        'dll': '⚙',
        'iso': '💿',
    }
    
    return icons.get(ext, '📄')


def render_interactive_tree(node, path="", search_result=None, parent_key="root"):
    """Render tree interaktif dengan expander yang bisa di-klik"""
    if node is None:
        return
    
    current_path = f"{path}/{node.name}" if path else node.name
    is_highlighted = bool(search_result and node.name == search_result)
    
    if node.is_dir:
        folder_icon = "📂" if is_highlighted else "📁"
        
        child_count = len(node.children)
        folder_info = f"({child_count} item{'s' if child_count != 1 else ''})"
        
        expander_key = f"{parent_key}_{node.name}"
        with st.expander(f"{folder_icon} **{node.name}** {folder_info}", expanded=is_highlighted):
            if child_count == 0:
                st.caption("Folder kosong")
            else:
                for idx, child in enumerate(node.children):
                    render_interactive_tree(child, current_path, search_result, f"{expander_key}_{idx}")
    else:
        file_icon = get_file_icon(node.name)
        
        if is_highlighted:
            st.markdown(
                f"<div style='background-color: #fff3cd; padding: 6px 10px; border-radius: 6px; "
                f"margin: 2px 0; border-left: 3px solid #ffc107;'>"
                f"{file_icon} <strong>{node.name}</strong>"
                f"</div>",
                unsafe_allow_html=True
            )
        else:
            st.markdown(f"{file_icon} {node.name}")


def render_compact_tree(node, level=0, search_result=None):
    """Render tree dalam format compact dengan indentasi"""
    if node is None:
        return
    
    indent = "  " * level
    is_highlighted = bool(search_result and node.name == search_result)
    
    if node.is_dir:
        folder_icon = "📂" if is_highlighted else "📁"
        
        if is_highlighted:
            st.markdown(
                f"<div style='background-color: #fff3cd; padding: 4px 8px; border-radius: 4px; "
                f"margin: 2px 0; font-weight: bold;'>{indent}{folder_icon} {node.name}/</div>",
                unsafe_allow_html=True
            )
        else:
            st.markdown(f"{indent}{folder_icon} **{node.name}/**")
        
        for child in node.children:
            render_compact_tree(child, level + 1, search_result)
    else:
        file_icon = get_file_icon(node.name)
        
        if is_highlighted:
            st.markdown(
                f"<div style='background-color: #fff3cd; padding: 4px 8px; border-radius: 4px; "
                f"margin: 2px 0; font-weight: bold; border-left: 3px solid #ffc107;'>"
                f"{indent}{file_icon} {node.name}</div>",
                unsafe_allow_html=True
            )
        else:
            st.markdown(f"{indent}{file_icon} {node.name}")


def init_session_state():
    """Initialize session state variables"""
    if 'file_system' not in st.session_state:
        st.session_state.file_system = None
    if 'flat_stack' not in st.session_state:
        st.session_state.flat_stack = None
    if 'search_result' not in st.session_state:
        st.session_state.search_result = None
    if 'page' not in st.session_state:
        st.session_state.page = 'select'


def select_file_system_page():
    """Page untuk memilih file system"""
    st.title("File System Explorer")
    st.markdown("---")
    
    st.subheader("Pilih File System")
    
    col1, col2 = st.columns(2)
    
    with col1:
        st.markdown("### File System Kecil - Menengah")
        
        if st.button("SmallOS", use_container_width=True, type="secondary"):
            st.session_state.file_system = create_small_file_system()
            st.session_state.flat_stack = flatten_tree_to_stack(st.session_state.file_system)
            st.session_state.page = 'explorer'
            st.rerun()
        st.caption("~ 20 nodes, 3 levels")
        
        if st.button("PacOS (Default)", use_container_width=True, type="primary"):
            st.session_state.file_system = create_sample_file_system()
            st.session_state.flat_stack = flatten_tree_to_stack(st.session_state.file_system)
            st.session_state.page = 'explorer'
            st.rerun()
        st.caption("~ 50 nodes, 4 levels")
        
        if st.button("FlatOS", use_container_width=True, type="secondary"):
            st.session_state.file_system = create_flat_file_system()
            st.session_state.flat_stack = flatten_tree_to_stack(st.session_state.file_system)
            st.session_state.page = 'explorer'
            st.rerun()
        st.caption("2 levels, root dan files")
    
    with col2:
        st.markdown("### File System Besar")
        
        if st.button("LargeOS", use_container_width=True, type="secondary"):
            st.session_state.file_system = create_large_file_system()
            st.session_state.flat_stack = flatten_tree_to_stack(st.session_state.file_system)
            st.session_state.page = 'explorer'
            st.rerun()
        st.caption("~ 100+ nodes, 5 levels")
        
        if st.button("BinaryTreeOS", use_container_width=True, type="secondary"):
            st.session_state.file_system = create_balanced_binary_tree()
            st.session_state.flat_stack = flatten_tree_to_stack(st.session_state.file_system)
            st.session_state.page = 'explorer'
            st.rerun()
        st.caption("Perfect Binary Tree, 31 nodes")
        
        st.markdown("### Custom Binary Tree")
        depth = st.slider("Kedalaman Tree", 3, 7, 4)
        if st.button("Buat Custom Tree", use_container_width=True, type="secondary"):
            st.session_state.file_system = create_perfect_binary_tree(depth)
            st.session_state.flat_stack = flatten_tree_to_stack(st.session_state.file_system)
            st.session_state.page = 'explorer'
            st.rerun()


def explorer_page():
    """Main explorer page dengan visualisasi tree dan search"""
    root = st.session_state.file_system
    flat_stack = st.session_state.flat_stack
    
    with st.sidebar:
        st.title("Pencarian File")
        
        if st.button("Ganti File System", use_container_width=True):
            st.session_state.page = 'select'
            st.session_state.search_result = None
            st.rerun()
        
        st.markdown("---")
        
        search_query = st.text_input("Cari file/folder", placeholder="Masukkan nama...")
        
        algorithm = st.radio(
            "Algoritma Pencarian",
            ["DFS Iteratif", "DFS Rekursif"],
            help="Pilih algoritma untuk mencari file"
        )
        
        if st.button("Cari", use_container_width=True, type="primary", disabled=not search_query):
            if search_query:
                if algorithm == "DFS Rekursif":
                    path, found = dfs_recursive(flat_stack, search_query, 0)
                else:
                    path, found = dfs_iterative(flat_stack, search_query)
                
                st.session_state.search_result = {
                    'query': search_query,
                    'path': path,
                    'found': found,
                    'algorithm': algorithm
                }
            else:
                st.session_state.search_result = None
        
        if st.button("Reset Pencarian", use_container_width=True):
            st.session_state.search_result = None
            st.rerun()
        
        st.markdown("---")
        
        st.subheader("Statistik")
        total_nodes, total_dirs, total_files, max_depth = get_file_system_stats(root)
        
        st.metric("Total Nodes", total_nodes)
        st.metric("Direktori", total_dirs)
        st.metric("Files", total_files)
        st.metric("Kedalaman", f"{max_depth} level")
    
    st.title(f"{root.name} - File System Explorer")
    
    if st.session_state.search_result:
        result = st.session_state.search_result
        
        if result['found']:
            st.success(f"File ditemukan menggunakan {result['algorithm']}!")
            st.info(f"Path: `{result['path']}`")
        else:
            st.error(f"File '{result['query']}' tidak ditemukan.")
    
    st.markdown("---")
    
    tab1, tab2 = st.tabs(["Interactive Explorer", "List View"])
    
    with tab1:
        st.subheader("File Explorer - Struktur Interaktif")
        st.caption("Klik folder untuk melihat isinya")
        
        highlight_name = None
        if st.session_state.search_result and st.session_state.search_result['found']:
            highlight_name = st.session_state.search_result['query']
        
        render_interactive_tree(root, "", highlight_name)
    
    with tab2:
        st.subheader("Semua File dan Folder - Daftar Lengkap")
        
        for i, node in enumerate(flat_stack):
            is_highlighted = (st.session_state.search_result and 
                            st.session_state.search_result['found'] and 
                            node.name == st.session_state.search_result['query'])
            
            if node.is_dir:
                icon = "📁"
                type_label = "[DIR]"
            else:
                icon = get_file_icon(node.name)
                type_label = "[FILE]"
            
            if is_highlighted:
                st.markdown(
                    f"<div style='background-color: #fff3cd; padding: 6px 10px; border-radius: 4px; "
                    f"margin: 2px 0; border-left: 3px solid #ffc107; font-weight: bold;'>"
                    f"{i+1}. {icon} {type_label} {node.path}</div>",
                    unsafe_allow_html=True
                )
            else:
                st.markdown(f"{i+1}. {icon} `{type_label}` {node.path}")


def main():
    st.set_page_config(
        page_title="File System Explorer",
        page_icon="📁",
        layout="wide",
        initial_sidebar_state="expanded"
    )
    
    st.markdown("""
        <style>
        .stButton>button {
            border-radius: 8px;
            transition: all 0.3s ease;
        }
        .stButton>button:hover {
            transform: translateY(-2px);
            box-shadow: 0 4px 8px rgba(0,0,0,0.1);
        }
        .stTextInput>div>div>input {
            border-radius: 8px;
        }
        h1 {
            color: #1f77b4;
        }
        h2, h3 {
            color: #2ca02c;
        }
        .stTabs [data-baseweb="tab-list"] {
            gap: 8px;
        }
        .stTabs [data-baseweb="tab"] {
            border-radius: 8px 8px 0 0;
        }
        .stExpander {
            border: 1px solid #e0e0e0;
            border-radius: 6px;
            margin: 4px 0;
            transition: all 0.2s ease;
        }
        .stExpander:hover {
            border-color: #1f77b4;
            box-shadow: 0 2px 6px rgba(31, 119, 180, 0.1);
        }
        div[data-testid="stExpander"] details summary {
            font-size: 14px;
            padding: 8px 12px;
        }
        </style>
    """, unsafe_allow_html=True)
    
    init_session_state()
    
    if st.session_state.page == 'select':
        select_file_system_page()
    else:
        explorer_page()


if __name__ == "__main__":
    main()
