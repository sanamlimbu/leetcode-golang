package leetcode

/*
	Write a function that reverses a string. The input string is given as an array of characters s.
	You must do this by modifying the input array in-place with O(1) extra memory.
*/

func ReverseString(s []byte) {
	length := len(s)
	for headIndex := 0; headIndex < length/2; headIndex++ {
		tailIndex := length - 1 - headIndex
		temp := s[headIndex]
		s[headIndex] = s[tailIndex]
		s[tailIndex] = temp
	}
}
