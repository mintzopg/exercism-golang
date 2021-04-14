package binarysearch

func SearchInts(a []int, key int) int {
	return binarySearch(a, 0, len(a)-1, key)
}

func binarySearch(a []int, low, high, x int) int {
	if high >= low {
		middle := int(low + (high-low)/2)
		if a[middle] == x { // x is in the middle
			return middle
		}
		if a[middle] > x { // go left
			return binarySearch(a, low, middle-1, x)
		}
		return binarySearch(a, middle+1, high, x) // go right
	}
	return -1
}
