package secret

// Handshake converts integer input into slice of strings
func Handshake(code int) (ret []string) {
	if code <= 0 || code >= 32 {
		return nil
	}

	for _, i := range []int{1, 2, 4, 8, 16} {
		switch code & i {
		case 1:
			ret = append(ret, "wink")
		case 2:
			ret = append(ret, "double blink")
		case 4:
			ret = append(ret, "close your eyes")
		case 8:
			ret = append(ret, "jump")
		case 16:
			ret = reverse(ret)
		}
	}

	return ret
}

func reverse(s []string) (r []string) {
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}
