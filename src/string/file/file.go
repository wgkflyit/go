package file

import (
	"fmt"
	"os"
)

func TestFile()  {

	file,error := os.Open("E:\\\\go_workspace\\a.txt")

	if error!=nil{
		panic("打开文件出错")
	}

	var bs = make([]byte,1024,2048)
	file.Read(bs)

	file.WriteString("helloworld")

	fmt.Println(bs)

	defer file.Close()
}
