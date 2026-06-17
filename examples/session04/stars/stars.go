// Session 04 (extra) — الگوهای ستاره با حلقه‌های تو در تو (کوه ستاره).
// Star patterns built with nested for-loops. A classic loop exercise.
// Run:  go run examples/session04/stars/stars.go
package main

import (
	"fmt"
	"strings"
)

// rightTriangle: یک مثلث راست‌گوشه.
// row 1 -> "*", row 2 -> "**", ...
func rightTriangle(height int) {
	for row := 1; row <= height; row++ {
		for star := 0; star < row; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// pyramid: کوه ستاره (هرم وسط‌چین).
// هر ردیف = چند فاصله برای وسط‌چین کردن + ستاره‌های فرد (1, 3, 5, ...).
func pyramid(height int) {
	for row := 1; row <= height; row++ {
		spaces := height - row // فاصله‌های سمت چپ برای وسط‌چین شدن
		stars := 2*row - 1     // تعداد ستاره‌های فرد در هر ردیف
		fmt.Print(strings.Repeat(" ", spaces))
		fmt.Print(strings.Repeat("*", stars))
		fmt.Println()
	}
}

// diamond: لوزی = یک هرم رو به بالا + یک هرم وارونه.
func diamond(height int) {
	// نیمه‌ی بالا (شامل وسط)
	for row := 1; row <= height; row++ {
		fmt.Print(strings.Repeat(" ", height-row))
		fmt.Println(strings.Repeat("*", 2*row-1))
	}
	// نیمه‌ی پایین (وارونه)
	for row := height - 1; row >= 1; row-- {
		fmt.Print(strings.Repeat(" ", height-row))
		fmt.Println(strings.Repeat("*", 2*row-1))
	}
}

func main() {
	const height = 5

	fmt.Println("۱) مثلث راست‌گوشه (right triangle):")
	rightTriangle(height)

	fmt.Println("\n۲) کوه/هرم ستاره (pyramid):")
	pyramid(height)

	fmt.Println("\n۳) لوزی (diamond):")
	diamond(height)
}
