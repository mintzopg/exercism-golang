package brackets

var valids = map[string]bool{
	"(": true,
	")": true,
	"[": true,
	"]": true,
	"{": true,
	"}": true,
}

func Bracket(s string) bool {
	// stack to hold brackets
	stack := []string{}
	for _, c := range s {
		char := string(c)
		if len(stack) == 0 {
			// if empty stack push char and continue
			if _, ok := valids[char]; ok {
				stack = append(stack, char)
			}
			continue
		}
		// check for matching pairs
		if char == ")" && stack[len(stack)-1] == "(" ||
			char == "]" && stack[len(stack)-1] == "[" ||
			char == "}" && stack[len(stack)-1] == "{" {
			stack = stack[:len(stack)-1] // pop matching char
		} else {
			// if valid char push to stack
			if _, ok := valids[char]; ok {
				stack = append(stack, char)
			}
		}
	}
	return len(stack) == 0
}
