package GoIO

import (
	"bufio"
	"os"
	"path/filepath"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2021/12/21 10:38
 * @Email 1328919715@qq.com
 * @Description:
 **/

/*
  about file
*/

/*
  Seek设置下一次读/写的位置。offset为相对偏移量，
  whence决定相对位置：0为相对文件开头，1为相对当前位置，2为相对文件结尾
  返回新的偏移量（相对开头）和可能的错误
*/
var (
	filePath       = "./test.log"
	seek     int64 = 10
	whence         = 0
)

func Test_readFile_bySeek(t *testing.T) {
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		t.Fatal(err)
	}
	os.Remove(filepath.Clean(absFilePath))

	fd, err := os.Create(filepath.Clean(absFilePath))
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()
	fdW := bufio.NewWriter(fd)
	fdW.WriteString("1234567890")
	fdW.WriteString("12345")
	fdW.WriteString("67890")
	fdW.Flush()

	ret, err := fd.Seek(seek, whence)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
	data := make([]byte, 0, 20)
	readN, err := fd.ReadAt(data, ret)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("ReadN:%d Data:%v", readN, string(data))
}
