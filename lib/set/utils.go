package set

var (
	IntKeyFunc = func(elem interface{}) int {
		return elem.(int)
	}
)
