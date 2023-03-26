package maths

import (
	"testing"
)

func TestPalindromenaval(t *testing.T) {

	of := IsPalindrome

	defer func() {
		IsPalindrome = of
	}()
	IsPalindrome = IsPalindromenaval

}

func TestPalindrome(t *testing.T) {
	testFindPalindome(t)
}

func testFindPalindome(t *testing.T) {

	/*Even in length should result in palindrome*/
	ispalindrome := IsPalindrome(4444)
	printResult(true, ispalindrome, 4444, t)

	/*Even in length should not result in palindrome*/
	ispalindrome = IsPalindrome(4443)
	printResult(false, ispalindrome, 4443, t)

	/*odd in length should result in palindrome*/
	ispalindrome = IsPalindrome(44444)
	printResult(true, ispalindrome, 44444, t)

	/*Even in length should result in palindrome with middle number as differnt*/
	ispalindrome = IsPalindrome(44344)
	printResult(true, ispalindrome, 44344, t)

	/*single digit is always considered palindrome*/
	ispalindrome = IsPalindrome(0)
	printResult(true, ispalindrome, 0, t)

	/*odd in length should not result in palindrome*/
	ispalindrome = IsPalindrome(44334)
	printResult(false, ispalindrome, 44334, t)

}

func printResult(e bool, r bool, i int, t *testing.T) {

	if e != r {
		t.Errorf("Expected %t but got %t for the input %d", e, r, i)
	} else {
		t.Logf("Expected %t  got %t for the input %d", e, r, i)
	}

}

func TestFactorialNaval(t *testing.T) {
	//Factorial = Factorialnaval

	i := FactorailRecursive(4)

	t.Logf("Expected %d and got %d", 24, i)
}

func TestFactorialTrailingZero(t *testing.T) {
	//Factorial = Factorialnaval

	i := FindFactorialTraiingZerosCount(5)

	t.Logf("Expected %d and got %d", 1, i)
}
