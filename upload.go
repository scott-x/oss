/*
* @Author: scottxiong
* @Date:   2019-07-29 02:29:25
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-09-14 15:08:52
 */
package oss

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func Upload(localFile string, uuid bool) string {
	var newFile string
	s, _ := newUUID()
	ext := path.Ext(strings.Trim(localFile, " "))
	if uuid {
		newFile = s + ext
	} else {
		newFile = strings.Trim(localFile, " ")
	}
	fd, err := os.Open(localFile)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
		os.Exit(-1)
	}
	defer fd.Close()

	// 上传文件流。
	err = bucket.PutObject(newFile, fd)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
		os.Exit(-1)
	}
	fmt.Printf("configuration.Endpoint:%s\n", *configuration.Endpoint)
	return ""
}

func UploadToSpecificFolder(localFile string, uuid bool, desFolder string) string {
	var newFile string
	s, _ := newUUID()
	ext := path.Ext(strings.Trim(localFile, " "))
	if uuid {
		newFile = s + ext
	} else {
		newFile = strings.Trim(localFile, " ")
	}
	fd, err := os.Open(localFile)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
		os.Exit(-1)
	}
	defer fd.Close()

	if !strings.Contains(desFolder, "/") {
		desFolder = desFolder + "/"
	}
	// 上传文件流。
	err = bucket.PutObject(desFolder+newFile, fd)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
		os.Exit(-1)
	}
	return "https://scott-x." + strings.Trim(configuration.Endpoint, "http://") + "/" + desFolder + newFile
}
