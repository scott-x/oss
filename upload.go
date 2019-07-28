/*
* @Author: scottxiong
* @Date:   2019-07-29 02:29:25
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-29 02:35:09
*/
package oss

import (
	"os"
	"fmt"
)
func Upload(localFile, newFile string){
	fd, err := os.Open(localFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	defer fd.Close()

	// 上传文件流。
	err = bucket.PutObject(newFile, fd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
