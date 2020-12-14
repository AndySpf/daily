package algorithm

// RDRDD  => R赢
func predictPartyVictory(senate string) string {
	m := map[byte]int{}
	for i := range senate {
		m[senate[i]]++
	}

	// 每个议员一定限制自己之后对方第一个行使权力的议员
	bs := []byte(senate)
	RDel, DDel := 0, 0
	for m['D'] != 0 && m['R'] != 0 {
		for i := range bs {
			if bs[i] == 'N' {
				continue
			}

			if bs[i] == 'R' {
				if RDel != 0 {
					RDel--
					bs[i] = 'N'
					continue
				}
				m['D']--
				if m['D'] <= 0 {
					return "Radiant"
				}
				DDel++
			}

			if bs[i] == 'D' {
				if DDel != 0 {
					bs[i] = 'N'
					DDel--
					continue
				}
				m['R']--
				if m['R'] <= 0 {
					return "Dire"
				}
				RDel++
			}
		}
	}

	if m['D'] == 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}

// 循环队列，双方哪个先发言，则把这个人添加到这一方队列的末尾，然后双方队列各自去掉队首。
func predictPartyVictory1(senate string) string {
	var radiant, dire []int
	for i, s := range senate {
		if s == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	// 每个议员一定限制自己之后对方第一个行使权力的议员
	for len(radiant) > 0 && len(dire) > 0 {
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}
	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
