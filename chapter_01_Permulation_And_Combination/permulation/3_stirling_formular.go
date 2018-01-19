package main

import (
	"fmt"
	"math/big"
	"math"
)

/*
   实现斯特林公式 Stirling 公式

   n! ~ sqrt(2*n*π) * pow(n/e, n)

 */
func realFactorial(n int) *big.Int {
	re := big.NewInt(1)
	for i:=int64(1); i<=int64(n); i++ {
		re.Mul(re, big.NewInt(i))
	}
	return re
}

// 使用近似法计算 搞得我头疼，先不搞了。
func stirlingFormular(n int) *big.Float {
	re := big.NewFloat(1)

	// 计算右边的一块
	rightRe := big.NewInt(1)
	rightRe.Exp( big.NewInt(int64(float64(n)/math.E)), big.NewInt(int64(n)), nil )

	fmt.Println("rightRe: ", rightRe, big.NewInt(int64(float64(n)/math.E)), big.NewInt(int64(n)),big.NewInt(1) )

	// copy 到float 里边
	re3 := big.NewFloat(1)
	re3.SetInt(rightRe)
	re = re.Mul(big.NewFloat(math.Sqrt(2*float64(n)*math.Pi)), re3)

	return re
}


func main() {
	real1 := realFactorial(123)
	real2 := big.NewFloat(1)
	real2.SetInt(real1)

	jinsi := stirlingFormular(100) //
	fmt.Println(real2, jinsi)

}