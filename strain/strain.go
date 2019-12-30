package strain

// Collections implemented on
type Ints []int
type Lists [][]int
type Strings []string

// Keep (predicate func) Ints // applies predicate function on collection and returns a new collection
func (ints Ints) Keep(fn func(i int) bool) Ints {
	var out Ints
	for _, val := range ints {
		if fn(val) {
			out = append(out, val)
		}
	}
	return out
}

// Discard (predicate func) Ints // applies predicate function on collection and returns a new collection
func (ints Ints) Discard(fn func(i int) bool) Ints {
	var out Ints
	for _, val := range ints {
		if !fn(val) {
			out = append(out, val)
		}
	}
	return out
}

// Keep (predicate func) Lists // applies predicate function on collection and returns a new collection
func (lists Lists) Keep(fn func(l []int) bool) Lists {
	var out Lists
	for _, val := range lists {
		if fn(val) {
			out = append(out, val)
		}
	}
	return out
}

// Discard (predicate func) Lists // applies predicate function on collection and returns a new collection
func (lists Lists) Discard(fn func(l []int) bool) Lists {
	var out Lists
	for _, val := range lists {
		if !fn(val) {
			out = append(out, val)
		}
	}
	return out
}

// Keep (predicate func) Strings // applies predicate function on collection and returns a new collection
func (str Strings) Keep(fn func(s string) bool) Strings {
	var out Strings
	for _, val := range str {
		if fn(val) {
			out = append(out, val)
		}
	}
	return out
}

// Discard (predicate func) Strings // applies predicate function on collection and returns a new collection
func (str Strings) Discard(fn func(s string) bool) Strings {
	var out Strings
	for _, val := range str {
		if !fn(val) {
			out = append(out, val)
		}
	}
	return out
}
