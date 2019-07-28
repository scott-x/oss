/*
* @Author: scottxiong
* @Date:   2019-07-29 02:29:25
* @Last Modified by:   sottxiong
* @Last Modified time: 2019-07-29 04:34:53
*/
package oss

import (
	"os"
	"fmt"
	"path"
)
func Upload(localFile string, uuid bool){
	var newFile string
	s,_ := newUUID()
	ext := path.Ext(localFile)
    if uuid {
    	newFile= s+ext
    }else{
    	newFile=localFile
    }
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
