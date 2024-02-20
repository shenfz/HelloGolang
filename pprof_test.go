package test

import (
	"github.com/pkg/profile"
	"math/rand"
	"os"
	"runtime/pprof"
	"testing"
	"time"
)

func generate1(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

// go tool pprof -http=:9999 cpu.pprof
func Test_Pprof(t *testing.T) {

	fd, _ := os.OpenFile("./cpu.profile", os.O_CREATE|os.O_RDWR, 0644)
	defer fd.Close()
	pprof.StartCPUProfile(fd)
	defer pprof.StopCPUProfile()
	n := 10
	for i := 0; i < 5; i++ {
		bubbleSort(generate1(n))
		n *= 10
	}
	t.Log("finished")
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func concat(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += randomString(n)
	}
	return s
}

func Test_ProfilePkg(t *testing.T) {
	defer profile.Start().Stop()
	concat(100)
}
