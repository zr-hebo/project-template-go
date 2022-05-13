package main

import (
	"fmt"
	"os"
	"strings"
	"unsafe"

	log "github.com/golang/glog"
	"github.com/mohae/deepcopy"
	"github.com/zr-hebo/utils/ezflag"
)

var (
	reservedKeys = []string{"alsologtostderr", "logtostderr", "help"}
)

func showCommand() {
	fmt.Println("++++++++++++++++++++++++")
	tmpRawArgs := deepcopy.Copy(os.Args)
	rawArgs := tmpRawArgs.([]string)
	err := ezflag.Parse(reservedKeys)
	if err != nil {
		log.Warningf("parse command line error occur:%s", err.Error())
	}
	fmt.Printf("start service with command: %s\n", steganography(rawArgs))
}

func steganography(rawArgs []string) string {
	os.Args = rawArgs
	fields := make([]string, 0, len(os.Args))
	for idx, field := range os.Args {
		if strings.Contains(field, "password") || strings.Contains(field, "passwd") {
			pwdStr := (*[2]uintptr)(unsafe.Pointer(&os.Args[idx]))
			pwdStartPos := strings.IndexByte(field, '=')
			if pwdStartPos >= 0 {
				for i := pwdStartPos + 1; i < len(os.Args[idx]); i++ {
					*(*uint8)(unsafe.Pointer((*pwdStr)[0] + uintptr(i))) = '*'
				}
			}
			field = os.Args[idx]
		}

		fields = append(fields, field)
	}

	return strings.Join(fields, " ")
}
