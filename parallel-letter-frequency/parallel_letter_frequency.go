package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency ([] string) FreqMap; parallel implementation of Frequency
func ConcurrentFrequency(strings []string) FreqMap {
	ch := make(chan FreqMap)
	defer close(ch)
	result := FreqMap{}

	for _, str := range strings {
		go func(s string) {
			ch <- Frequency(s)
		}(str)
	}

	for range strings {
		for r, count := range <-ch {
			result[r] += count
		}
	}
	return result
}
