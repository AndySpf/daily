package algorithm

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	if len(plants) == 0 {
		return 0
	}
	a := capacityA
	b := capacityB
	fillCount := 0
	for i, j := 0, len(plants)-1; i <= j; i, j = i+1, j-1 {
		if i == j {
			if a < plants[i] && b < plants[i] {
				fillCount++
			}
			break
		}
		if a < plants[i] {
			a = capacityA
			fillCount++
		}
		a -= plants[i]
		if b < plants[j] {
			b = capacityB
			fillCount++
		}
		b -= plants[j]
	}
	return fillCount
}
