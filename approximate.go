package approximate

import (
	"math"
	"strconv"
	"strings"
)

var units = map[string]string{
	"thousand": "K",
	"million":  "M",
	"billion":  "G",
}

func round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow

	if intermed < 0.0 {
		intermed -= 0.5
	} else {
		intermed += 0.5
	}
	rounder = float64(int64(intermed))
	return rounder / float64(pow)
}

func Approximate(number int) string {
	divisor := 1.0
	unit := ""
	numberFloat := float64(number)
	absNumber := math.Abs(numberFloat)
	if absNumber >= 999950000 { // billions
		divisor = 1e9
		unit = units["billion"]
	} else if absNumber >= 999950 { // millions
		divisor = 1e6
		unit = units["million"]
	} else if absNumber >= 1000 { // thousands
		divisor = 1e3
		unit = units["thousand"]
	}
	division := numberFloat / divisor

	precision := 0

	if math.Abs(division) < 10 && division-round(division, 0) != 0 {
		precision = 1
	}
	numberString := strconv.FormatFloat(round(division, precision), 'f', -1, 64)
	numberString = strings.Replace(numberString, "1000", "999", -1) + unit
	return numberString
}
