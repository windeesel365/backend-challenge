package stringtonum

import (
	"decodetonum/validitystring"
	"strconv"
	"strings"
)

func Stringtonum(encoded string) string {

	validitystring.CheckString(encoded)

	n := len(encoded) + 1
	nums := make([]int, n)

	// กำหนด ค่าเริ่มต้นให้numsตัวแรก
	nums[0] = 0

	for i := 0; i < len(encoded); i++ {
		var substrbefore string

		if i >= 2 {
			// case ปกติ
			substrbefore = encoded[i-2 : i+1]
		} else if i == 1 {
			// case i แค่ 1 ใส่ "X" นำหน้า ป้องกัน outofbound
			substrbefore = "X" + string(encoded[i-1:i+1])
		} else {
			// case i = 0 ใส่ "XX" นำหน้า ป้องกัน outofbound
			substrbefore = "XX" + string(encoded[i:i+1])
		}

		var substrafter string
		if i <= len(encoded)-2 {
			// case ปกติ
			substrafter = encoded[i:]
		} else if i == len(encoded)-1 {
			// case i แค่ 1 ใส่ "X" นำหน้า ป้องกัน outofbound
			substrafter = string(encoded[i:]) + "X"
		} else {
			// case i = 0 ใส่ "XX" นำหน้า ป้องกัน outofbound
			substrafter = string(encoded[i:]) + "XX"
		}

		switch encoded[i] {
		case 'L':
			//ถ้า string มี net count R มากกว่า L  ตอนลด ค่าเลขต่อไปจึงลดระดับลงลึกไปอีกระดับ // RR ทำให้ก่อนหน้านั้นต้องต่ำมาก
			// เกิด special L ตอนลดลง

			if strings.Count(encoded, "R") > strings.Count(encoded, "L") && strings.Contains(substrbefore, "RRL") {
				nums[i+1] = nums[i] - (strings.Count(encoded[:i], "R") - strings.Count(encoded[i+1:], "L"))

			} else { //กรณี net count  L เท่ากับ R จะขยับปกติ ไม่มีการเพิ่มหรือลดระดับเพิ่มได้
				nums[i+1] = nums[i] - 1
			}

		case 'R':
			//ถ้า string มี net count  L มากกว่า R  ตอนเพิ่ม ค่าเลขต่อไปจึงขึ้นเพิ่มไปอีกระดับ  // LL ทำให้ก่อนหน้านั้นต้องสูงเพื่อให้ LL แล้วลงไปเจอ 0 พอดี
			// เกิด special R ตอนเพิ่มขึ้น
			if strings.Count(encoded, "L") > strings.Count(encoded, "R") && strings.Contains(substrafter, "RLL") {

				nums[i+1] = nums[i] + (-strings.Count(encoded[:i], "R") + strings.Count(encoded[i+1:], "L"))

			} else { //กรณี net count  L เท่ากับ R จะขยับปกติ ไม่มีการเพิ่มหรือลดระดับเพิ่มได้
				nums[i+1] = nums[i] + 1
			}
		case '=':
			nums[i+1] = nums[i]
		}
	}

	//หาค่าต่ำสุดในชุดตัวเลข
	minVal := nums[0]
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
	}

	//ปรับตัวเลขทั้งหมดให้เป็นบวกโดยใช้ค่าเริ่มต้นที่ต่ำสุด
	offset := 0
	if minVal < 0 {
		offset = -minVal //ถ้า minVal ติดลบ จะยิ่งเป็นบวกที่ offset
	}

	for i := range nums {
		nums[i] += offset //เอาค่าoffset ไปบวกให้ทุกๆค่าในnums
	}

	// แปลงตัวเลขเป็นสตริง
	var result strings.Builder
	for _, v := range nums {
		result.WriteString(strconv.Itoa(v))
	}

	return result.String()
}
