package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

const ZERO = "0"

func add(a, b string) string {
	a, b = normalize(a, b)
	var carry int = 0
	var sum int = 0
	res := ""
	for i := len(a) - 1; i >= 0; i-- {
		sum = carry + int(a[i]) + int(b[i]) - 96
		res = strconv.FormatInt(int64(sum%10), 10) + res
		carry = (sum / 10)
	}
	if carry > 0 {
		res = strconv.FormatInt(int64(carry), 10) + res
	}
	return res
}

func sub(a, b string) string {
	a, b = normalize(a, b)
	var carry int = 0
	res := ""
	for i := len(a) - 1; i >= 0; i-- {
		sum := int(a[i]) - int(b[i]) - carry
		if sum < 0 {
			sum = sum + 10
			carry = 1
		} else {
			carry = 0
		}
		res = strconv.FormatInt(int64(sum), 10) + res
	}
	return res
}

func normalize(a, b string) (string, string) {
	lenA := len(a)
	lenB := len(b)
	if lenA < lenB {
		zeros := strings.Repeat(ZERO, lenB-lenA)
		a = zeros + a
	} else if lenB < lenA {
		zeros := strings.Repeat(ZERO, lenA-lenB)
		b = zeros + b
	}
	return a, b
}

func multiply(a, b string) string {
	lenA := len(a)
	lenB := len(b)

	if lenA <= 9 && lenB <= 9 {
		iA, _ := strconv.ParseInt(a, 10, 64)
		iB, _ := strconv.ParseInt(b, 10, 64)
		return strconv.FormatInt(iA*iB, 10)
	}

	a, b = normalize(a, b)
	lenA = len(a)
	lenB = len(b)
	if lenA&1 == 1 {
		a = ZERO + a
		b = ZERO + b
		lenA = lenA + 1
		lenB = lenB + 1
	}
	middle := lenA / 2
	x1, x0 := a[:middle], a[middle:]
	y1, y0 := b[:middle], b[middle:]
	z2 := multiply(x1, y1)
	z0 := multiply(x0, y0)
	z1 := sub(multiply(add(x0, x1), add(y0, y1)), add(z0, z2))

	z2 += strings.Repeat(ZERO, middle*2)
	z1 += strings.Repeat(ZERO, middle)
	return add(add(z2, z1), z0)
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var a, b string
	var n int
	fmt.Fscan(reader, &n)
	for i := 0; i < n; i++ {
		//fmt.Fscan(reader, &a, &b)
		a = strings.Repeat("9", rand.Intn(10000))
		b = strings.Repeat("8", rand.Intn(10000))
		fmt.Fprintln(writer, strings.TrimLeft(multiply(a, b), "0"))
	}
	writer.Flush()
}
