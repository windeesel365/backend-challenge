package jsonhandler

import (
	"encoding/json"
	"log"
	"os"
)

func ReadJSONFile(filename string) [][]int {
	//เช็คการเปิด JSON file ว่าเปิดได้
	//เพื่อvalidity ขั้นแรก
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	//ถ้า validity ด้านบนผ่าน; อ่านข้อมูลจาก JSON file
	byteValue, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	// slice int 2 มิติ เพื่อเก็บ data เพราะข้อมูลเรามีทางลึกและทางกว้าง
	var data [][]int

	// แปลง JSON เป็น slice ยัดใส่ data 2d-slice ต้องpointer
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
	}

	return data
}
