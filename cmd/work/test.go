package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"sync"
	"time"
)

func toHash(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

func main() {
	ifname := "data/ttest.txt"
	ifile, ierr := os.Open(ifname)
	if ierr != nil {
		_ = fmt.Errorf("%sを開けません", ifname)
	}
	defer ifile.Close()

	var n int
	fmt.Printf("スレッド数を入力してください: ")
	fmt.Scanf("%d", &n)

	scan := bufio.NewScanner(ifile)

	var wg sync.WaitGroup
	var mu sync.Mutex

	start := time.Now()
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for scan.Scan() {
				mu.Lock()
				str := scan.Text()
				if err := scan.Err(); err != nil {
					_ = fmt.Errorf("%sを読み込めません", ifname)
					break
				}
				hash := toHash(str)
				fmt.Printf("%s\n", hash)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	measure := time.Since(start)
	fmt.Printf("time: %f[s]", measure.Seconds())
}