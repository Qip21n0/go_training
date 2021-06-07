package main

import (
	"bufio"
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
	fmt.Printf("読み込むファイルを入力してください: ")
	fmt.Scanf("%s", &ifname)
	ifile, ierr := os.Open(ifname)
	if ierr != nil {
		_ = fmt.Errorf("%sを開けません", ifname)
	}
	defer ifile.Close()

	cpus := runtime.NumCPU()
	fmt.Printf("あなたのCPU数: %d\n", cpus)

	var n int
	fmt.Printf("goroutine数を入力してください: ")
	fmt.Scanf("%d", &n)

	var hashs = []string{}

	var wg sync.WaitGroup
	//var mu sync.Mutex

	scan := bufio.NewScanner(ifile)

	start := time.Now()
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for scan.Scan() {
				if err := scan.Err(); err != nil {
					_ = fmt.Errorf("%sを読み込めません", ifname)
					break
				}
				//mu.Lock()
				str := scan.Text()
				hashs = append(hashs, toHash(str))
				time.Sleep(time.Duration(i*10) * time.Millisecond)
				//mu.Unlock()
			}
		}(i)
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