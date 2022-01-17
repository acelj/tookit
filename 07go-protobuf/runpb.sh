# proto 文件中加入了option go_package="./testpb"; 才能成功,否则报错
protoc --go_out=. ./test.proto   

go run test.go
