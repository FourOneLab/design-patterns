package object_oriented

import "fmt"

const DEFAULT_CAPACITY = 10

// 多态这种机制需要编程语言提供特殊的语法机制来实现。
// - 1、通过接口实现对不同对象的引用
// - 2、通过组合将 DynamicArray 的属性传递给 SortedDynamicArray
// - 3、通过覆盖在 SortedDynamicArray 中重写 DynamicArray 中的 Add 方法
type Array interface {
	Add(e int)
	Get(i int) int
}

type DynamicArray struct {
	size     int
	capacity int
	elements []int
}

func NewDynamicArray() Array {
	return &DynamicArray{
		size:     0,
		capacity: DEFAULT_CAPACITY,
		elements: make([]int, 0),
	}
}

func (a *DynamicArray) Size() int {
	return a.size
}

func (a *DynamicArray) Get(index int) int {
	return a.elements[index]
}

func (a *DynamicArray) Add(e int) {
	a.ensureCapacity()
	a.elements = append(a.elements, e)
}

// ensureCapacity 如果数组满了就扩容
func (a *DynamicArray) ensureCapacity() {}

type SortedDynamicArray struct {
	DynamicArray
}

func NewSortedDynamicArray() Array {
	return &SortedDynamicArray{DynamicArray{
		size:     0,
		capacity: DEFAULT_CAPACITY,
		elements: make([]int, DEFAULT_CAPACITY),
	}}
}

func (s *SortedDynamicArray) Add(e int) {
	s.ensureCapacity()

	i := s.size - 1
	// 寻找合适的插入位置
	for ; i >= 0; i-- {
		if s.elements[i] > e {
			s.elements[i+1] = s.elements[i]
		} else {
			break
		}
	}
	s.elements[i+1] = e
	s.size++
}

func (s *SortedDynamicArray) Get(index int) int {
	return s.elements[index]
}

// ---------------更简单一些的示例----------------

// Iterator 是一个接口，定义来一个可以遍历集合数据的迭代器
type Iterator interface {
	hasNext() bool
	next() string
	remove() string
}

// MyArray 实现 Iterator
type MyArray struct {
	data []string
}

func (m MyArray) hasNext() bool {
	return false
}

func (m MyArray) next() string {
	return ""
}

func (m MyArray) remove() string {
	return ""
}

type LinkedListNode struct {
	value string
	next  *LinkedListNode
}

// LinkedList 实现 Iterator
type LinkedList struct {
	head LinkedListNode
}

func (l LinkedList) hasNext() bool {
	return false
}

func (l LinkedList) next() string {
	return l.head.next.value
}

func (l LinkedList) remove() string {
	return l.next()
}

// prints 传入不同类型的实现类，动态调用 hasNext 函数
func prints(iterator Iterator) {
	if iterator.hasNext() {
		fmt.Println(iterator.next())
	}
}
