package list

//代表一个节点
type MyListNode struct{
	Next *MyListNode
	Value interface{}
}
type List struct{
	head *MyListNode
	len int
}