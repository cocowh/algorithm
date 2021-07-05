package easy

func replaceSpace(s string) string {
	res := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res = append(res, []byte("%20")...)
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
