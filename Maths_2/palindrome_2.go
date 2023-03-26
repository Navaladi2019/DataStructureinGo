package maths

func IsPalindromenaval(i int) bool {

	if uint(i) < 10 {
		return true
	}

	target := make([]int, 0)

	for i > 0 {
		r := i % 10
		i = (i - r) / 10
		target = append(target, r)
	}

	middelIndex := len(target) / 2

	if len(target)/2%2 > 0 {
		middelIndex = middelIndex + 1
	}

	IsPalindrome := true

	for l := 0; l <= middelIndex; l++ {
		if target[l] != target[len(target)-1-l] {
			IsPalindrome = false
			break
		}
	}

	return IsPalindrome
}

var IsPalindrome = func(i int) bool {

	lastdigit := i
	reversedInt := 0

	for lastdigit != 0 {
		reversedInt = reversedInt*10 + (lastdigit % 10)
		lastdigit = lastdigit / 10
	}

	return reversedInt == i
}

var Factorial = func(i int) int {
	target := 1

	for l := 1; l <= i; l++ {
		target = target * l
	}
	return target
}

func FactorailRecursive(i int) int {
	if i == 1 {
		return 1
	}

	return i * FactorailRecursive(i-1)
}

func Factorialnaval(i int) int {

	// n! = n *(n-1)*(n-2)*.....2,1
	target := i
	for i > 0 {
		target = target * (i - 1)
		i = i - 1
	}
	return target
}

func FindFactorialTraiingZerosCount(i int) int {

	fact := FactorailRecursive(i)
	z := 0
	for fact%10 == 0 && fact > 10 {
		z++
		fact = fact / 10
	}
	return z
}
