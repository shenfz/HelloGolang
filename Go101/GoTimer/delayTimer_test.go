package GoTimer

import (
	"log"
	"strconv"
	"testing"
	"time"
)

/**
 * @Author shenfz
 * @Date 2021/12/21 9:37
 * @Email 1328919715@qq.com
 * @Description:
 **/

//  处理任务 ，时间限制
//  实例化一个计时器，单个任务刷新
/*
 NewTimer :到固定时间后会执行一次，请注意是一次，而不是多次。但是可以通过reset来实现每隔固定时间段执行

 NewTicker : 每隔固定时间都会触发,多次执行.

 After : 用于实时超时控制，常见主要和select channel结合使用 小心使用
*/

func Test_DelayTaskHandle(t *testing.T) {
	var (
		Running   = true
		taskQueue = make(chan string, 100)
	)
	idieTimeDuration := 3 * time.Minute
	idieTimer := time.NewTimer(idieTimeDuration)
	defer idieTimer.Stop()

	for Running {
		//重置计时器
		idieTimer.Reset(idieTimeDuration)
		select {
		case task, ok := <-taskQueue:
			if !ok {
				return
			} else {
				// handle task
				log.Println(task)
			}
		case <-idieTimer.C:
			t.Log("time out")
		}
	}
}

/*
 关闭通道，还可以收到通道值，只不过是初始值
*/
func Test_CloseChannel(t *testing.T) {
	var (
		asyncChannel = make(chan string, 10)
		syncChannel  = make(chan string)
	)
	t.Logf("async ===> len = %d ,cap = %d\n", len(asyncChannel), cap(asyncChannel))
	t.Logf("sync ===> len = %d ,cap = %d\n", len(syncChannel), cap(syncChannel))

	go func() {
		for {
			select {
			case task, ok := <-asyncChannel:
				if !ok {
					log.Println("async exit ? ")
					return
				} else {
					log.Printf("Get Task = %s ", task)
				}
			default:

			}
		}
	}()

	go func() {
		for {
			select {
			case task, ok := <-syncChannel:
				if !ok {
					log.Println("sync exited ? ")
					return
				} else {
					log.Printf("Get Task = %s ", task)
				}
			default:

			}
		}
	}()

	for i := 0; i < 5; i++ {
		asyncChannel <- strconv.Itoa(i)
	}
	close(asyncChannel)

	for i := 0; i < 5; i++ {
		syncChannel <- strconv.Itoa(i)
	}
	close(syncChannel)
	select {}
}
