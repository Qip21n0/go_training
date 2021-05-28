# The Manual for this directory

## Overview
This directory is for Go lang training.

I will tarin **grammar** and **construction** of Go lang in this directory.

Actually, I'm going to store anything about Go lang.

## Commands
```fish
# confirm whether there is a file named "docker-compose.yml" in this directory.

# launch a container with docker-compose
docker-compose up -d

# execute any Go program file
# this is an example code in ./basic/helloworld.go
docker-compose exec app go run ./basic/helloworld.go
# -> Hello World

# stop the container
docker-compose down
```

**or**

```fish
./run.fish
# -> Enter the Path of code you wanna run.
# -> path: basic/helloworld.go

# -> Hello World
```

### Work hard, play hard!!!!!