package leetcode

/*
	Question:
	Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
	An input string is valid if:
	Open brackets must be closed by the same type of brackets.
	Open brackets must be closed in the correct order.
	Every close bracket has a corresponding open bracket of the same type.
*/

/*
	Solution idea:
	Use a stack to store the opening brackets
  	Pop where the matching closing brackets appear in the string
  	Return true when the final stack is empty

 	Complexity:
  	Time: O(n) - we need to loop through the string once
  	Space: O(n) - worst case we would need to store the whole string in a stack
*/

func IsValid(s string) bool {
	// slice used as stack
	var stack []string

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if !isClosing(c) {
			stack = append(stack, c)
		} else {
			// pop stack
			l := len(stack)
			if l == 0 {
				return false
			}
			lastItem := string(stack[l-1])
			stack = stack[:l-1]

			// and see if matching
			if lastItem != getMatching(c) {
				return false
			}
		}
	}
	return len(stack) == 0
}

func isClosing(s string) bool {
	if s == ")" || s == "}" || s == "]" {
		return true
	}

	return false
}

func getMatching(s string) string {
	if s == ")" {
		return "("
	}
	if s == "}" {
		return "{"
	}
	if s == "]" {
		return "["
	}

	return ""
}
