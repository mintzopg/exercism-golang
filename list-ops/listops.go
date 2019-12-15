package listops

// IntList of integers equipped with functional methods
type IntList []int
type unaryFunc func(int) int
type binFunc func(int, int) int
type predFunc func(int) bool

// Length method () -> int; returns the length of IntList
func (v IntList) Length() int {
	len := 0
	for range v {
		len++
	}
	return len
}

// Append method(IntList) -> IntList; returns a new list with v and appended elements
func (v IntList) Append(l IntList) IntList {
	u := make(IntList, v.Length()+l.Length())
	for i := 0; i < u.Length(); i++ {
		if i < v.Length() {
			u[i] = v[i]
		} else {
			u[i] = l[i-v.Length()]
		}
	}
	return u
}

// Foldl method (binFunc) -> int; Fold left method
func (v IntList) Foldl(fn binFunc, acc int) int {
	for i := 0; i < v.Length(); i++ {
		acc = fn(acc, v[i])
	}
	return acc
}

// Foldr method (binFun) -> int; Fold right method
func (v IntList) Foldr(fn binFunc, acc int) int {
	for i := v.Length() - 1; i >= 0; i-- {
		acc = fn(v[i], acc)
	}
	return acc
}

// Map method (unaryFunc) -> List; returns a new list with unaryFunc applied on each element  of the old list
func (v IntList) Map(fn unaryFunc) IntList {
	// make a new IntList with same Length as IntList and return it
	u := make(IntList, v.Length())
	for i := 0; i < u.Length(); i++ {
		u[i] = fn(v[i])
	}
	return u
}

// Filter method (predFunc) -> IntList; returns a new list with elements satisfying the predicate
func (v IntList) Filter(fn predFunc) IntList {
	// make a new empty IntList append elements and return it
	u := IntList{}
	for i := 0; i < v.Length(); i++ {
		if fn(v[i]) {
			u = u.Append(IntList{v[i]})
		}
	}
	return u
}

// Reverse method () -> IntList; returns a new list with reversed elements
func (v IntList) Reverse() IntList {
	u := IntList{}
	for i := v.Length() - 1; i >= 0; i-- {
		u = u.Append(IntList{v[i]})
	}
	return u
}

// Concat method ([]IntList) -> IntList; returns a new list with the concatenation of "v" and "lists"
func (v IntList) Concat(lists []IntList) IntList {
	u := IntList{}.Append(v)
	for _, l := range lists {
		u = u.Append(l)
	}
	return u
}
