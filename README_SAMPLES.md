# Sample File Systems - Dokumentasi

## Overview
Program ini menyediakan 3 variasi file system dengan kompleksitas berbeda untuk testing dan benchmark DFS.

## 1. SmallOS (File System Kecil)

### Karakteristik
- **Total Nodes**: ~20 nodes
- **Kedalaman**: 3 level
- **Struktur**: Sederhana dan mudah di-debug
- **Use Case**: Testing cepat, debugging, pembelajaran

### Struktur
```
SmallOS/
├── Documents/
│   ├── Work/
│   │   ├── quarterly_report
│   │   ├── annual_report
│   │   ├── client_presentation
│   │   ├── budget_sheet
│   │   └── meeting_notes
│   └── Personal/
│       ├── vacation_photo
│       ├── personal_diary
│       └── favorite_recipes
└── Applications/
    ├── Utilities/
    │   ├── calculator_app
    │   ├── text_editor
    │   └── file_manager
    └── Games/
        ├── puzzle_game
        ├── chess_game
        ├── solitaire_game
        └── tetris_game
```

### Statistik
- **Direktori**: 6
- **Files**: 15
- **Level Maksimum**: 3

### Contoh File untuk Testing
- `quarterly_report` (Level 3)
- `calculator_app` (Level 3)
- `tetris_game` (Level 3)

---

## 2. PacOS (File System Medium - Default)

### Karakteristik
- **Total Nodes**: ~50 nodes
- **Kedalaman**: 4 level
- **Struktur**: Balanced, representasi sistem nyata
- **Use Case**: Testing standar, demonstrasi

### Struktur
```
PacOS/
├── System/
│   ├── Kernel/
│   ├── Drivers/
│   └── Logs/
├── Users/
│   ├── Admin/
│   │   ├── Desktop/
│   │   ├── Documents/
│   │   └── Settings/
│   ├── Guest/
│   └── Student/
│       ├── Homework/
│       ├── Projects/
│       └── Downloads/
├── Programs/
│   ├── Games/
│   ├── Office/
│   └── Development/
└── Temp/
```

### Statistik
- **Direktori**: ~20
- **Files**: ~30
- **Level Maksimum**: 4

### Contoh File untuk Testing
- `bootloader` (Level 3)
- `system_monitor` (Level 4)
- `math_assignment` (Level 4)

---

## 3. LargeOS (File System Besar)

### Karakteristik
- **Total Nodes**: ~100+ nodes
- **Kedalaman**: 5 level
- **Struktur**: Kompleks, simulasi enterprise system
- **Use Case**: Stress testing, benchmark performa

### Struktur
```
LargeOS/
├── System/
│   ├── Kernel/ (6 files)
│   ├── Drivers/ (8 files)
│   └── Config/ (5 files)
├── Users/
│   ├── Admin/
│   │   ├── Documents/ (7 files)
│   │   ├── Projects/
│   │   │   ├── WebApp/ (4 files)
│   │   │   ├── MobileApp/ (4 files)
│   │   │   ├── Database/ (4 files)
│   │   │   ├── API/
│   │   │   └── Infrastructure/
│   │   └── Settings/
│   ├── User1/
│   ├── User2/
│   └── Guest/
├── Programs/
│   ├── Office/ (6 files)
│   ├── Development/ (7 files)
│   ├── Multimedia/
│   ├── Internet/
│   └── Utilities/
└── Data/
    ├── Databases/ (5 files)
    ├── Backups/ (6 files)
    └── Logs/ (8 files)
```

### Statistik
- **Direktori**: ~30
- **Files**: ~70+
- **Level Maksimum**: 5

### Contoh File untuk Testing
- `bootloader` (Level 3)
- `monthly_report` (Level 4)
- `frontend_code` (Level 5)
- `database_backup` (Level 3)

---

## Perbandingan

| Aspek | SmallOS | PacOS | LargeOS |
|-------|---------|-------|---------|
| Total Nodes | ~20 | ~50 | ~100+ |
| Kedalaman | 3 | 4 | 5 |
| Kompleksitas | Rendah | Medium | Tinggi |
| Waktu Benchmark | Cepat | Sedang | Lambat |
| Use Case | Debug/Learn | Demo/Standard | Stress Test |

---

## Cara Menggunakan

### 1. Memilih File System
Saat program dijalankan, pilih salah satu:
```
=== PILIH FILE SYSTEM ===
1. PacOS (Default - Medium, ~50 nodes)
2. SmallOS (Small - ~20 nodes)
3. LargeOS (Large - ~100+ nodes)
Pilih: 
```

### 2. Melihat Statistik
Program akan otomatis menampilkan statistik file system yang dipilih:
```
=== STATISTIK FILE SYSTEM ===
Nama Root       : SmallOS
Total Nodes     : 20
Total Direktori : 6
Total Files     : 14
Kedalaman Max   : 3 level
```

### 3. Benchmark Recommendation

**SmallOS**
- Iterasi: 5000-10000
- Waktu: < 1 detik
- Ideal untuk: Quick testing

**PacOS**
- Iterasi: 1000-5000
- Waktu: 1-3 detik
- Ideal untuk: Standard benchmark

**LargeOS**
- Iterasi: 100-1000
- Waktu: 3-10 detik
- Ideal untuk: Performance analysis

---

## Tips Benchmark

1. **Pilih file di berbagai kedalaman** untuk hasil yang representatif
2. **SmallOS**: Gunakan untuk verifikasi algoritma benar
3. **PacOS**: Gunakan untuk demo dan presentasi
4. **LargeOS**: Gunakan untuk analisis performa dan optimasi

## Fungsi Utility

### `getFileSystemStats(root *Node)`
Menghitung statistik file system:
- Total nodes
- Total direktori
- Total files
- Kedalaman maksimum

### `printFileSystemStats(root *Node)`
Menampilkan statistik dalam format yang mudah dibaca.

---

## Contoh Output

```
=== PILIH FILE SYSTEM ===
1. PacOS (Default - Medium, ~50 nodes)
2. SmallOS (Small - ~20 nodes)
3. LargeOS (Large - ~100+ nodes)
Pilih: 3

=== STATISTIK FILE SYSTEM ===
Nama Root       : LargeOS
Total Nodes     : 108
Total Direktori : 32
Total Files     : 76
Kedalaman Max   : 5 level

=== SISTEM PENCARIAN FILE SYSTEM DENGAN DFS ===
1. Tampilkan Diretory
2. Pencarian dengan DFS Rekursif
3. Pencarian dengan DFS Iteratif
4. Benchmark & Komparasi Waktu
5. Exit
```
