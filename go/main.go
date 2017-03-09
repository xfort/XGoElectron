package main

import (
	"bufio"
	"os"
	"log"
	"encoding/json"
	"fmt"
)

func main() {

	startRead()
}

//开始循环读取输入数据
func startRead() {

	var readerSize = 1024 * 16
	reader := bufio.NewReaderSize(os.Stdin, readerSize)
	var lineByte []byte
	var prefixByte []byte
	var isPrefix = false
	var err error
	for {
		lineByte, isPrefix, err = reader.ReadLine()
		if err != nil {
			log.Println("ReadLine() error", err)
			continue
		}

		if isPrefix {
			if prefixByte == nil {
				prefixByte = make([]byte, 0, readerSize*8)
			}
			prefixByte = append(prefixByte, lineByte...)

		} else {
			if prefixByte == nil {
				go handleLine(lineByte)
			} else {
				go handleLine(append(prefixByte, lineByte...))
				prefixByte = nil
			}
		}
	}
}
//处理接收到的消息数据
func handleLine(lineByte []byte) {
	log.Println("handleLine()", string(lineByte))
	var eventMap map[string]interface{}
	json.Unmarshal(lineByte, &eventMap)

}

//发送结果数据
func sendMessageByte(msg []byte) {
	fmt.Println(string(msg))
}
