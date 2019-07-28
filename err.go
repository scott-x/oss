/*
* @Author: scottxiong
* @Date:   2019-07-29 02:33:38
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-29 02:34:16
*/

package oss

import (
	"os"
	"fmt"
)

func handleError(err error) {
    fmt.Println("Error:", err)
    os.Exit(-1)
}
