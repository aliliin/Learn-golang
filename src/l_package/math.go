package main

import (
	"fmt"
	"math"
)

func main() {
	// 绝对值计算（float64）返回值也是 float64
	x := math.Abs(-2)
	fmt.Printf("%.1f\n", x) // 2.0

	y := math.Abs(2)
	fmt.Printf("%.1f\n", y) // 2.0

	// 进一法取整
	c := math.Ceil(1.49)
	fmt.Printf("%.1f\n", c) // 2.0
	d := math.Ceil(-3.14)
	fmt.Printf("%.1f\n", d) // -3.0

	// 余弦
	fmt.Printf("%.2f\n", math.Cos(math.Pi)) // -1.00
	// 双曲余弦
	fmt.Printf("%.2f\n", math.Cosh(0)) // 1.00

	// 返回x-y或0的最大值。
	fmt.Printf("%.2f\n", math.Dim(4, 2))  // 2
	fmt.Printf("%.2f\n", math.Dim(-2, 0)) // 0

	// 舍去法取整 (返回小于或等于x的最大整数值)
	fmt.Printf("%.1f\n", math.Floor(1.51)) // 1

	// 返回最大值 float64
	fmt.Printf("%.1f\n", math.Max(2, 3)) // 3.0
	// 返回最小值 float64
	fmt.Printf("%.1f\n", math.Min(2, 3)) // 2.0

	// 取余数
	fmt.Printf("%.1f\n", math.Mod(7, 4)) // 3.0
	fmt.Printf("%.1f\n", math.Mod(8, 4)) // 0.0

	// 指数表达式 (返回 x 的 y 次方的幂)
	fmt.Printf("%.1f\n", math.Pow(2, 4)) // 16.0
	fmt.Printf("%.1f\n", math.Pow10(2))  // 100.0

	// 舍入返回最接近的整数，从零舍入一半。
	fmt.Printf("%.1f\n", math.Round(10.5))  // 11
	fmt.Printf("%.1f\n", math.Round(-10.5)) // -11

	// 平方根
	fmt.Printf("%.1f\n", math.Sqrt(9)) // 3.0
}
