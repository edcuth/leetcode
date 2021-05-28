//20. Valid Parentheses
//https://leetcode.com/problems/valid-parentheses/
func remove(s []byte, i int) []byte {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func isValid(s string) bool {
	c1 := 0
	c2 := 0
	c3 := 0
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			c1++
			stack = append(stack, s[i])
		} else if string(s[i]) == ")" {
			c1--
			if c1 < 0 || string(stack[len(stack)-1]) != "(" {
				fmt.Println("c1 is false")
				return false
			}
			stack = remove(stack, len(stack)-1)
		} else if string(s[i]) == "[" {
			c2++
			stack = append(stack, s[i])
		} else if string(s[i]) == "]" {
			c2--
			if c2 < 0 || string(stack[len(stack)-1]) != "[" {
				fmt.Println("c2 is false")
				return false
			}
			stack = remove(stack, len(stack)-1)
		} else if string(s[i]) == "{" {
			c3++
			stack = append(stack, s[i])
		} else if string(s[i]) == "}" {
			c3--
			if c3 < 0 || string(stack[len(stack)-1]) != "{" {
				fmt.Println("c3 is false")
				return false
			}
			stack = remove(stack, len(stack)-1)
		}
	}
	if c1 == 0 && c2 == 0 && c3 == 0 {
		return true
	}
	fmt.Println("the end is false")
	return false
}