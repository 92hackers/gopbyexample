package foo

func ReverseMap(m map[string]int) map[int]string {
	return func() (_gop_ret map[int]string) {
		_gop_ret = map[int]string{}
		for k, v := range m {
			_gop_ret[v] = k
		}
		return
	}()
}
