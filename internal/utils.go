package internal

func RepeatSlice(slice []string, n int) []string {
	ret := make([]string, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, slice...)
	}
	return ret
}
