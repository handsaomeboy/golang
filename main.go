package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	paths := os.Args
	md := Md5Sum{}
	md.Read(paths)
}

func (m *Md5Sum) Read(paths []string) {
	for index, path := range paths {
		if index == 0 {
			continue
		}
		content, err := os.ReadFile(path)
		// 内存大小
		if err != nil {
			fmt.Println("read file error,file path is:" + path)
		}
		m.md5Sum(path, content)

	}
}

func (m *Md5Sum) md5Sum(name string, content []byte) {
	re := md5.Sum(content)
	fmt.Printf(" %x \t %s \n", re, name)
}

type Md5Sum struct {
}

// 实现一个管道文件  socket 、fd
// 理解管道 linux 云计算
// 虚拟化技术 容器 namespace基础上
//