package utils

func IFSelector(is bool, d1 interface{}, d2 interface{}) interface{} {
	if is {
		return d1
	} else {
		return d2
	}
}
