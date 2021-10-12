package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	d, err := os.ReadFile("./dataEnc.out")
	check(err)
	k, err := os.ReadFile("./key.out")
	check(err)
	data := string(d)
	keys := strings.Split(string(k), " ")
	fmt.Println(reflect.TypeOf(keys[0]))
	i := 0
	var out string = ""
	for _, c := range data {
		if !unicode.IsSpace(c) {
			temp, err := strconv.Atoi(keys[i])
			check(err)
			out += string(rune(int(c) - temp))
			i++
		} else {
			out += string(c)
		}

	}

	fout, err := os.Create("data.out")
	check(err)
	dataDec := []byte(out)
	fout.Write(dataDec)
}
