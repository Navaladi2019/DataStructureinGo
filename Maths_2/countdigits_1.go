package maths

import (
	"fmt"
	"strconv"
)

/*My solution initially to count digits*/
func Coutdigits(i int) {
	str := strconv.FormatInt(int64(i), 10)

	fmt.Println(len(str))
}

func CountDigitsWithoutConvertingToStrint(i uint) {

	digits := 0
	for i > 0 {

		i = i / 10
		digits++
	}

	fmt.Println(digits)
}
