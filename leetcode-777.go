package main

// 设计链表
// 设计链表的实现，可以选择单链表或者双链表，单链表中的节点应该具有两个属性：val和next
// val是当前节点的值，next是指向下一个节点的指针/引用，如果要使用双链表，则还需要一个属性prev指向前一个节点
// 要支持以下操作：
// 1. get(index)：获取链表中第index个节点的值，如果索引无效，则返回-1
// 2. addAtHead(val)：在链表的第一个元素之前添加一个值为val的节点，插入后，新节点将成为链表的第一个节点
// 3. addAtTail(val)：将值为val的节点追加到链表的最后一个元素
// 4. addAtIndex(index,val)：在链表中的第index个节点之前添加值为val的节点，如果index等于链表的长度，则该节点将附加到链表的末尾 如果index大于链表长度，则不会插入节点，如果index小于0，则在头部插入节点
// 5. deleteAtIndex(index)：如果索引index有效，则删除链表中的第index个节点

// 单链表实现

type MyLinkedNode struct {
	Val  int
	Next *MyLinkedNode
}

type MyLinkedList struct {
	head *MyLinkedNode
	size int
	//next *MyLinkedNode
	//prev *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		&MyLinkedNode{},
		0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	// 当索引无效则退出
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size || index < 0 {
		return
	}
	index = max(index, 0)
	this.size++
	pred := this.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	toAdd := &MyLinkedNode{val, pred.Next}
	pred.Next = toAdd
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index > this.size || index < 0 {
		return
	}
	this.size--
	pred := this.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	pred.Next = pred.Next.Next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
