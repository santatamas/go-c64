package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func loadFile(path string) (Memory, byte, byte) {
	result := newMemory()

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error! Can't open file")
	}

	byteContent, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	startPCH := byteContent[0]
	startPCL := byteContent[1]
	startAddress := toInt16([]byte{startPCL, startPCH})

	currentAddress := startAddress
	for i := 2; i < len(byteContent); i++ {
		result.WriteAbsolute(currentAddress, byteContent[i])
		currentAddress++
	}
	//fmt.Printf("%08b", b)
	return result, startPCH, startPCL
}
