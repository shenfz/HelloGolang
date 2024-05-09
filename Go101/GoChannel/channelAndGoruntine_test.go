package ChannelQ

import "fmt"

/**
 * @Author shenfz
 * @Date 2021/8/30 16:51
 * @Email 1328919715@qq.com
 * @Description: channel 和 gor 组合
 **/

// 单个 channel接收数据的，优先使用 for-range ，当channel关闭时会自动退出
func getData(ch chan struct{}) {
	defer fmt.Println("GetData exit")
	for range ch {
		fmt.Println("Get Data ")
	}
}

// 多个chan 多事件处理 使用 for-select
// 如果函数内关闭一个声明为 只读通道 <-chan ，编译出错
// 如果只想关闭一个接收事件的通道，不影响其他，则直接设为 nil ,select 则忽略
func getDatas(ch1 chan struct{}, ch chan int, ch1Stop chan struct{}) {
	for {
		select {
		case _, ok := <-ch1:
			if !ok {
				fmt.Println("getData exit")
				return
			}
			fmt.Println("get Data")

		case <-ch1Stop:
			close(ch)
			return
		case data, ok := <-ch:
			if !ok {
				ch = nil // 从 nil 通道接受，永远阻塞；其次，置于nil,更容易被GC回收
			}
			fmt.Println("GetData Int = ", data)

		default:
		}
	}
}
