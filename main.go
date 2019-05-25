package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	readString func() string
	readBytes  func() []byte
	stdout     *bufio.Writer
)

func init() {
	readString, readBytes = newReadString(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
}
func newReadString(ior io.Reader) (func() string, func() []byte) {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), int(1e+11))
	r.Split(bufio.ScanWords)
	f1 := func() string {
		if !r.Scan() {
			panic("panic")
		}
		return r.Text()
	}
	f2 := func() []byte {
		if !r.Scan() {
			panic("panic")
		}
		return r.Bytes()
	}
	return f1, f2
}
func readInt() int {
	return int(readInt64())
}
func readInt64() int64 {
	i, err := strconv.ParseInt(readString(), 0, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}
func readFloat64() float64 {
	f, err := strconv.ParseFloat(readString(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func readStrings(len int) []string {
	var ss []string
	for i := 0; i < len; i++ {
		s := readString()
		ss = append(ss, s)
	}
	return ss
}

func readInts(len int) []int {
	var ii []int
	for j := 0; j < len; j++ {
		i := readInt()
		ii = append(ii, i)
	}
	return ii
}

func uniqueInt(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	n := readInt()
	d := readInts(n)

	ud := uniqueInt(d)
	// sort.Ints(ud)

	fmt.Println(len(ud))
}
