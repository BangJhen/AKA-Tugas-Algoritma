package main

import (
	"fmt"
	"math"
	"time"
)

// BenchmarkResult menyimpan hasil benchmark
type BenchmarkResult struct {
	Method   string
	Duration time.Duration
	Found    bool
}

// runBenchmark menjalankan benchmark untuk kedua metode
func runBenchmark(root *Node, target string, iterations int) []BenchmarkResult {
	var results []BenchmarkResult

	// Benchmark DFS Rekursif
	var totalRecursive time.Duration
	var foundRecursive bool
	var i int

	for i = 0; i < iterations; i++ {
		start := time.Now()
		_, foundRecursive = dfsRecursive(root, target)
		elapsed := time.Since(start)
		totalRecursive += elapsed
	}
	avgRecursive := totalRecursive / time.Duration(iterations)

	results = append(results, BenchmarkResult{
		Method:   "Rekursif",
		Duration: avgRecursive,
		Found:    foundRecursive,
	})

	// Benchmark DFS Iteratif
	var totalIterative time.Duration
	var foundIterative bool

	for i = 0; i < iterations; i++ {
		start := time.Now()
		_, foundIterative = dfsIterative(root, target)
		elapsed := time.Since(start)
		totalIterative += elapsed
	}
	avgIterative := totalIterative / time.Duration(iterations)

	results = append(results, BenchmarkResult{
		Method:   "Iteratif",
		Duration: avgIterative,
		Found:    foundIterative,
	})

	return results
}

// printBenchmarkResults menampilkan hasil benchmark dengan visualisasi
func printBenchmarkResults(results []BenchmarkResult) {
	fmt.Println("\n=== HASIL BENCHMARK ===")
	fmt.Println("Metode\t\t| Waktu Rata-rata\t| Status")
	fmt.Println("----------------------------------------------------------")

	var i int
	for i = 0; i < len(results); i++ {
		result := results[i]
		status := "Tidak Ditemukan"
		if result.Found {
			status = "Ditemukan"
		}
		fmt.Printf("%s\t| %v\t\t| %s\n", result.Method, result.Duration, status)
	}

	fmt.Println("\n=== VISUALISASI GRAFIK PERBANDINGAN WAKTU ===")
	printBarChart(results)

	// Hitung persentase perbedaan
	if len(results) >= 2 {
		recursive := results[0].Duration.Nanoseconds()
		iterative := results[1].Duration.Nanoseconds()

		var diff float64
		var faster string

		if recursive > iterative {
			diff = float64(recursive-iterative) / float64(recursive) * 100
			faster = "Iteratif"
		} else {
			diff = float64(iterative-recursive) / float64(iterative) * 100
			faster = "Rekursif"
		}

		fmt.Printf("\n%s lebih cepat %.2f%% dari metode lainnya\n", faster, diff)
	}
}

// printBarChart membuat visualisasi bar chart sederhana
func printBarChart(results []BenchmarkResult) {
	if len(results) == 0 {
		return
	}

	// Cari nilai maksimum untuk scaling
	var maxDuration time.Duration
	var i int
	for i = 0; i < len(results); i++ {
		if results[i].Duration > maxDuration {
			maxDuration = results[i].Duration
		}
	}

	// Lebar maksimum bar (dalam karakter)
	maxBarWidth := 50

	// Print bar untuk setiap hasil
	for i = 0; i < len(results); i++ {
		result := results[i]

		// Hitung panjang bar
		var barLength int
		if maxDuration > 0 {
			ratio := float64(result.Duration) / float64(maxDuration)
			barLength = int(math.Round(ratio * float64(maxBarWidth)))
		}

		// Print label metode
		fmt.Printf("%-10s | ", result.Method)

		// Print bar
		var j int
		for j = 0; j < barLength; j++ {
			fmt.Print("█")
		}

		// Print waktu
		fmt.Printf(" %v\n", result.Duration)
	}
}

// runMultipleBenchmarks menjalankan benchmark untuk beberapa file
func runMultipleBenchmarks(root *Node, targets []string, iterations int) {
	fmt.Println("\n=== BENCHMARK UNTUK MULTIPLE FILE ===")
	fmt.Printf("Jumlah iterasi per file: %d\n", iterations)

	var allRecursive []time.Duration
	var allIterative []time.Duration
	var i int

	for i = 0; i < len(targets); i++ {
		target := targets[i]
		fmt.Printf("\n--- Testing file: %s ---\n", target)

		results := runBenchmark(root, target, iterations)
		printBenchmarkResults(results)

		if len(results) >= 2 {
			allRecursive = append(allRecursive, results[0].Duration)
			allIterative = append(allIterative, results[1].Duration)
		}
	}

	// Print summary
	if len(allRecursive) > 0 && len(allIterative) > 0 {
		fmt.Println("\n=== RINGKASAN KESELURUHAN ===")

		var totalRecursive, totalIterative time.Duration
		for i = 0; i < len(allRecursive); i++ {
			totalRecursive += allRecursive[i]
			totalIterative += allIterative[i]
		}

		avgRecursive := totalRecursive / time.Duration(len(allRecursive))
		avgIterative := totalIterative / time.Duration(len(allIterative))

		fmt.Printf("Rata-rata Rekursif : %v\n", avgRecursive)
		fmt.Printf("Rata-rata Iteratif : %v\n", avgIterative)

		// Visualisasi summary
		summaryResults := []BenchmarkResult{
			{Method: "Rekursif", Duration: avgRecursive, Found: true},
			{Method: "Iteratif", Duration: avgIterative, Found: true},
		}

		fmt.Println("\n=== GRAFIK RATA-RATA KESELURUHAN ===")
		printBarChart(summaryResults)
	}
}

// testSingleIteration mengukur waktu untuk 1 iterasi pencarian
func testSingleIteration(root *Node) {
	fmt.Println("\n=== TEST KECEPATAN LAPTOP (1 ITERASI) ===")
	fmt.Println("Mengukur waktu yang dibutuhkan untuk 1 kali pencarian")

	var target string
	fmt.Print("\nMasukkan nama file/direktori yang dicari: ")
	fmt.Scanln(&target)

	fmt.Println("HASIL PENGUKURAN")

	// Test DFS Rekursif
	fmt.Println("\n1. DFS REKURSIF:")
	start := time.Now()
	path, found := dfsRecursive(root, target)
	elapsed := time.Since(start)

	if found {
		fmt.Printf("   Status      : ✓ Ditemukan\n")
		fmt.Printf("   Path        : %s\n", path)
	} else {
		fmt.Printf("   Status      : ✗ Tidak ditemukan\n")
	}
	fmt.Printf("   Waktu       : %v\n", elapsed)
	fmt.Printf("   Nanosecond  : %d ns\n", elapsed.Nanoseconds())
	fmt.Printf("   Microsecond : %.3f µs\n", float64(elapsed.Nanoseconds())/1000.0)

	// Test DFS Iteratif
	fmt.Println("\n2. DFS ITERATIF:")
	start = time.Now()
	path, found = dfsIterative(root, target)
	elapsed = time.Since(start)

	if found {
		fmt.Printf("   Status      : ✓ Ditemukan\n")
		fmt.Printf("   Path        : %s\n", path)
	} else {
		fmt.Printf("   Status      : ✗ Tidak ditemukan\n")
	}
	fmt.Printf("   Waktu       : %v\n", elapsed)
	fmt.Printf("   Nanosecond  : %d ns\n", elapsed.Nanoseconds())
	fmt.Printf("   Microsecond : %.3f µs\n", float64(elapsed.Nanoseconds())/1000.0)

	fmt.Println("\n========================================")
	fmt.Println("         INFORMASI")
	fmt.Println("========================================")
	fmt.Println("• 1 mikrodetik (µs) = 0.000001 detik")
	fmt.Println("• 1 nanodetik (ns)  = 0.000000001 detik")
	fmt.Println("• Waktu sangat kecil karena pencarian sangat cepat")
	fmt.Println("• Untuk hasil lebih akurat, gunakan benchmark")
	fmt.Println("  dengan banyak iterasi (menu 2-4)")
}

// benchmarkMenu menampilkan menu benchmark
func benchmarkMenu(root *Node) {
	var pilih int
	var target string
	var iterations int

	for pilih != 5 {
		fmt.Println("\n=== MENU BENCHMARK DFS ===")
		fmt.Println("1. Test Kecepatan Laptop (1 Iterasi)")
		fmt.Println("2. Benchmark Single File")
		fmt.Println("3. Benchmark Multiple Files")
		fmt.Println("4. Benchmark Custom")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Print("Pilih: ")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			testSingleIteration(root)

		case 2:
			fmt.Print("Masukkan nama file yang dicari: ")
			fmt.Scanln(&target)
			fmt.Print("Jumlah iterasi (default 1000): ")
			fmt.Scanln(&iterations)
			if iterations <= 0 {
				iterations = 1000
			}

			results := runBenchmark(root, target, iterations)
			printBenchmarkResults(results)

		case 3:
			fmt.Print("Jumlah iterasi per file (default 1000): ")
			fmt.Scanln(&iterations)
			if iterations <= 0 {
				iterations = 1000
			}

			// Test dengan beberapa file di berbagai kedalaman
			targets := []string{
				"bootloader",           // Level 3
				"system_monitor",       // Level 4
				"math_assignment",      // Level 4
				"programming_tutorial", // Level 4
				"temp_backup",          // Level 2
			}

			runMultipleBenchmarks(root, targets, iterations)

		case 4:
			fmt.Print("Masukkan nama file yang dicari: ")
			fmt.Scanln(&target)
			fmt.Print("Jumlah iterasi: ")
			fmt.Scanln(&iterations)
			if iterations <= 0 {
				iterations = 100
			}

			results := runBenchmark(root, target, iterations)
			printBenchmarkResults(results)

		case 5:
			fmt.Println("Kembali ke menu utama...")

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
