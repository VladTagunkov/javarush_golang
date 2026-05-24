package main

import "fmt"

func main() {
	var steps int
	fmt.Scan(&steps)

	var sum64 float64 = 0.0
	var sum32 float32 = 0

	for i := 0; i < steps; i++ {
		// TODO: прибавьте 0.1 к sum64 (float64) и sum32 (float32) на каждой итерации
		sum64 += 0.1
		sum32 += 0.1
	}

	expected := float64(steps) / 10.0
	// TODO: вычислите ожидаемое значение суммы (expected) для заданного числа шагов

	delta64 := expected * 0
	// TODO: вычислите delta64 как модуль разницы между sum64 и expected через if (без math.Abs)
	if expected >= sum64 {
		delta64 = expected - sum64
	} else {
		delta64 = sum64 - expected
	}

	delta32 := expected * 0
	// TODO: вычислите delta32 как модуль разницы между float64(sum32) и expected через if (без math.Abs)
	if expected >= float64(sum32) {
		delta32 = expected - float64(sum32)
	} else {
		delta32 = float64(sum32) - expected
	}

	fmt.Printf("sum64=%.17f\n", sum64)
	fmt.Printf("sum32=%.17f\n", sum32)
	fmt.Printf("delta64=%.17f\n", delta64)
	fmt.Printf("delta32=%.17f\n", delta32)
}
