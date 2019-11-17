/*
 * @Description: 不相交集
 * @Author: longerQiu
 * @Date: 2019-11-10 11:43:59
 * @LastEditTime: 2019-11-10 13:29:56
 */
package main

import "fmt"

func main()  {
	var s [16]int
	initialize(&s)
	setUnion(&s, 5, 6)
	setUnion(&s, 7, 8)
	setUnion(&s, 5, 7)
	setUnionBySize(&s, 5, 4)
	e := find(6, s)
	fmt.Printf("%d \n", e)
	e = find(4, s)
	fmt.Printf("%d \n", e)
	fmt.Println(s)
}

func initialize(S *[16]int)  {
	var i int
	for i = 15; i >= 0; i-- {
		S[i] = -1
	}
}

func find(x int, S [16]int) int  {
	if S[x] <= 0 {
		return x
	} else {
		return find(S[x], S)
	}
}

func setUnion(S *[16]int, r1, r2 int)  {
	S[r2] = r1
}

func setUnionBySize(S *[16]int, r1, r2 int) {
	if S[r2] < S[r1] {
		S[r1] = r2
	} else {
		if S[r1] == S[r2] {
			S[r1]--
		}
		
		S[r2] = r1
	}
}