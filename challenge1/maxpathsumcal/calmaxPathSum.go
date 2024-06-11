package maxpathsumcal

// func เพื่อเทียบค่า a กับ b เลือกค่ามากกว่า
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxPathSum(triangle [][]int) int {
	// n เป็นจำนวนสมาชิกของ two dimensional array(อันนี้คือจำนวน rows ทางลึก)
	n := len(triangle)

	// สร้าง instance dp เป็น two dimensional array ที่มีขนาดเท่ากับ triangle
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
		copy(dp[i], triangle[i]) // คัดลอกค่าจาก triangle ไปยัง dp
	}

	// ใช้วิธีBottom-up เพื่อค่อยๆเทียบมากกว่าของคู่และพอกค่าขึ้นมา
	for i := n - 2; i >= 0; i-- { // เริ่มจากรวบแถวรองสุดท้ายขึ้นมาเรื่อยๆ
		//ต้องเป็น n-2 เพราะลำดับแถวก่อนสุดท้าย -1 และลำดับคือการหัก -1 จาก len
		for j := 0; j <= i; j++ { // เดินทางจากซ้ายไปขวาในแต่ละแถว
			// อัปเดตค่า triangle โดยหาค่าสูงสุดจากลูกทั้งสองของแต่ละจุด
			dp[i][j] += max(dp[i+1][j], dp[i+1][j+1])
			//dp[i+1]คือแถวล่างกว่า1ขั้นที่ถูกยกมารวบ
			//่j กับ j+1 เป็น ลำดับที่ใกล้กันที่เราต้องเทียบกันว่าเอาอันไหนเพื่อได้มากกว่า
			//พอรวบขึ้นมา แถวล่างสุดของloopจะยังอยู่ แต่มีการนำค่าที่เทียบค่ามากกว่ามาทบรวมกับก้อนแถวปัจจุบัน ทำให้แถวปัจจุบันค่ามากขึ้น
		}
	}

	return dp[0][0]
}
