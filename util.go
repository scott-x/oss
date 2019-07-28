/*
* @Author: scottxiong
* @Date:   2019-07-29 02:39:58
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-29 02:40:42
*/
package oss

import (
	"fmt"
	"os"
)
//whether the obj exists or not
func Exists(obj string) bool{
	// 判断文件是否存在。
	isExist, err := bucket.IsObjectExist(obj)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	//fmt.Println("Exist:", isExist)
	return isExist
}