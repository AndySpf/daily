package algorithm

//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。
//通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
//
//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
func intToRoman(num int) string {
	MCount, MRemainder := num/1000, num%1000
	DCount, DRemainder := MRemainder/500, MRemainder%500
	CCount, CRemainder := DRemainder/100, DRemainder%100
	LCount, LRemainder := CRemainder/50, CRemainder%50
	XCount, XRemainder := LRemainder/10, LRemainder%10
	VCount, VRemainder := XRemainder/5, XRemainder%5 // 0/1
	ICount := VRemainder / 1

	result := ""
	for i := 0; i < MCount; i++ {
		result += "M"
	}
	for i := 0; i < DCount; i++ {
		if DCount == 1 && CCount == 4 {
			break
		}
		result += "D"
	}
	if CCount == 4 {
		tmp := "CD"
		if DCount == 1 {
			tmp = "CM"
		}
		result += tmp
	} else {
		for i := 0; i < CCount; i++ {
			result += "C"
		}
	}

	for i := 0; i < LCount; i++ {
		if LCount == 1 && XCount == 4 {
			break
		}
		result += "L"
	}
	if XCount == 4 {
		tmp := "XL"
		if LCount == 1 {
			tmp = "XC"
		}
		result += tmp
	} else {
		for i := 0; i < XCount; i++ {
			result += "X"
		}
	}

	for i := 0; i < VCount; i++ {
		if VCount == 1 && ICount == 4 {
			break
		}
		result += "V"
	}
	if ICount == 4 {
		tmp := "IV"
		if VCount == 1 {
			tmp = "IX"
		}
		result += tmp
	} else {
		for i := 0; i < ICount; i++ {
			result += "I"
		}
	}
	return result
}

//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
func romanToInt(s string) int {
	m := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	integer := 0
	for i := 0; i < len(s); {
		if i+2 <= len(s) {
			if v, ok := m[s[i:i+2]]; ok {
				integer += v
				i += 2
				continue
			}
		}

		integer += m[s[i:i+1]]
		i++
	}
	return integer
}
