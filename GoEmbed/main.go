package main

/**
 * @Author shenfz
 * @Date 2024/3/7 20:12
 * @Email 1328919715@qq.com
 * @Description: 在embed中，可以将静态资源文件嵌入到三种类型的变量，分别为：字符串、字节数组、embed.FS文件类型
 **/
import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed version.txt
var version string

//go:embed bytes.txt
var bytesData []byte

//go:embed public
var embededFiles embed.FS

func main() {
	fmt.Println("Get Embed string : ", version)
	fmt.Println("Get Embed []byte : ", string(bytesData))

	useOS := len(os.Args) > 1 && os.Args[1] == "live"
	http.Handle("/", http.FileServer(getFileSystem(useOS)))
	http.ListenAndServe(":8888", nil)
}

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		log.Print("using live mode")
		return http.FS(os.DirFS("public"))
	}

	log.Print("using embed mode")

	fsys, err := fs.Sub(embededFiles, "public")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}
