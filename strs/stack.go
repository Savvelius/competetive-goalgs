package strs

func CorrectBrackets(s string) bool {
	var stack = []byte{}

	for i := range s {
		switch s[i] {
		case '(', '{', '[':
			stack = append(stack, s[i])
		case ')', '}', ']':
			var open byte
			switch s[i] {
			case ')':
				open = '('
			case '}':
				open = '{'
			case ']':
				open = '['
			}

			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top != open {
				return false
			}
			stack = stack[:len(stack)-1]

		}
	}

	return len(stack) == 0
}
