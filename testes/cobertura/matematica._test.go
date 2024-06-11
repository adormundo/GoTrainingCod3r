package cobertura

import (
	"testing"
)

func TestMedia(t *testing.T) {
	tests := []struct {
		numeros  []float64
		esperado float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3.0},
		{[]float64{1.5, 2.5, 3.5}, 2.5},
		{[]float64{5, 5, 5, 5}, 5.0},
		{[]float64{}, 0.0}, // Teste para slice vazio
	}

	for _, tt := range tests {
		resultado := Media(tt.numeros...)
		if resultado != tt.esperado {
			t.Errorf("Para %v, esperado %.2f, mas obtido %.2f", tt.numeros, tt.esperado, resultado)
		}
	}
}
