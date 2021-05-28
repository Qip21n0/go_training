#!/usr/local/bin/fish

echo "Enter the Path of code you wanna run."
read -p 'echo "path: "' path
echo ""

docker-compose exec app go run ./$path