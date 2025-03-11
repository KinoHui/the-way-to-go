package myheap

import (
	"container/heap"
	"fmt"
)

// Go 语言中可以通过实现 heap.Interface 来构建整数大顶堆
// 实现 heap.Interface 需要同时实现 sort.Interface
type intHeap []any

// Push heap.Interface 的方法，实现推入元素到堆
func (h *intHeap) Push(x any) {
	// Push 和 Pop 使用 pointer receiver 作为参数
	// 因为它们不仅会对切片的内容进行调整，还会修改切片的长度。
	*h = append(*h, x.(int))
}

// Pop heap.Interface 的方法，实现弹出堆顶元素
// 在 Go 的 heap 包中，当你调用 heap.Pop(h) 时，实际上是先通过堆调整过程，将堆顶元素与最后一个元素交换，然后再移除并返回最后一个元素。
// Pop 方法的职责：Pop 方法的职责是移除并返回最后一个元素。这个时候，堆的调整已经完成了，也就是说堆顶元素已经被交换到最后，并且堆的其他部分已经被调整为一个合法的堆。
func (h *intHeap) Pop() any {
	// 待出堆元素存放在最后
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

// Len sort.Interface 的方法
func (h *intHeap) Len() int {
	return len(*h)
}

// Less sort.Interface 的方法
func (h *intHeap) Less(i, j int) bool {
	// 如果实现小顶堆，则需要调整为小于号
	return (*h)[i].(int) > (*h)[j].(int)
}

// Swap sort.Interface 的方法
func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Top 获取堆顶元素
func (h *intHeap) Top() any {
	return (*h)[0]
}

/* Driver Code */
func TestHeap() {
	/* 初始化堆 */
	// 初始化大顶堆
	maxHeap := &intHeap{}
	heap.Init(maxHeap)
	/* 元素入堆 */
	// 调用 heap.Interface 的方法，来添加元素
	heap.Push(maxHeap, 1)
	heap.Push(maxHeap, 3)
	heap.Push(maxHeap, 2)
	heap.Push(maxHeap, 4)
	heap.Push(maxHeap, 5)

	/* 获取堆顶元素 */
	top := maxHeap.Top()
	fmt.Printf("堆顶元素为 %d\n", top)

	/* 堆顶元素出堆 */
	// 调用 heap.Interface 的方法，来移除元素
	heap.Pop(maxHeap) // 5
	heap.Pop(maxHeap) // 4
	heap.Pop(maxHeap) // 3
	heap.Pop(maxHeap) // 2
	heap.Pop(maxHeap) // 1

	/* 获取堆大小 */
	size := len(*maxHeap)
	fmt.Printf("堆元素数量为 %d\n", size)

	/* 判断堆是否为空 */
	isEmpty := len(*maxHeap) == 0
	fmt.Printf("堆是否为空 %t\n", isEmpty)
}
