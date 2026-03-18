// Лабораторна робота №3
// Управляючі конструкції. Функції. Введення та виведення. Псевдовипадкові числа.
package main

import (
	"fmt"
	"math"
)

// Параметри варіанту з таблиці 3.2
type variant struct {
	a, c, m, n, K int
}

var table32 = map[int]variant{
	1:  {1103515245, 12345, 1 << 31, 100, 10000},
	2:  {1664525, 1013904223, 1 << 32, 150, 20000},
	3:  {16807, 0, 1<<31 - 1, 200, 30000},
	4:  {22695477, 1, 1 << 32, 250, 40000},
	5:  {48271, 0, 1<<31 - 1, 300, 50000},
	6:  {214013, 2531011, 1 << 32, 100, 50000},
	7:  {2147483629, 2147483587, 1<<31 - 1, 150, 40000},
	8:  {65539, 0, 1 << 32, 200, 30000},
	9:  {1140671485, 12820163, 1 << 24, 250, 20000},
	10: {69069, 1, 1 << 32, 300, 10000},
	11: {1103515245, 12345, 1 << 31, 300, 50000},
	12: {1664525, 1013904223, 1 << 32, 250, 40000},
	13: {16807, 0, 1<<31 - 1, 200, 30000},
	14: {22695477, 1, 1 << 32, 150, 20000},
	15: {48271, 0, 1<<31 - 1, 100, 10000},
	16: {214013, 2531011, 1 << 32, 300, 10000},
	17: {2147483629, 2147483587, 1<<31 - 1, 250, 20000},
	18: {65539, 0, 1 << 32, 200, 30000},
	19: {1140671485, 12820163, 1 << 24, 150, 40000},
	20: {134775813, 1, 1 << 32, 100, 50000},
}

func main() {
	fmt.Println("=== Лабораторна робота №3 ===")
	fmt.Println("Лінійний конгруентний метод: Xk+1 = (a*Xk + c) mod m")
	fmt.Print("Введіть номер варіанту (1-20) або 0 для варіанту 1: ")
	var vNum int
	_, err := fmt.Scanf("%d", &vNum)
	if err != nil || vNum < 0 || vNum > 20 {
		vNum = 1
	}
	if vNum == 0 {
		vNum = 1
	}
	vr, ok := table32[vNum]
	if !ok {
		vr = table32[1]
		vNum = 1
	}
	a, c, m, n, K := vr.a, vr.c, vr.m, vr.n, vr.K
	seed := 12345
	fmt.Printf("Обрано варіант %d: a=%d, c=%d, m=2^31/32, діапазон [0, %d), K=%d\n\n", vNum, a, c, n, K)

	// 1. Генерація цілочислової послідовності
	intSeq := randIntSeq(a, c, m, seed, n, K)
	fmt.Printf("Згенеровано %d цілих значень у діапазоні [0, %d)\n\n", len(intSeq), n)

	// 2. Обробка масиву
	freq := frequency(intSeq, n)
	fmt.Println("--- Частота інтервалів (інтервал = 1) ---")
	printFreqSample(freq, n, 10)

	probs := statisticalProbability(freq, K)
	fmt.Println("\n--- Статистична ймовірність ---")
	printProbSample(probs, n, 10)

	mx := mathExpectation(probs, n)
	fmt.Printf("\nМатематичне сподівання M(X) = %.6f\n", mx)

	dx := variance(probs, n, mx)
	fmt.Printf("Дисперсія D(X) = %.6f\n", dx)

	sigma := math.Sqrt(dx)
	fmt.Printf("Середньоквадратичне відхилення σ(X) = %.6f\n", sigma)

	// 3. Генерація послідовності дійсних значень
	fmt.Println("\n--- Послідовність псевдовипадкових дійсних значень ---")
	realSeq := randRealSeq(a, c, m, seed, n, K)
	fmt.Printf("Згенеровано %d дійсних значень у діапазоні [0, %d)\n", len(realSeq), n)
	printRealSample(realSeq, 15)
}

// randIntSeq генерує цілочислову послідовність псевдовипадкових значень
// лінійним конгруентним методом: Xk+1 = (a*Xk + c) mod m
// Повертає K значень у діапазоні [0, n).
func randIntSeq(a, c, m, x0, n, K int) []int {
	seq := make([]int, K)
	x := int64(x0)
	mm := int64(m)

	for i := 0; i < K; i++ {
		x = (int64(a)*x + int64(c)) % mm
		if x < 0 {
			x += mm
		}
		seq[i] = int(x) % n
	}
	return seq
}

// randRealSeq генерує послідовність псевдовипадкових дійсних значень у [0, n).
func randRealSeq(a, c, m, x0, n, K int) []float64 {
	seq := make([]float64, K)
	x := int64(x0)
	mm := int64(m)
	nf := float64(n)

	for i := 0; i < K; i++ {
		x = (int64(a)*x + int64(c)) % mm
		if x < 0 {
			x += mm
		}
		seq[i] = (float64(x) / float64(mm)) * nf
	}
	return seq
}

// frequency обчислює частоту інтервалів (кількість появи кожного значення 0..n-1).
func frequency(seq []int, n int) []int {
	freq := make([]int, n)
	for _, v := range seq {
		if v >= 0 && v < n {
			freq[v]++
		}
	}
	return freq
}

// statisticalProbability P(i) = L/K — відносна частота.
func statisticalProbability(freq []int, K int) []float64 {
	probs := make([]float64, len(freq))
	for i := range freq {
		probs[i] = float64(freq[i]) / float64(K)
	}
	return probs
}

// mathExpectation M(X) = sum(xi * pi).
func mathExpectation(probs []float64, n int) float64 {
	var mx float64
	for i := 0; i < n; i++ {
		mx += float64(i) * probs[i]
	}
	return mx
}

// variance D(X) = sum((xi - M(X))^2 * pi).
func variance(probs []float64, n int, mx float64) float64 {
	var dx float64
	for i := 0; i < n; i++ {
		dx += (float64(i) - mx) * (float64(i) - mx) * probs[i]
	}
	return dx
}

func printFreqSample(freq []int, n, sample int) {
	if sample > n {
		sample = n
	}
	for i := 0; i < sample; i++ {
		fmt.Printf("  значення %d: частота %d\n", i, freq[i])
	}
	fmt.Printf("  ... (показано перші %d з %d)\n", sample, n)
}

func printProbSample(probs []float64, n, sample int) {
	if sample > n {
		sample = n
	}
	for i := 0; i < sample; i++ {
		fmt.Printf("  P(%d) = %.6f\n", i, probs[i])
	}
	fmt.Printf("  ... (показано перші %d з %d)\n", sample, n)
}

func printRealSample(seq []float64, count int) {
	if count > len(seq) {
		count = len(seq)
	}
	for i := 0; i < count; i++ {
		fmt.Printf("  [%d] = %.4f\n", i, seq[i])
	}
	fmt.Printf("  ... (показано перші %d з %d)\n", count, len(seq))
}
