package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	// ... lots of work...
	time.Sleep(1 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

// func double(x int) int {
// 	return x + x
// }
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

func readFiles(filenames []string) error {
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close() // NOTE: risky; could run out of file descriptors
	}
	return nil
}

func readFiles2(filenames []string) error {
	for _, filename := range filenames {
		if err := doFile(filename); err != nil {
			return err
		}
	}
	return nil
}
func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// ...
	return nil
}

func main() {
	bigSlowOperation()

	_ = double(4)
}
