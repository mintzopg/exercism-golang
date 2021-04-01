package railfence

import "strings"

func Encode(s string, r int) string {
	fence := make([][]string, r)
	tmp := make([]string, r)
	inc := true
	i := 0
	for _, ch := range s {
		if i == 0 {
			inc = true
		}
		if i == r-1 {
			inc = false
		}
		fence[i] = append(fence[i], string(ch))
		if inc {
			i++
		} else {
			i--
		}
	}
	for i = 0; i < r; i++ {
		tmp[i] = strings.Join(fence[i], "")
	}
	return strings.Join(tmp, "")
}

func Decode(s string, r int) string {
	var sb strings.Builder
	n := len(s)
	counter := make([]int, r)
	tmp := make([]string, r)
	inc := true
	i := 0
	for range s {
		if i == 0 {
			inc = true
		}
		if i == r-1 {
			inc = false
		}
		counter[i]++
		if inc {
			i++
		} else {
			i--
		}
	}
	for i = 0; i < r; i++ {
		tmp[i] = s[:counter[i]]
		s = s[counter[i]:]
	}
	i = 0
	for j := 0; j < n; j++ {
		if i == 0 {
			inc = true
		}
		if i == r-1 {
			inc = false
		}
		sb.WriteString(string(tmp[i][0]))
		tmp[i] = tmp[i][1:]
		if inc {
			i++
		} else {
			i--
		}
	}
	return sb.String()
}
