package heap

var (
	IntCompFunc = func(first, second interface{}) int {
		firstInt :=first.(int)
		secondInt := second.(int)
		if firstInt < secondInt {
			return -1
		} else if firstInt > secondInt {
			return 1
		} else {
			return 0
		}
	}
)