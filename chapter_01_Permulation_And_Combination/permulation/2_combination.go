package main

import "fmt"

/*
   字典序法生成排列 参见 组合数学 Page 18
   [1, n]
 */
func CombinationGenrator(n int, r int) (c chan []int) {
	c = make(chan []int)
	var (
		currCom = make([]int, r)
		count int
	)
	for i:=1; i<=r; i++ {
		currCom[i-1] = i
	}

	go func() {
		newCom := make([]int, r)
		copy(newCom, currCom)
		count++
		c <- newCom

		for {

			var (
				j, i int
			)

			// S1: 求满足 cj < n-r+j, 使j的值达到最大，设 i = max{j | cj < n-r+j}. 如果找不到，中止.
			for j=r-1; j>=0; j-- {
				if currCom[j] < n-r+j+1 {
					i = j
					break
				}
				if j == 0 { // 没有找到
					close(c)
					fmt.Println("n, r, count: ", n, r, count)
					return
				}
			}

			// S2: ci <- (ci)+1
			currCom[i] = currCom[i] + 1

			// S3: cj <- (cj-1) + 1, j = i+1, i+2, ..., r
			for k:=i+1; k<=r-1; k++ {
				currCom[k] = currCom[k-1] + 1
			}

			newCom := make([]int, r)
			copy(newCom, currCom)
			count++
			c <- newCom

		}

	}()
	return
}

//func main() {
//	c := CombinationGenrator(10, 3)
//	for i := range c {
//		fmt.Println(i)
//	}
//}
