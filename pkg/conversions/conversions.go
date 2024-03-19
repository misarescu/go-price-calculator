package conversions

import (
	"strconv"
)

func StringsToFloats(input []string) ([]float64, error) {
	result := []float64{}
	for _, strNum := range input {
		floatNum, err := strconv.ParseFloat(strNum, len(strNum))
		if err != nil {
			return []float64{}, err
		}
		result = append(result, floatNum)
	}
	return result, nil
}
