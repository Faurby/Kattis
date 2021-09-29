package main

import "os"

const BUFFER_SIZE = 1 << 16

var buffer = make([]byte, BUFFER_SIZE)
var bufferPointer = 0
var bytesRead = 0

func ReadLine() string {
	var buf []byte
	c := Read()
	for c != byte(255) {
		if c == '\n' {
			break
		}
		buf = append(buf, c)
		c = Read()
	}
	return string(buf)
}
func Skip(x int) {
	if x < 0 && 0 < bufferPointer-x {
		bufferPointer += x
	} else if x > 0 {
		if bufferPointer+x < bytesRead {
			bufferPointer += x
		} else {
			for i := 0; i < x; i++ {
				Read()
			}
		}
	}
}
func Next() string {
	var buf []byte
	c := Read()
	for c != byte(255) {
		if c == ' ' || c == '\n' {
			break
		}
		buf = append(buf, c)
		c = Read()
	}
	return string(buf)
}

func NextInt() int {
	ret := 0
	c := Read()
	for c <= ' ' {
		c = Read()
	}
	neg := c == '-'
	for c >= '0' && c <= '9' {
		ret = ret*10 + int(c) - '0'
		c = Read()
	}
	if neg {
		return -ret
	}
	return ret
}

func FillBuffer() {
	bytesRead, _ = os.Stdin.Read(buffer)
	if bytesRead == 0 {
		buffer[0] = 255
	}
	bufferPointer = 0
}

func Read() byte {
	if bytesRead == bufferPointer {
		FillBuffer()
	}
	bufferPointer++
	return buffer[bufferPointer-1]
}
