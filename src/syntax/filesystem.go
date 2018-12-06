package main

import (
	"os"
	"fmt"
	"archive/zip"
	"io"
	)

// 压缩ZIP
func Compress(src []string, zippath string) bool {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	zipfile, err := os.OpenFile(zippath, flags, 0644)
	defer zipfile.Close()
	if err != nil {
		fmt.Println("不能打开ZIP文件：", zippath)
		return false
	}
	zipwtr := zip.NewWriter(zipfile)
	defer zipwtr.Close()
	for i:=0; i<len(src); i++ {
		stat, err := os.Stat(src[i])
		if err != nil {
			fmt.Println("无法获取文件状态：", src[i])
			continue
		}
		if stat.IsDir() {
			fmt.Println("是目录：", src[i])
			continue
		}
		f, err := os.Open(src[i])
		if err != nil {
			fmt.Println("没能打开文件：", src[i])
			continue
		}
		defer f.Close()
		wtr, err := zipwtr.Create(src[i])
		if err != nil {
			fmt.Println("注册ZIP失败：", zippath, src[i])
		}else {
			_, err := io.Copy(wtr, f)
			if err != nil {
				fmt.Println("写入ZIP失败：", zippath, src[i])
			}
		}
	}
	return true
}
// 解读ZIP
func List(zip string) []string {
	return []string{}
}
// 解压ZIP
func Uncompress(zip string, dst string) bool {
	return true
}

func main() {
	fmt.Println("当前工作目录：", os.Args[0])
	args := os.Args[1:]
	for i:=0; i<len(args); i++ {
		fmt.Println("第",i,"个参数是：",args[i])
	}

	f, err := os.Create("./x.test")
	if err != nil {
		fmt.Println("无法创建文件：","./x.test")
	}else{
		defer f.Close()
		f.Write([]byte("abcde"))
	}

	Compress([]string{".","./none.test","./x.test","./y.test","./z.test"}, "./filesystem.go.zip")
}
