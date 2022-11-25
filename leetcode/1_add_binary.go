package leetcode

func AddBinary(a string, b string) string {
	var i, j int
	var carry bool = false
	var str, strA, strB string
	if len(a) >= len(b) {
		i = len(a) - 1
		j = len(b) - 1
	} else {
		i = len(b) - 1
		j = len(a) - 1
		//swap 'a' and 'b'
		temp := a
		a = b
		b = temp
	}
	for ; i >= 0; i-- {
		strA = string(a[i])
		if j >= 0 {
			strB = string(b[j])
			if strA == "1" && strB == "1" {
				if carry {
					str = "1" + str
				} else {
					str = "0" + str
				}
				carry = true
			} else if strA == "0" && strB == "0" {
				if carry {
					str = "1" + str
				} else {
					str = "0" + str
				}
				carry = false
			} else {
				if carry {
					str = "0" + str
					carry = true
				} else {
					str = "1" + str
					carry = false
				}
			}
			j--
		} else {
			if strA == "1" {
				if carry {
					str = "0" + str
					carry = true
				} else {
					str = "1" + str
					carry = false
				}
			} else {
				if carry {
					str = "1" + str
				} else {
					str = "0" + str
				}
				carry = false
			}
		}
	}
	if carry {
		str = "1" + str
	}
	return str
}
