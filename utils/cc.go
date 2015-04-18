package utils

import (
	"strconv"
	"time"
)

// 时间耗时计算器
type CC struct {
	var unit time.Duration
	var start int64
	var stop int64
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
func (c *CC) Cost() int64{
	return c.stop - c.start	
}

// 耗时的显示， 会自动转换单位
func (c *CC) CostStr() string{
	cost := Cost()	
	return strconv.Itoa(cost)	
}
