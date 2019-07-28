/*
* @Author: scottxiong
* @Date:   2019-07-29 02:29:40
* @Last Modified by:   sottxiong
* @Last Modified time: 2019-07-29 04:41:05
*/
package oss

import (
	"fmt"
    "os"
    "io/ioutil"
    "path"
    "strings"
    "io"
    "bytes"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func StreamDownload(obj string){
	// 下载文件到流。
	body, err := bucket.GetObject(obj)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作。
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("data:", string(data))
}

func CashDownload(obj string){
	// 下载文件到缓存。
	body, err := bucket.GetObject(obj)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	defer body.Close()

	buf := new(bytes.Buffer)
	io.Copy(buf, body)
	fmt.Println("buf:", buf)
}

//very useful api
func LocalDownload(obj,localFile string){
	// 下载文件到本地文件流。
	body, err := bucket.GetObject(obj)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	defer body.Close()

	fd, err := os.OpenFile(localFile, os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1) 
	}
	defer fd.Close()

	io.Copy(fd, body)
}

func DownloadToLocal(obj,localFile string){
	// 下载文件到本地文件。
	err = bucket.GetObjectToFile(obj, localFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

// 分片下载。3个协程并发下载分片，开启断点续传下载。
// 其中"<yourObjectName>"为objectKey， "LocalFile"为filePath，100*1024为partSize。
func BreakPointDownload(obj,localFile string, routines int){
    err = bucket.DownloadFile(obj, localFile, 100*1024, oss.Routines(routines), oss.Checkpoint(true, ""))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

//zipMethod: zip gzip
func ZipFileDownload(obj string){
	// 文件压缩下载。
	zipMethod := strings.Split(path.Ext(obj),".")[1]
	localFile :=obj
	err = bucket.GetObjectToFile(obj, localFile, oss.AcceptEncoding(zipMethod))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
// 定义进度条监听器。
type OssProgressListener struct {
}
// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		fmt.Printf("Transfer Started, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferDataEvent:
		fmt.Printf("\rTransfer Data, ConsumedBytes: %d, TotalBytes %d, %d%%.",
			event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
	case oss.TransferCompletedEvent:
		fmt.Printf("\nTransfer Completed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferFailedEvent:
		fmt.Printf("\nTransfer Failed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	default:
	}
}

//with percentage download
func WithProgressDownload(obj,localFile string){
	// 带进度条的下载。
	err = bucket.GetObjectToFile(obj, localFile, oss.Progress(&OssProgressListener{}))
	if err != nil {
	fmt.Println("Error:", err)
	os.Exit(-1)
	}
	fmt.Println("Transfer Completed.")
}