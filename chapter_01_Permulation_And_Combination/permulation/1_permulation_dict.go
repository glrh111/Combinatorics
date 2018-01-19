package main

import "fmt"

/*
   字典序法生成排列 参见 组合数学 Page 18
   [1, n]
 */
func GeneratorDict(n int) (c chan []int) {
	c = make(chan []int)
	var (
		firstPer = make([]int, n)
		nextPer = make([]int, n)
		count int
	)
	for i:=1; i<=n; i++ {
		firstPer[i-1] = i
	}

	go func() {
		// 每次push到里边的时候，需要注意，copy一份出来. 不然，接收者使用的时候，会变
		newPer := make([]int, n)
		copy(newPer, firstPer)
		count++
		c <- newPer

		for {
			// 通过 firstPer 生成 nextPer

			var (
				j, i, h, k int
			)

			// S1: 求满足关系式 p_(j-1) < p_j 的最大值，设为 i，即 i = max{j | p_(j-1) < p_j}
			for j=n-1; j>=1; j-- {
				if firstPer[j] > firstPer[j-1] {
					i = j+1
					break
				}
			}
			if i == 0 { // 没找到
				close(c)
				fmt.Println("n count: ", n, count)
				return
			}

			// S2: 求满足关系式 p_(i-1) < p_k 的k的最大值，设为 h, 即 h = max{k | p_(i-1) < p_k}
			for k=n-1; k>=0; k-- {
				if firstPer[k] > firstPer[i-2] {
					h = k + 1
					break
				}
			}

			// S3: p_(i-1)与p_h 互换，得到 p1'p2'...pn'
			firstPer[i-2], firstPer[h-1] = firstPer[h-1], firstPer[i-2]

			// S4: 令 p1'p2'...p_(i-1)'p_i'p_(i+1)'...p_n 中 i ~ n 的顺序逆转，即得到下一个排列
			copy(nextPer, firstPer)

			for y:=i-1; y<=n-1; y++ {
				nextPer[y] = firstPer[n+i-2-y] // 反转
			}

			newPer := make([]int, n)
			copy(newPer, nextPer)
			c <- newPer
			count++

			copy(firstPer, nextPer)
		}
	}()
	return
}

//func main() {
//	c := GeneratorDict(6)
//	for i := range c {
//		fmt.Println(i)
//	}
//}

