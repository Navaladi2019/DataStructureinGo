package intro

import "fmt"

/*This method has o(n) notation
but if we see SumNaturalNumbersPerfect it uses o(1) notation*/
func SumNaturalNumbers(n int) {

	op := 0
	for i := 1; i <= n; i++ {
		op = op + i
	}

	fmt.Println(op)
}

func SumNaturalNumbersPerfect(n int) {

	op := n * (n + 1) / 2

	fmt.Println(op)
}
