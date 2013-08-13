package approximate

import (
	"fmt"
	"strconv"
	"testing"
)

func test(input int, expected string, units map[string]string, t *testing.T) {
	output := "* " + strconv.Itoa(input) + " - "
	result := Approximate(input)
	colorGreen := "\x1b[32;1m"
	colorRed := "\x1b[31;1m"
	resetColor := "\x1b[0m"

	if result == expected {
		output += colorGreen + "ok" + resetColor
		fmt.Println(output)
	} else {
		output += colorRed + "expected " + expected + ", got " + result + resetColor
		fmt.Println(output)
		t.Fail()
	}
}

func TestApproximate(t *testing.T) {
	units := make(map[string]string)

	test(7, "7", units, t)
	test(800, "800", units, t)
	test(1200, "1.2K", units, t)
	test(1000, "1K", units, t)
	test(87821, "88K", units, t)
	test(412011, "412K", units, t)
	test(999949, "999K", units, t)
	test(999950, "1M", units, t)
	test(1000000, "1M", units, t)
	test(6123922, "6.1M", units, t)
	test(27969374, "28M", units, t)
	test(28001341, "28M", units, t)
	test(28061341, "28M", units, t)
	test(812192442, "812M", units, t)
	test(999949999, "999M", units, t)
	test(999950000, "1G", units, t)
	test(1000000000, "1G", units, t)

	test(-8, "-8", units, t)
	test(-5000, "-5K", units, t)
	test(-999949999, "-999M", units, t)
	test(-999950000, "-1G", units, t)
}
