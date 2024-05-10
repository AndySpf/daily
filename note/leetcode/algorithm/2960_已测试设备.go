package algorithm

func countTestedDevices(batteryPercentages []int) int {
	var testCount int
	for _, battery := range batteryPercentages {
		if battery-testCount >= 1 {
			testCount++
		}
	}
	return testCount
}
