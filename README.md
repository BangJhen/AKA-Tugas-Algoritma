# File System Explorer - PacOS

Program pencarian file system menggunakan algoritma DFS (Depth-First Search) dengan dua implementasi: Rekursif dan Iteratif.

## Persyaratan Sistem

- Python 3.8 atau lebih baru
- pip (Python package installer)

## Instalasi

1. Pastikan Python sudah terinstall:
```bash
python --version
```

2. Install dependencies untuk GUI (opsional):
```bash
pip install -r requirements.txt
```

## Cara Menjalankan

### 1. Mode Command Line (CLI)

Jalankan program berbasis terminal:

```bash
python main.py
```

**Fitur CLI:**
- Menu interaktif di terminal
- Pilih file system (SmallOS, PacOS, LargeOS, dll)
- Tampilkan struktur direktori
- Cari file dengan DFS Rekursif atau Iteratif

**Contoh Penggunaan:**
```
=== PILIH FILE SYSTEM ===
1. PacOS (Default - Medium, ~50 nodes)
2. SmallOS (Small - ~20 nodes)
...
Pilih: 1

=== SISTEM PENCARIAN FILE SYSTEM DENGAN DFS ===
1. Tampilkan Diretory
2. Pencarian dengan DFS Rekursif
3. Pencarian dengan DFS Iteratif
4. Exit
Pilih: 2

Masukkan nama file yang dicari: bootloader
File Ditemukan! Path: PacOS/System/Kernel/bootloader
```

### 2. Mode Grafis (GUI)

Jalankan aplikasi dengan interface web modern:

```bash
streamlit run streamlit_app.py
```

Browser akan otomatis terbuka di `http://localhost:8501`

**Fitur GUI:**
- Interface modern seperti File Explorer
- 3 mode tampilan: Interactive, Compact, List
- Search bar dengan pilihan algoritma
- Highlight hasil pencarian
- Statistik file system real-time
- Folder dapat diklik untuk expand/collapse

## Struktur File

```
.
├── main.py                    # Program CLI
├── streamlit_app.py           # Program GUI
├── node.py                    # Definisi Node dan FlatNode
├── dfs_recursive.py           # Algoritma DFS Rekursif
├── dfs_iterative.py           # Algoritma DFS Iteratif
├── sample_filesystems.py      # Contoh file system
├── requirements.txt           # Dependencies
├── test.py                    # Script testing
└── README.md                  # Dokumentasi ini
```

## Pilihan File System

1. **SmallOS** - Kecil (~20 nodes, 3 levels)
   - Cocok untuk testing cepat

2. **PacOS** - Medium (~50 nodes, 4 levels)
   - File system default, balanced

3. **LargeOS** - Besar (~100+ nodes, 5 levels)
   - Testing performa dengan data besar

4. **FlatOS** - Sederhana (2 levels)
   - Hanya root dan files, tanpa subfolder

5. **BinaryTreeOS** - Terstruktur (31 nodes)
   - Perfect Binary Tree

6. **Custom Binary Tree** - Kustom (depth 3-7)
   - Sesuaikan kedalaman tree sendiri

## Algoritma Pencarian

### DFS Rekursif (Linear Recursion)
- Menggunakan rekursi untuk menelusuri tree
- Kompleksitas: O(n)
- Relasi rekurensi: T(n) = T(n-1) + 1

### DFS Iteratif (Linear Loop)
- Menggunakan loop untuk menelusuri tree
- Kompleksitas: O(n)
- Linear search pada flat stack

## Tips Penggunaan

### CLI
- Pilih file system yang sesuai dengan kebutuhan testing
- Gunakan "Tampilkan Diretory" untuk melihat struktur
- Cari file dengan nama lengkap (case-sensitive)

### GUI
- Gunakan sidebar untuk pencarian
- Pilih algoritma sebelum mencari
- Klik folder di Interactive Explorer untuk melihat isi
- Gunakan tab berbeda untuk view yang berbeda
- Reset pencarian untuk membersihkan highlight

## Contoh File yang Bisa Dicari

**Di PacOS:**
- bootloader
- calculator
- ai_project
- system_monitor
- graphics_driver

**Di SmallOS:**
- puzzle_game
- quarterly_report
- text_editor

## Troubleshooting

### Program tidak jalan
```bash
# Pastikan di direktori yang benar
cd "path/to/Sistem File System PacOS"

# Jalankan ulang
python main.py
```

### Streamlit tidak bisa diakses
```bash
# Install ulang Streamlit
pip install --upgrade streamlit

# Jalankan dengan port berbeda
streamlit run streamlit_app.py --server.port 8502
```

### Import error
```bash
# Pastikan semua file Python ada di folder yang sama
ls *.py

# Hasil harus menampilkan: main.py, node.py, dfs_recursive.py, dll
```

## Testing

Jalankan script testing untuk memverifikasi program:

```bash
python test.py
```

Output yang benar:
```
Test Program File System PacOS
========================================
[OK] bootloader -> PacOS/System/Kernel/bootloader
[OK] calculator -> PacOS/Programs/Office/calculator
...
Test selesai!
```

## Kompleksitas Algoritma

| Algoritma | Time | Space | Keterangan |
|-----------|------|-------|------------|
| DFS Rekursif | O(n) | O(h) | h = tinggi tree |
| DFS Iteratif | O(n) | O(n) | Flat stack |

Kedua algoritma memiliki kompleksitas waktu yang sama (O(n)), perbedaan utama adalah penggunaan memory.

## Lisensi

Program ini dibuat untuk keperluan akademik - Tugas AKA Semester 3.

## Kontak

Untuk pertanyaan atau masalah, silakan hubungi pengembang.

---

**Selamat mencoba!** 🚀
