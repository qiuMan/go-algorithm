/*
 * @Description: 
 * @Author: longerQiu
 * @Date: 2019-11-23 11:14:09
 * @LastEditTime: 2019-11-23 11:14:09
 */
package test

import (
	"testing"
	"go-algorithm/sort/function"
)

func TestShell(t *testing.T) {
	array := []int{90, 12, 34, 333, 1111, 1, 3, 9, 2, 8, 110, 34, 53, 97, 89, 200}
	n := len(array)
	sort := function.Shell(array, n)
	t.Log(sort)
}