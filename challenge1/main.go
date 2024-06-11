package main

import (
	"fmt"
	"maxpathsummod/jsonhandler"
	"maxpathsummod/maxpathsumcal"
)

func main() {

	data1 := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}
	result1 := maxpathsumcal.MaxPathSum(data1)
	fmt.Println(result1)

	filename := "./jsonhandler/hard.json"
	data2 := jsonhandler.ReadJSONFile(filename)
	//เรียก func maxPathSum ในpackage maxpathsumcal
	result2 := maxpathsumcal.MaxPathSum(data2)
	fmt.Println(result2)
}
