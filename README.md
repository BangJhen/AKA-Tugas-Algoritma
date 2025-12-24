# Sistem Pencarian File System dengan DFS

Program sederhana untuk mencari file dalam struktur file system menggunakan algoritma Depth-First Search (DFS) yang diimplementasikan dalam bahasa Go.

## Fitur

- Implementasi DFS secara **Rekursif**
- Implementasi DFS secara **Iteratif** (menggunakan Stack)
- Struktur kode modular dan clean
- Hanya menggunakan library standard Go (fmt)
- Tidak menggunakan short declaration (:=)
- Menggunakan function sederhana tanpa method pada struct

## Struktur File

```
.
├── main.go           # File utama program
├── node.go           # Definisi struktur data Node dan Stack
├── dfs_recursive.go  # Implementasi DFS rekursif
├── dfs_iterative.go  # Implementasi DFS iteratif
└── README.md         # Dokumentasi
```

## Cara Menjalankan

```bash
go run .
```

atau

```bash
go run main.go node.go dfs_recursive.go dfs_iterative.go sample_filesystems.go
```

## Struktur Data

### Node
Merepresentasikan file atau direktori dalam file system:
- `Name`: Nama file atau direktori
- `IsDir`: Boolean untuk menandai apakah ini direktori
- `Children`: Array pointer ke node children (untuk direktori)

### Stack
Struktur data stack untuk implementasi DFS iteratif:
- `pushStack`: Function untuk menambah item ke stack
- `popStack`: Function untuk mengambil item teratas
- `isStackEmpty`: Function untuk mengecek apakah stack kosong

## Algoritma DFS

### DFS Rekursif
- Menggunakan rekursi untuk traversal tree
- Base case: node nil atau file ditemukan
- Recursive case: cek semua children dari direktori

### DFS Iteratif
- Menggunakan stack eksplisit untuk menyimpan node yang akan dikunjungi
- Loop hingga stack kosong
- Push children ke stack dalam urutan terbalik

## Contoh Output

Program akan menampilkan:
1. Struktur file system yang dibuat
2. Hasil pencarian dengan DFS Rekursif
3. Hasil pencarian dengan DFS Iteratif
4. Test untuk file yang tidak ada

## Kompleksitas

- **Time Complexity**: O(n) dimana n adalah jumlah total node
- **Space Complexity**: 
  - Rekursif: O(h) dimana h adalah tinggi tree (call stack)
  - Iteratif: O(n) worst case untuk stack
