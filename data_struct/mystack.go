package datastruct

import (
	"sync"
)

//Stack 栈结构实现
type Stack struct {
	length  int
	topNode *node
	lock    *sync.RWMutex
}

type node struct {
	value   interface{}
	preNode *node
}

//NewStack 创建新栈
func NewStack() *Stack {
	return &Stack{lock: &sync.RWMutex{}}
}

//Push 压入元素
func (s *Stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	newNode := &node{value: v, preNode: s.topNode}
	s.topNode = newNode
	s.length++
}

//Pop 取栈顶元素
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.length == 0 {
		return nil
	}
	top := s.topNode
	s.topNode = top.preNode
	s.length--
	return top.value
}

//Peek 查看栈顶元素
func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.topNode.value
}

//Length 查看栈元素个数
func (s *Stack) Length() int {
	return s.length
}
