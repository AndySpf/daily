package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"unsafe"
)

func ExecuteCMD(cmd string) (string, error) {
	log.Debug(cmd)
	c := exec.Command("sh", "-c", cmd)
	bs, err := c.Output()
	if err != nil {
		log.Println(err)
		return "nil", err
	}
	return string(bs), nil
}

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func FileExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func AppendFile(path string, bs []byte) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(append([]byte{10}, bs...)); err != nil {
		return err
	}
	return nil
}
