/*
 * @Description: 不相交集
 * @Author: longerQiu
 * @Date: 2019-11-10 11:43:59
 * @LastEditTime: 2019-11-24 22:32:17
 */
package dis_set

type DisSet struct {
	size int
	union []int
}

func NewDisSet() *DisSet {
	return new(DisSet)
}

func (ds *DisSet) CreateDisSet(size int)  {
	ds.union = make([]int, size+1)
}

func (ds *DisSet) Find(x int) int  {
	if ds.union[x] <= 0 {
		return x
	} else {
		return ds.Find(ds.union[x])
	}
}

func (ds *DisSet) SetUnion(r1, r2 int)  {
	ds.union[r2] = r1
}

func (ds *DisSet) SetUnionBySize(r1, r2 int) {
	if ds.union[r2] < ds.union[r1] {
		ds.union[r1] = r2
	} else {
		if ds.union[r1] == ds.union[r2] {
			ds.union[r1]--	
		}

		ds.union[r2] = r1
	}
}

func (ds *DisSet) GetUnion() []int {
	return ds.union
}