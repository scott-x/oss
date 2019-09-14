/*
* @Author: scottxiong
* @Date:   2019-07-29 02:29:25
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-09-14 15:15:04
 */
package oss

import (
	"encoding/json"
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
	f, err := os.Open("./oss-conf.json")
	if err != nil {
		fmt.Printf("Open configuration error: %s\n", err)
		return ""
	}
	defer f.Close()
	configuration := &conf{}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(configuration)
	if err != nil {
		panic(err)
	}
	//fmt.Println(configuration)

	endpoint := configuration.Endpoint
	bucketName := configuration.BucketName
	return "https://" + *bucketName + "." + strings.Trim(*endpoint, "http://") + "/" + desFolder + newFile
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
