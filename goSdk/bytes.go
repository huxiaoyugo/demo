package goSdk

import (
	"bytes"
	"fmt"
	"github.com/lunny/log"
	"io"
	"os"
	"unsafe"
)

func BytesReader() {

	str := "helloworld你好"
	reader := bytes.NewReader([]byte(str))

	for {
		ch, _, er := reader.ReadRune()
		if er == io.EOF {
			log.Info(er)
			break
		}
		if er != nil {
			log.Error(er)
			break
		}
		fmt.Printf("%c\n", ch)
	}

}


func BytesRead() {
	str := "helloworld你好"
	reader := bytes.NewReader([]byte(str))

	fmt.Println(len(str))
	b := make([]byte,16)
	reader.Read(b)

	fmt.Println(*(*string)(unsafe.Pointer(&b)))
}

func BytesWriteTo() {
	str := "helloworld你好\n"
	reader := bytes.NewReader([]byte(str))

	reader.WriteTo(os.Stdout)
}


func StringRange() {

	str := "hello world 你好"

	for i, c := range str {
		fmt.Printf("%d %c\n", i, c)
	}

	fmt.Printf("%s", str[12:15])
}