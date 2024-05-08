package algorithm

func wateringPlants(plants []int, capacity int) int {
	if len(plants) == 0 {
		return 0
	}
	var steps, currentCap = 0, capacity

	for index := -1; index < len(plants); {
		if currentCap >= plants[index+1] {
			currentCap -= plants[index+1]
			steps++
			index++
		} else {
			steps += 2 * (index + 1)
			currentCap = capacity
		}

		if index >= len(plants)-1 {
			break
		}
	}
	return steps
}
