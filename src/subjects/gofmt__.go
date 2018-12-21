package main

import "fmt"

// 源代码文件名以“_”开头在构建时被忽略。

var(
	huan int = 9527
	qiuxiang int = 1314
)

func main(){
fmt.Println("huaan is", huan)
fmt.Println("qiuxiang is", qiuxiang)
}

/*

// -d 调用本地diff命令显示格式化前后的区别
// -w output-file 格式化并写入指定文件

D:\Important!\happy\howtogo>gofmt src\subjects\gofmt__.go
package main

import "fmt"

// 源代码文件名以“_”开头在构建时被忽略。

var (
        huan     int = 9527
        qiuxiang int = 1314
)

func main() {
        fmt.Println("huaan is", huan)
        fmt.Println("qiuxiang is", qiuxiang)
}

//

 */