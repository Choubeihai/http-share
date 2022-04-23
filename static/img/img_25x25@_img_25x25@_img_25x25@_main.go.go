package main

import (
	"os"
	p "path"
	"path/filepath"
)

func main() {
	// 遍历文件夹，获取文件路径
	paths := make([]string, 0)
	filepath.Walk("/Users/yanghsl/GolandProjects/http-share/http/img", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	// 遍历文件路径，修改文件名
	for _, path := range paths {
		newPath := filepath.Join(filepath.Dir(path), "img_25x25@_"+p.Base(path))
		os.Rename(path, newPath)
	}
}
