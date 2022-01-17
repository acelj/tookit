package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"07go-protobuf/testpb"

	"github.com/golang/protobuf/proto"
)

func write() {
	p1 := &testpb.Person{
		Id:   1,
		Name: "小张",
		Phones: []*testpb.Phone{
			{Type: testpb.PhoneType_HOME, Number: "111111111"},
			{Type: testpb.PhoneType_WORK, Number: "222222222"},
		},
		Meta: map[string]string{"name": "校长", "age": "20"},
	}
	p2 := &testpb.Person{
		Id:   2,
		Name: "小王",
		Phones: []*testpb.Phone{
			{Type: testpb.PhoneType_HOME, Number: "333333333"},
			{Type: testpb.PhoneType_WORK, Number: "444444444"},
		},
		Meta: map[string]string{"name": "小王111", "age": "20"},
	}

	//创建地址簿
	book := &testpb.ContactBook{}
	book.Persons = append(book.Persons, p1)
	book.Persons = append(book.Persons, p2)

	//编码数据
	data, _ := proto.Marshal(book)
	//把数据写入文件
	ioutil.WriteFile("./test.txt", data, os.ModePerm)
}

func read() {
	//读取文件数据
	data, _ := ioutil.ReadFile("./test.txt")
	book := &testpb.ContactBook{}
	//解码数据
	proto.Unmarshal(data, book)
	for _, v := range book.Persons {
		fmt.Println(v.Id, v.Name)
		for _, vv := range v.Phones {
			fmt.Println(vv.Type, vv.Number)
		}
		for k, v := range v.Meta {
			fmt.Println(k, v) // 打印map的信息
		}
	}
}

func main() {
	fmt.Println("This is main...")

	msg_test := &testpb.Per{
		Name: *proto.String("Davie"),
		Age:  *proto.Int32(19),
		From: *proto.String("China"),
	}
	// 序列化
	msgDataEncoding, err := proto.Marshal(msg_test)
	if err != nil {
		panic(err.Error())
		return
	}

	// fmt.Printf("%T\n", msgDataEncoding)
	fmt.Println("msgDataEncoding type is ", reflect.TypeOf(msgDataEncoding))
	fmt.Printf("二进制打印:[%v]\n", msgDataEncoding)
	fmt.Printf("长度:[%v]\n", len(msgDataEncoding))
	fmt.Printf("字符串打印:[%s]\n", msgDataEncoding)

	fmt.Println("\n")

	//反序列化
	msgEntity := testpb.Per{}
	err = proto.Unmarshal(msgDataEncoding, &msgEntity)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}

	fmt.Println("msgEntity type is ", reflect.TypeOf(msgEntity))
	fmt.Printf("姓名：%s\n", msgEntity.GetName())
	fmt.Printf("年龄：%d\n", msgEntity.GetAge())
	fmt.Printf("国籍：%s\n", msgEntity.GetFrom())

	fmt.Println("==========================================")

	write()
	read()
}
