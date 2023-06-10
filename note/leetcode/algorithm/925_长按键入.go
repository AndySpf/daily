package algorithm

func isLongPressedName(name string, typed string) bool {
	if name == "" || typed == "" {
		if name == "" && typed == "" {
			return true
		} else {
			return false
		}
	}
	if name[0] != typed[0] {
		return false
	}

	ptr1, ptr2 := 1, 1
	for {
		if ptr1 == len(name) && ptr2 == len(typed) {
			break
		}
		if ptr1 < len(name) && ptr2 == len(typed) {
			return false
		}
		if ptr1 == len(name) && ptr2 < len(typed) {
			for i := ptr2; i < len(typed)-1; i++ {
				if typed[i] != name[ptr1-1] {
					return false
				}
			}
			return true
		}
		s := name[ptr1]
		if typed[ptr2] == s {
			ptr1++
			ptr2++
			continue
		}
		if typed[ptr2] != s && typed[ptr2] == name[ptr1-1] {
			ptr2++
			continue
		}
		return false
	}
	return true
}

func isLongPressedName1(name string, typed string) bool {
	ptr1, ptr2 := 0, 0
	for ptr2 < len(typed) {
		if ptr1 < len(name) && name[ptr1] == typed[ptr2] {
			ptr1++
			ptr2++
		} else if ptr2 > 0 && typed[ptr2] == typed[ptr2-1] {
			ptr2++
		} else {
			return false
		}
	}
	return ptr1 == len(name)
}
