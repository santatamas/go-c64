package main

import (
	"fmt"
	"github.com/santatamas/go-c64/RAM"
	n "github.com/santatamas/go-c64/numeric"
	"io/ioutil"
	"log"
	"os"
)

func loadFile(path string) (RAM.Memory, byte, byte) {
	result := RAM.NewMemory()

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
	startAddress := n.ToInt16([]byte{startPCL, startPCH})

	currentAddress := startAddress
	for i := 2; i < len(byteContent); i++ {
		result.WriteAbsolute(currentAddress, byteContent[i])
		currentAddress++
	}
	//fmt.Printf("%08b", b)
	return result, startPCH, startPCL
}
