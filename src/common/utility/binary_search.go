package utility

import (
	"sort"
)

func StringBinarySearch(a []string, x string) (status bool) {
	i := sort.Search(len(a), func(i int) bool { return x <= a[i] })
	if i < len(a) && a[i] == x {
		status = true
	} else {
		status = false
	}
	return
}
