//快排的实现

package main

import (
	"fmt"
	"math/rand"
)

const (
	LOG = true
)

func log(params ...interface{}) {
	if LOG {
		fmt.Println(params...)
	}
}

// 分区
func partition(array []int, l int, r int) int {
	povit := array[r]
	c := l
	log(" -- povit:", povit)
	for i := l; i < r; i++ {
		if array[i] <= povit {
			// 交换
			log(fmt.Sprintf(" >> i: %d, c: %d", i, c))
			if i != c {
				log(fmt.Sprintf("  >>  value of i: %d, c: %d", array[i], array[c]))
				array[i], array[c] = array[c], array[i]
				log("  >>  ", array)
			}
			// 游标前移
			c++
		}
	}

	array[r], array[c] = array[c], array[r]
	return c
}

// 排序
func QuickSort(array []int, l int, r int) {
	if l < r {
		povit := partition(array, l, r)
		QuickSort(array, l, povit-1)
		QuickSort(array, povit+1, r)
	}
}

// 排序全部
func QuickSortAll(array []int) {
	QuickSort(array, 0, len(array)-1)
}

func main() {
	var pool = make([]int, 0, 16)
	capcity := cap(pool)
	random := rand.New(rand.NewSource(123))
	for i := 0; i < capcity; i++ {
		pool = append(pool, random.Intn(32))
	}

	fmt.Println("start ===:")
	fmt.Println(pool, "len", len(pool))
	QuickSortAll(pool)
	fmt.Println("end ===:")
	fmt.Println(pool, "len", len(pool))

}
