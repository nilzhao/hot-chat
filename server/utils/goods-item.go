package utils

import (
	"math/rand"
	"time"
)

const min = int64(1)

// 二倍均值算法
// 红包的数量，红包金额
// 金额单位为分，1元钱=100分
func DoubleAverage(count, amount int64) int64 {
	if count <= 0 {
		return 0
	}
	if count == 1 {
		return amount
	}
	// 计算醉倒可用金额
	max := amount - min*count
	avg := max / count
	avg = 2*avg + min
	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(avg) + min
	return x
}
