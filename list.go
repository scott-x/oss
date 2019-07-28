/*
* @Author: scottxiong
* @Date:   2019-07-28 22:50:07
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-29 02:43:17
*/
package oss

import (
   "os"
   "fmt"
)

func ListAll(){
	 // 列举所有文件。
    marker := ""
    for {
        lsRes, err := bucket.ListObjects(oss.Marker(marker))
        if err != nil {
            handleError(err)
        }

        // 打印列举文件，默认情况下一次返回100条记录。 
        for _, object := range lsRes.Objects {
            fmt.Println("Bucket: ", object.Key)
        }

        if lsRes.IsTruncated {
            marker = lsRes.NextMarker
        } else {
            break
        }
    }
}

func ListMax(n int){
	 // 设置列举文件的最大个数，并列举文件。
    lsRes, err := bucket.ListObjects(oss.MaxKeys(n))
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }

    // 打印结果。
    fmt.Println("Objects:", lsRes.Objects)
    for _, object := range lsRes.Objects {
        fmt.Println("Object:", object.Key)
    }
}


func ListPrefix(prefix string){
	 // 列举包含指定前缀的文件。默认列举100个文件。
    lsRes, err := bucket.ListObjects(oss.Prefix(prefix))
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }

    // 打印结果。
    fmt.Println("Objects:", lsRes.Objects)
    for _, object := range lsRes.Objects {
        fmt.Println("Object:", object.Key)
    }
}

func RemoveOne(obj string){
	// 删除单个文件。
	err = bucket.DeleteObject(obj)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func RemoveMany(obj ...string){
	var files []string
	for _,o :=range obj{
        files=append(files,o)
	}
	// 返回删除成功的文件。
	delRes, err := bucket.DeleteObjects(files)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("Deleted Objects:", delRes.DeletedObjects)

}

//同一存储空间内拷贝文件
func CopyObj(obj,destObjectName string){
	// 拷贝文件到同一个存储空间的另一个文件。
	_, err = bucket.CopyObject(obj, destObjectName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
