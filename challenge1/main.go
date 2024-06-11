package main

import (
	"fmt"
	"maxpathsummod/jsonhandler"
	//"summaxpath/jsonhandler"
	//"summaxpath/maxpathsumcal"
)

func main() {

	filename := "./jsonhandler/hard.json"
	data := jsonhandler.ReadJSONFile(filename)
	//เรียก func maxPathSum ในpackage maxpathsumcal
	//result := maxpathsumcal.MaxPathSum(data)

	fmt.Println(result)
}
