#!/usr/local/bin/fish

read -p 'echo "n: "' n
echo ""

set path "data/test.txt"
echo "==== make test.txt ===="
docker-compose exec app go run ./application/mktxt.go $n $path
echo "Done."
echo ""

echo "==== excute standard.go ===="
docker-compose exec app go run ./work/standard.go
echo ""

echo "==== excute sha256hex.go ===="
docker-compose exec app go run ./work/sha256hex.go
echo ""

echo "diff is ..."
diff -c cmd/data/test_sha256_hex.txt cmd/data/standard.txt