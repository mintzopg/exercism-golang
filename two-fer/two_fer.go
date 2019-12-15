package twofer

// ShareWith(name string) string
// ...return a phrase depending on given string name
func ShareWith(name string) string {
	if len(name) == 0 {
		return "One for you, one for me."
	} else {
		return "One for " + name + ", one for me."
	}
}
