package leetcode

import "fmt"

func AddStrings(num1 string, num2 string) string {
	var sum string
	var i, j int = len(num1) - 1, len(num2) - 1
	var carry int = 0

	// swap 'num1' and 'num2'
	if j > i {
		tempIndex := i
		i = j
		j = tempIndex

		tempNum := num1
		num1 = num2
		num2 = tempNum
	}

	var intOne, intTwo, reminder, temp int

	for ; i >= 0; i-- {
		// convert rune to actual int not ASCII value
		intOne = int(num1[i] - '0')
		if j >= 0 {
			// convert rune to actual int not ASCII value
			intTwo = int(num2[j] - '0')
			temp = intOne + intTwo + carry
		} else {
			temp = intOne + carry
		}
		reminder = temp % 10
		if reminder == 0 {
			sum = "0" + sum
		} else {
			sum = fmt.Sprint(reminder) + sum
		}
		carry = temp / 10
		j--
	}
	if carry != 0 {
		sum = fmt.Sprint(carry) + sum
	}
	return sum
}
