package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

const (
	KNOWN_GOOD_CHECKSUM = 4010533300 // fill in with known checksum from 30megabytes.bin
)

// doSomething1/test1 uses a slice to iterate over the bytes.

func doSomething1(block []byte) int {
	checksum := 0
	for _, b := range block {
		checksum += int(b)
	}
	return checksum
}

func test1() time.Duration {
	data, err := os.ReadFile("30megabytes.bin")
	if err != nil {
		panic(err)
	}

	checksum := 0
	tStart := time.Now()
	for len(data) > 0 {
		checksum += doSomething1(data[:512])
		data = data[512:]
	}

	if checksum != KNOWN_GOOD_CHECKSUM {
		panic("Checksum mismatch")
	}

	return time.Now().Sub(tStart)
}

// doSomething2/test2 uses a byte array and a range to iterate over the bytes.

func doSomething2(block []byte, start int, end int) int {
	checksum := 0
	for i := start; i < end; i++ {
		checksum += int(block[i])
	}
	return checksum
}

func test2() time.Duration {
	data, err := os.ReadFile("30megabytes.bin")
	if err != nil {
		panic(err)
	}

	checksum := 0
	tStart := time.Now()
	dataLen := len(data)
	for start := 0; start < dataLen; start += 512 {
		end := min(start+512, dataLen)
		checksum += doSomething2(data, start, end)
	}

	if checksum != KNOWN_GOOD_CHECKSUM {
		panic("Checksum mismatch")
	}

	return time.Now().Sub(tStart)
}

// doSomething3/test3 uses a pointer to the byte array.

func doSomething3(block *[]byte, start int, end int) int {
	checksum := 0
	for i := start; i < end; i++ {
		checksum += int((*block)[i])
	}
	return checksum
}

func test3() time.Duration {
	data, err := os.ReadFile("30megabytes.bin")
	if err != nil {
		panic(err)
	}

	checksum := 0
	tStart := time.Now()
	dataLen := len(data)
	for start := 0; start < dataLen; start += 512 {
		end := min(start+512, dataLen)
		checksum += doSomething3(&data, start, end)
	}

	if checksum != KNOWN_GOOD_CHECKSUM {
		panic("Checksum mismatch")
	}

	return time.Now().Sub(tStart)
}

// doSomething4/test4 is what would happen if it actually copied the entire slice every time.
// This is slow. But just for kicks, here is what it would take.

func doSomething4(block []byte) int {
	checksum := 0
	for _, b := range block {
		checksum += int(b)
	}
	return checksum
}

func test4() time.Duration {
	data, err := os.ReadFile("30megabytes.bin")
	if err != nil {
		panic(err)
	}

	checksum := 0
	tStart := time.Now()
	for len(data) > 0 {
		checksum += doSomething1(bytes.Clone(data[:512]))
		data = bytes.Clone(data[512:])
	}

	if checksum != KNOWN_GOOD_CHECKSUM {
		panic("Checksum mismatch")
	}

	return time.Now().Sub(tStart)
}

type testfunc func() time.Duration

func bench(name string, tf testfunc, iterations int) {
	var min, max time.Duration
	var total time.Duration

	for i := 0; i < iterations; i++ {
		duration := tf()
		total += duration
		if i == 0 || duration < min {
			min = duration
		}
		if duration > max {
			max = duration
		}
	}

	avg := total / time.Duration(iterations)
	fmt.Printf("%s,%v,%v,%v\n", name, min, max, avg)
}

func main() {
	// suitable for https://ozh.github.io/ascii-tables/
	fmt.Printf("Approach,Min,Max,Avg\n")
	bench("Approach 1 (Slicing)", test1, 1000)
	bench("Approach 2 (Array,Range)", test2, 1000)
	bench("Approach 3 (Array by Reference, Range)", test3, 1000)
	bench("Approach 3 (Slicing with forced copies)", test4, 2) // this will take a while...
	os.Exit(0)
}
