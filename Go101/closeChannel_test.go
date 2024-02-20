package Go101

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/9/16 10:50
 * @Desc:
 */
//func foo(c <-chan int) {
//	close(c) // error: 不能关闭单向接收通道
//}

func foo(c chan int) {
	close(c) // error: 不能关闭单向接收通道
}
