//@Author: wulinlin
//@Description:
//@File:  encrypt
//@Version: 1.0.0
//@Date: 2023/12/03 17:55

package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

/** 加密方式 **/

func Md5ByString(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		panic(any(err))
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

func Md5ByBytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}
