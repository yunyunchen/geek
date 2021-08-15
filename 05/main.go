package main

import (
	"fmt"
	"sync"
	"time"
)

// SlidingWindowLimiter 滑动窗口计数器限流
type SlidingWindowLimiter struct {
	Interval    int64       // 总计数时间，*1000，转换为豪秒
	LastTime    int64       // 上一个窗口时间,time.Now().UnixNano()
	Lck         *sync.Mutex // 锁
	WinCount    []int64     // 窗口中请求当前数量
	TicketSize  int64       // 窗口最大容量
	TicketCount int64       // 窗口个数
	CurIndex    int64       // 目前使用的窗口下标
}

// NewSlidingWindowLimiter 初始化滑动窗口计数器限流
func NewSlidingWindowLimiter(interval int64, ticketCount int64, ticketSize int64) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		Interval:    interval * 1000,
		LastTime:    time.Now().UnixNano(),
		TicketSize:  ticketSize,
		TicketCount: ticketCount,
		WinCount:    make([]int64, ticketCount, ticketCount),
		CurIndex:    0,
		Lck:         new(sync.Mutex),
	}
}

// slidingCounterLimit 滑动窗口计数器限流实现
func (r *SlidingWindowLimiter) slidingCounterLimit() bool {
	r.Lck.Lock()
	defer r.Lck.Unlock()
	// 一个窗口的时间段
	eachTicketTime := r.Interval / r.TicketCount
	// 当前时间
	now := time.Now().UnixNano()
	//fmt.Println(now/1e6,r.LastTime/1e6,now/1e6-r.LastTime/1e6,eachTicketTime)
	// 如果间隔时间超过一个窗口的时间 当前窗口置0 指向下一个窗口
	if (now/1e6)-(r.LastTime/1e6) > eachTicketTime {
		r.WinCount[r.CurIndex] = 0
		r.CurIndex = (r.CurIndex + 1) % r.TicketCount
		r.LastTime = now
	}

	// 当前窗口未满则正常计数
	flag := false
	if r.WinCount[r.CurIndex] < r.TicketSize {
		r.WinCount[r.CurIndex]++
		flag = true
	}

	fmt.Printf("当前时间段:%d,窗口中请求当前数量：%d,一个周期内窗口数据：%d", r.CurIndex, r.WinCount[r.CurIndex], r.WinCount)

	return flag
}
func main() {
	fmt.Println(time.Now().Unix())
	//var wg sync.WaitGroup
	// 定义1秒5个时间窗口 每个窗口大小为10  即1秒25个请求
	r := NewSlidingWindowLimiter(1, 5, 5)
	for i := 0; i < 40; i++ {
		//wg.Add(1)
		ok := r.slidingCounterLimit()
		if ok {
			fmt.Println("pass ", i)
		} else {
			fmt.Println("limit ", i)
		}
		time.Sleep(30 * time.Millisecond)
		//wg.Done()
	}
	//wg.Wait()

	fmt.Println(time.Now().Unix())
}
