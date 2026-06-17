// پروژه‌ی کوه ستاره
// این برنامه یک «کوه» (هرم متقارن) از ستاره‌ها را در خروجی چاپ می‌کند.
// ارتفاع کوه را کاربر تعیین می‌کند.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("ارتفاع کوه را وارد کنید: ")
	input, _ := reader.ReadString('\n')

	height, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || height <= 0 {
		fmt.Println("خطا: لطفاً یک عدد صحیح و مثبت وارد کنید.")
		return
	}

	printMountain(height)
}

// printMountain یک هرم متقارن از ستاره‌ها با ارتفاع داده‌شده چاپ می‌کند.
// در هر ردیف:
//
//	تعداد فاصله‌های سمت چپ = height - row
//	تعداد ستاره‌ها        = 2*row - 1  (همیشه فرد، تا کوه متقارن شود)
func printMountain(height int) {
	for row := 1; row <= height; row++ {
		spaces := strings.Repeat(" ", height-row)
		stars := strings.Repeat("*", 2*row-1)
		fmt.Println(spaces + stars)
	}
}
