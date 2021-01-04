package solution11

func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		privot := low + (high-low)>>1
		if numbers[privot] < numbers[high] {
			high = privot
		} else if numbers[privot] > numbers[high] {
			low = privot + 1
		} else {
			high--
		}
	}
	return numbers[low]
}
