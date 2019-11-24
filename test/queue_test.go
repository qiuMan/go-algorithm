/*
 * @Description: 
 * @Author: longerQiu
 * @Date: 2019-11-23 11:04:24
 * @LastEditTime: 2019-11-24 10:23:30
 */
 package test

 import (
	 "testing"
	 "go-algorithm/structure"
 )
 
 func TestQueue(t *testing.T) {
	 q := structure.NewQueue()
	 q.EnQueue(0)
	 q.EnQueue(1)
	 q.EnQueue(2)
	 q.EnQueue(3)
	 q.EnQueue(4)
	 
	 t.Log(q.GetQueueLen())
	 
	 for !q.IsEmpty() {
		 e, _ := q.DeQueue()
		t.Log(e) 
	 }

	 t.Log(q.GetQueueLen())
	 
 }