package cobertura

import (
	"math"
)

func Media(numeros ...float64) float64 {
	if len(numeros) == 0 {
		return 0.0
	}

	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))

	mediaArredondada := math.Round(media*100) / 100 // Arredonda para duas casas decimais

	return mediaArredondada
}
