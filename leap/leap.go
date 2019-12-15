package leap

// IsLeapYear function(int)bool
func IsLeapYear(n int) bool {
	if n%4 == 0 {
		if n%100 == 0 {
			if n%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
