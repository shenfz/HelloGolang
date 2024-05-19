package t_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

/**
 * @Author shenfz
 * @Date 2024/3/29 13:07
 * @Email 1328919715@qq.com
 * @Description: 善用io-copy 替代 ioutil。ReadAll ,后者起始大小512字节，存在多次扩容，耗时较大，指定的内存大小是需要读取的数据大小的两倍时，效率达到最高
 **/

var (
	filename           = "./test.log"
	enableDoubleBuffer = true
)

func Test_IoCopy(t *testing.T) {
	absFilePath, _ := filepath.Abs(filename)
	file, err := os.Open(absFilePath)
	if err != nil {
		t.Errorf("open err:%v", err)
		return
	}
	defer file.Close()
	fileInfo, er := file.Stat()
	if er != nil {
		t.Errorf("state err:%v", err)
		return
	}
	var size int64 = fileInfo.Size()
	if enableDoubleBuffer {
		size = size * 2
	}
	buf := bytes.NewBuffer(make([]byte, 0, size))
	_, err = io.Copy(buf, file)
	if err != nil {
		t.Errorf("copy err:%v", err)
		return
	}
}

func TestReadAll(t *testing.T) {
	absFilePath, _ := filepath.Abs(filename)
	file, err := os.Open(absFilePath)
	if err != nil {
		t.Errorf("open err:%v", err)
		return
	}
	_, err = ioutil.ReadAll(file)
	if err != nil {
		t.Errorf("readall err:%v", err)
		return
	}
}
