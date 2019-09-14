/*
* @Author: scottxiong
* @Date:   2019-07-29 02:39:58
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-09-14 15:30:28
 */
package oss

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

//whether the obj exists or not
func Exists(obj string) bool {
	// 判断文件是否存在。
	isExist, err := bucket.IsObjectExist(obj)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	//fmt.Println("Exist:", isExist)
	return isExist
}

func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x%x%x%x%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
