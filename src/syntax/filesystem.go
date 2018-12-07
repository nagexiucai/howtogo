package main

import (
	"os"
	"fmt"
	"archive/zip"
	"io"
	"path"
	"bufio"
	"strings"
	"log"
	"path/filepath"
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
		wtr, err := zipwtr.Create(path.Base(src[i]))
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
func List(zippath string) []string {
	result := []string{}
	stat, err := os.Stat(zippath)
	if err != nil {
		fmt.Println("无法获取文件状态：",zippath)
		return result
	}
	if stat.IsDir() {
		fmt.Println("是目录：",zippath)
		return result
	}
	zipfile, err := zip.OpenReader(zippath)
	if err == nil {
		defer zipfile.Close()
		for i, innerfile := range zipfile.File {
			fmt.Println(i,"-----",innerfile.Name)
			result = append(result, innerfile.Name)
			f, err := innerfile.Open()
			if err == nil {
				defer f.Close()
				buf := []byte{}
				f.Read(buf) // TODO: 没读出来？
				raw := string(buf)
				data := bufio.NewScanner(strings.NewReader(raw))
				data.Split(bufio.ScanRunes)
				for data.Scan() {
					fmt.Print(data.Text())
				}
			}else{
				fmt.Println("打开内部归档文件失败：", innerfile.Name)
			}
		}
	}else{
		fmt.Println("打开ZIP失败：",zippath)
	}
	return result
}
// 解压ZIP
func Uncompress(zippath string, dst string) bool {
	zipfile, err := zip.OpenReader(zippath)
	if err != nil {
		fmt.Println("打开ZIP失败：",zippath)
	}else{
		defer zipfile.Close()
		for _, innerfile := range zipfile.Reader.File {
			zippedfile, err := innerfile.Open()
			if err != nil {
				log.Fatalln(err)
			}
			defer zippedfile.Close()

			tgt := filepath.Join(dst,path.Base(innerfile.Name))
			if innerfile.FileInfo().IsDir() {
				os.MkdirAll(tgt, innerfile.Mode())
			}else{
				outputfile, err := os.OpenFile(tgt,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,innerfile.Mode())
				if err != nil {
					log.Fatalln(err)
				}
				defer outputfile.Close()
				_, err =io.Copy(outputfile, zippedfile)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}
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
		f.Write([]byte{'f','g','h','i','j'})
	}

	Compress([]string{".","./none.test","./x.test","./y.test","./z.test"}, "./filesystem.go.zip")

	List("./filesystem.go.zip")

	Uncompress("./filesystem.go.zip", ".")
}
