package utils

import (
	"strconv"
	"time"
)

// 时间耗时计算器
type CC struct {
	unit  time.Duration
	start int64
	stop  int64
}

// 通过一个计时单位
func NewCC(unit time.Duration) *CC {
	c := &CC{
		unit:  unit,
		start: 0,
		stop:  0,
	}
	return c
}

// 以毫秒为单位
func NewDefaultCC() *CC {
	return NewCC(time.Millisecond)
}

// 开始点
func (c *CC) Start() {
	c.start = time.Now().UnixNano()
}

// 结束点
func (c *CC) Stop() {
	c.stop = time.Now().UnixNano()
}

// 耗时, 单位Nano
func (c *CC) Cost() int64 {
	return c.stop - c.start
}

// 得到单位后的
func (c *CC) CostWithUnit() int64 {
	return c.Cost() / int64(c.unit)
}

// 耗时的显示， 会自动转换单位
func (c *CC) CostStr() string {
	cost := c.CostWithUnit()
	return strconv.FormatInt(cost, 10) + c.GetUnitStr()
}

// 获取转换单位后带小数的结果
func (c *CC) CostFloatStr() string {
	cost := float64(c.Cost()) / float64(c.unit)
	return strconv.FormatFloat(cost, 'f', 4, 32) + c.GetUnitStr()
}

// 获取单位
func (c *CC) GetUnit() time.Duration {
	return c.unit
}

// 获取单位的字符串
func (c *CC) GetUnitStr() string {
	switch c.unit {
	case time.Nanosecond:
		return "ns"
	case time.Microsecond:
		return "micros"
	case time.Millisecond:
		return "ms"
	case time.Second:
		return "s"
	case time.Minute:
		return "m"
	case time.Hour:
		return "h"
	}
	return "unkown"
}
