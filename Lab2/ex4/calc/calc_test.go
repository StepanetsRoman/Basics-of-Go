package calc

import "testing"

// TestMin тестує функцію Min
func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		c        float64
		expected float64
	}{
		{"Позитивні числа", 5, 2, 8, 2},
		{"Негативні числа", -3, -1, -5, -5},
		{"Однакові числа", 10, 10, 10, 10},
		{"Змішані числа", -5, 0, 5, -5},
		{"Дробові числа", 1.5, 2.3, 0.8, 0.8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Min(tt.a, tt.b, tt.c)
			if result != tt.expected {
				t.Errorf("Min(%f, %f, %f) = %f; очікується %f", tt.a, tt.b, tt.c, result, tt.expected)
			}
		})
	}
}

// TestAverage тестує функцію Average
func TestAverage(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		c        float64
		expected float64
	}{
		{"Позитивні числа", 5, 10, 15, 10.0},
		{"Послідовні числа", 1, 2, 3, 2.0},
		{"Змішані числа", -5, 0, 5, 0.0},
		{"Однакові числа", 3, 3, 3, 3.0},
		{"Дробові числа", 1.5, 2.5, 3.5, 2.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Average(tt.a, tt.b, tt.c)
			if result != tt.expected {
				t.Errorf("Average(%f, %f, %f) = %f; очікується %f", tt.a, tt.b, tt.c, result, tt.expected)
			}
		})
	}
}

// TestSolveLinearEquation тестує функцію SolveLinearEquation
func TestSolveLinearEquation(t *testing.T) {
	tests := []struct {
		name        string
		a            float64
		b            float64
		expectedX    float64
		expectError  bool
	}{
		{"Нормальне рівняння 1", 2, -4, 2.0, false},
		{"Нормальне рівняння 2", 3, 6, -2.0, false},
		{"Нормальне рівняння 3", -2, 4, 2.0, false},
		{"Ділення на нуль", 0, 5, 0, true},
		{"Дробові коефіцієнти", 0.5, -1, 2.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SolveLinearEquation(tt.a, tt.b)
			if tt.expectError {
				if err == nil {
					t.Errorf("SolveLinearEquation(%f, %f) очікувала помилку, але отримала результат %f", tt.a, tt.b, result)
				}
			} else {
				if err != nil {
					t.Errorf("SolveLinearEquation(%f, %f) повернула помилку: %v", tt.a, tt.b, err)
				}
				if result != tt.expectedX {
					t.Errorf("SolveLinearEquation(%f, %f) = %f; очікується %f", tt.a, tt.b, result, tt.expectedX)
				}
			}
		})
	}
}

// BenchmarkMin бенчмарк для функції Min
func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Min(5, 2, 8)
	}
}

// BenchmarkAverage бенчмарк для функції Average
func BenchmarkAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Average(5, 10, 15)
	}
}

// BenchmarkSolveLinearEquation бенчмарк для функції SolveLinearEquation
func BenchmarkSolveLinearEquation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SolveLinearEquation(2, -4)
	}
}
