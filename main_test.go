package main

import "testing"

func TestFindSum(t *testing.T) {
	data := []struct {
		size        int
		expectedSum int
	}{
		{5, 15},
		{6, 21},
	}

	for _, testData := range data {
		actualSum := FindSum(testData.size)
		if actualSum == testData.expectedSum {
			t.Log("SUCCESS")
		} else {
			t.Errorf("Failure, expected %d, got %d", testData.expectedSum, actualSum)
		}
	}
}
