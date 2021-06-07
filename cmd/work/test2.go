package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"sync"
	"strings"
)

func main() {
	var ifname string
	fmt.Printf("読み込むファイルを入力してください: ")
	fmt.Scanf("%s", &ifname)
	ifile, ierr := os.Open(ifname)
	if ierr != nil {
		_ = fmt.Errorf("%sを開けません", ifname)
	}
	defer ifile.Close()

	var n int
	fmt.Printf("スレッド数を入力してください: ")
	fmt.Scanf("%d", &n)

	var str = []string{}

	var wg sync.WaitGroup
	var mu sync.Mutex

	scan := bufio.NewScanner(ifile)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			for scan.Scan() {
				mu.Lock()
				if err := scan.Err(); err != nil {
					_ = fmt.Errorf("%sを読み込めません", ifname)
					break
				}
				str = append(str, scan.Text())
				mu.Unlock()
			}
			defer wg.Done()
		}()
	}

	wg.Wait()

	fname := strings.Split(ifname, ".")
	ofname := fname[0] + "_sha256_hex." + fname[1]
	ofile, oerr := os.Create(ofname)
	if oerr != nil {
		_ = fmt.Errorf("%sを開けません", ofname)
	}

	for i, s := range str {
		hash := sha256.Sum256([]byte(s))
		fmt.Fprintf(ofile, "%x\n", hash[:])
		fmt.Printf("%d %s\n", i, s)
	}
}