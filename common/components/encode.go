package components

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5Encode(password string) string {
	w := md5.New()
	io.WriteString(w, password)
	md5str := string(fmt.Sprintf("%x", w.Sum(nil)))
	return md5str
}
