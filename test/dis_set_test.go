/*
 * @Description: 
 * @Author: longerQiu
 * @Date: 2019-11-23 11:04:24
 * @LastEditTime: 2019-11-24 22:32:30
 */
 package test

 import (
	 "testing"
	 "go-algorithm/dis_set"
 )
 
 func TestDisSet(t *testing.T) {
	 ds := dis_set.NewDisSet()
	 ds.CreateDisSet(16)
	 ds.SetUnion(1, 2)
	 ds.SetUnion(1, 8)
	 ds.SetUnion(2, 9)
	 ds.SetUnion(2, 6)
	 ds.SetUnion(9, 10)
	 t.Log(ds.Find(10))
	 
	 union := ds.GetUnion()

	 t.Log(union)
 }