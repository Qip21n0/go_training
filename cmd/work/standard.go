package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
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

	var hashs = []string{}

	scan := bufio.NewScanner(ifile)
	start := time.Now()
	for scan.Scan() {
		if err := scan.Err(); err != nil {
			_ = fmt.Errorf("%sを読み込めません", ifname)
			break
		}
		str := scan.Text()
		hashs = append(hashs, toHash(str))
	}

	ofname := "data/standard.txt"
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