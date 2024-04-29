package algorithm

type changeNode struct {
	snap  int
	value int
}

type SnapshotArray struct {
	lastChange [][]changeNode
	change     map[int]int
	snap       int
}

// 第一种 内存不足
// [[1, 2, 3, 4] [1, 0, 3, 4]]

// 第二种 用map记录改变的为止，内存降了，耗时较大
//	   1, 2 -> change 0 1->1, 2->2
// [1, 2, 3, 4]

// 第三种 双向链表
// nochange 0: snap3
// change   1: snap1->2

// Constructor Constructor
func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		lastChange: make([][]changeNode, length),
		change:     map[int]int{},
		snap:       0,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.change[index] = val
}

func (this *SnapshotArray) Snap() int {
	if len(this.change) > 0 {
		for k, v := range this.change {
			this.lastChange[k] = append(this.lastChange[k], changeNode{
				snap:  this.snap,
				value: v,
			})
		}
		this.change = map[int]int{}
	}
	this.snap++
	return this.snap - 1
}

// 2,3,5,8,10
func (this *SnapshotArray) Get(index int, snap_id int) int {
	if snap_id >= this.snap {
		return -1
	}

	start := 0
	end := len(this.lastChange[index])
	if len(this.lastChange[index]) == 0 {
		return 0
	}
	if snap_id < this.lastChange[index][0].snap {
		return 0
	}
	if snap_id > this.lastChange[index][end-1].snap {
		return this.lastChange[index][end-1].value
	}
	for end > start {
		middle := (end-start)/2 + start
		if snap_id >= this.lastChange[index][middle].snap && snap_id < this.lastChange[index][middle+1].snap {
			return this.lastChange[index][middle].value
		}

		if snap_id < this.lastChange[index][middle].snap {
			end = middle
		} else {
			start = middle + 1
		}
	}
	return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
