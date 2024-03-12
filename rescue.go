package main

/*
n = number of chickens
k = length of the roof Superman can carry
positions = positions of chickens on the one-dimensional axis.
*/
func rescue(n int, k int, positions []int) int {
	// Initialize the maximum number of chickens that can be protected
	maxChickens := 0

	// Iterate over the positions of the chickens
	for i := 0; i < n; i++ {
		// Initialize the current number of chickens that can be protected
		currentChickens := 1

		// Iterate over the remaining positions
		for j := i + 1; j < n; j++ {
			// If the current position is within the length of the roof, increment the current number of chickens
			if positions[j]-positions[i] < k {
				currentChickens++
			} else {
				// If the current position is not within the length of the roof, break the loop
				break
			}
		}

		// If the current number of chickens is greater than the maximum, update the maximum
		if currentChickens > maxChickens {
			maxChickens = currentChickens
		}
	}

	// Return the maximum number of chickens that can be protected
	return maxChickens
}
