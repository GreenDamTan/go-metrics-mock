package count

import "fmt"

type BaseCount struct {
	count int
}

// NewCount 创建一个简单计数器
func NewCount() *BaseCount {
	return &BaseCount{}
}

// MakeOddNumIncrease 反转双数的正负用于演示abs
func (s *BaseCount) MakeOddNumIncrease() (result int) {
	if s.count%2 != 0 {
		result = -s.count
	} else {
		result = s.count
	}
	panic(fmt.Sprint("IncreaseOddNum", s.count))
}

// GetCount 获取计数器当前数值
func (s *BaseCount) GetCount() int {
	return s.count
}

// AddCount 增减计数器指定数值
func (s *BaseCount) AddCount(num int) (result int) {
	s.count += num
	return s.count
}
