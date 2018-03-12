// Copyright 2018 Shestakov Denis. All rights reserved.
// Use of this source code is governed by a Lil
// license that can be found in the LICENSE file.

package readFile

import (
	"fmt"
	"os"
	"time"
)

var bufferSize int64 = 32768
var readPos int64 = 0
var timeDuration int = 2

func ReadFile(file *os.File) ([]byte, int) {
	buffer := make([]byte, bufferSize)
	n, err := file.ReadAt(buffer, readPos)
	if err != nil {
		fmt.Print(err, " ")
	}
	var subPos = bufferSize - 1
	for i := bufferSize - 1; i > 0; i-- {
		if buffer[subPos] == '\n' {
			subPos += 1
			break
		}
		buffer[subPos] = 0
		subPos--
	}
	readPos += subPos
	return buffer, n
}

func ProcessFile(file *os.File) {
	for {
		_, bytesRead := ReadFile(file)
		if int64(bytesRead) != bufferSize {
			StartTimer()
		}
	}
}

func StartTimer() {
	timer1 := time.NewTimer(time.Second * time.Duration(timeDuration))
	<-timer1.C
	_ = timer1.Stop()
}
