package util

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

/**
 * created: 2019/5/17 14:47
 * By Will Fan
 */
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))  //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
