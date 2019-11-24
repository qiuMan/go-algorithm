/*
 * @Description: 队列
 * @Author: longerQiu
 * @Date: 2019-11-24 09:15:20
 * @LastEditTime: 2019-11-24 10:22:15
 */
package structure

import "errors"

type Queue struct {
	elementType []interface{}
	size int //设置队列长度
}

func NewQueue() *Queue {
	return new(Queue)	
}

func (q *Queue) SetSize(size int) {
	if len(q.elementType) > size {
		q.elementType = q.elementType[0:size]
	}

	q.size = size
}

func (q *Queue) GetQueueLen() int {
	return len(q.elementType)
}

func (q *Queue) EnQueue(item interface{}) error {
	if q.IsFull() {
		return errors.New("Queue is full")
	}

	q.elementType = append(q.elementType, item)
	return nil
}

func (q *Queue) DeQueue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty") 
	}

	e := q.elementType[0]
	q.elementType = q.elementType[1:len(q.elementType)]
	return e, nil 
}

func (q *Queue) IsFull() bool {
	return q.size > 0 && len(q.elementType) == q.size
}

func (q *Queue) IsEmpty() bool {
	return len(q.elementType) == 0
}

func (q *Queue) MakeEmpty() {
	q.elementType = q.elementType[0:0]
}

