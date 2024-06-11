package maxpathsumcal

import (
	"maxpathsummod/jsonhandler"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		triangle [][]int
		expected int
	}{
		{
			triangle: [][]int{
				{59},
				{73, 41},
				{52, 40, 53},
				{26, 53, 6, 34},
			},
			expected: 237,
		},
		{
			triangle: [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			},
			expected: 21,
		},
		{
			triangle: [][]int{
				{1},
				{2, 3},
				{4, 5, 6},
			},
			expected: 10,
		},
		{
			triangle: [][]int{
				{10},
			},
			expected: 10,
		},
	}

	//เราใช้ hard.json จาก foler jsonhandler นำมา unit test
	filename := "../jsonhandler/hard.json"
	data := jsonhandler.ReadJSONFile(filename)

	//appendส่วน hard.jsonเข้าสู่ tests
	tests = append(tests, struct {
		triangle [][]int
		expected int
	}{
		triangle: data,
		expected: 7273,
	})

	for _, test := range tests {
		result := MaxPathSum(test.triangle)
		if result != test.expected {
			t.Errorf("MaxPathSum(%v) = %d; expected %d", test.triangle, result, test.expected)
		}
	}
}
