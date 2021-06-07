package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"sync"
	"strings"
	"time"
)

func toHash(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

func main() {
	var ifname string
	fmt.Printf("読み込むファイル名を入力してください: ")
	fmt.Scanf("%s", &ifname)
	ifile, ierr := os.Open(ifname)
	if ierr != nil {
		_ = fmt.Errorf("%sを開けません", ifname)
	}
	defer ifile.Close()

	var str = []string{}
	var s string

	for {
		n, err := fmt.Fscanf(ifile, "%s", &s)
		if n == 0 || err != nil {
			break
		} else {
			str = append(str, s)
		}
	}

	cpus := runtime.NumCPU()
	fmt.Printf("あなたのCPU数: %d\n", cpus)

	var n int
	fmt.Printf("goroutine数を入力してください: ")
	fmt.Scanf("%d", &n)

	var hashs = make([]string, len(str))

	var wg sync.WaitGroup
	var mu sync.Mutex

	start := time.Now()
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i, s := range str {
				mu.Lock()
				hash := toHash(s)
				hashs[i] = hash
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	fname := strings.Split(ifname, ".")
	ofname := fname[0] + "_sha256_hex." + fname[1]
	ofile, oerr := os.Create(ofname)
	if oerr != nil {
		_ = fmt.Errorf("%sを開けません", ofname)
	}
	defer ofile.Close()

	for _, hash := range hashs {
		fmt.Fprintf(ofile, "%s\n", hash)
	}

	measure := time.Since(start)
	fmt.Printf("time: %f[s]\n", measure.Seconds())
}