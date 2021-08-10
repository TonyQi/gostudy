package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFileByIoUtil(path string) {
	ret, error := ioutil.ReadFile(path)
	if error != nil {
		fmt.Println("read file failed %s \n", error)
		return
	}
	fmt.Println(string(ret))

}

//打开文件
func main() {

	fileObj, error := os.Open("D:\\data\\logs\\oauth2-server\\spring.log")
	if error != nil {
		fmt.Println("")
	}
	defer fileObj.Close()
	length := 1024
	var temp []byte
	for {
		temp = make([]byte, length)
		n, err := fileObj.Read(temp)
		if err != nil {
			fmt.Printf("read from file failed,error is %s!!!\n", err)
			return
		}
		if n < length {
			fmt.Println("读取结束")
			break
		}
		//fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(temp[:n]))
	}
	readFileByIoUtil("D:\\data\\logs\\oauth2-server\\spring.log")
}
