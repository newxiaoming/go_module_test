package main

import (
	"context"
	"fmt"
	"go_modules_test/src/es_proc"
	"time"
)

var TmpArray = []es_proc.ESStudentInfo{
	{"es_stu_info", "language", &es_proc.StudentInfo{1, "Go语言学习", "Go语言", "Go是一门很好的语言","Go语言使用场景很多", "田庆阳", 100}},
	{"es_stu_info", "language", &es_proc.StudentInfo{2, "C++语言学习", "C++语言", "C++是一门很好的语言","C++语言使用场景很多", "田庆钛", 200}},
	{"es_stu_info", "language", &es_proc.StudentInfo{3, "Python语言学习", "Python语言", "Python是一门很好的语言","Python语言使用场景很多", "田海洋", 300}},
}

func main() {
	if err := es_proc.Init(); err != nil {
		fmt.Println("init fail", err)
		return
	}

	// ES是无序的，所以通过sleep控制顺序
	put()
	time.Sleep(2*time.Second)
	query()
	time.Sleep(2*time.Second)
	fmt.Println("*****************************************")
	delete()
	time.Sleep(2*time.Second)
	query()
	fmt.Println("*****************************************")

	return
}

func put()  {
	var err error
	for index, ele := range TmpArray {
		if err = es_proc.PutData(context.Background(), &ele); err != nil {
			fmt.Println(err, index)
		}
	}
}

func query() {

	queryStr := `{
		"match": {
      		"content": "语言"
    	}
	}
	`

	array, err := es_proc.GetData(context.Background(), queryStr, &es_proc.ESStudentInfo{EsIndex:"es_stu_info", EsType:"language", EsData:nil})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range array {
		fmt.Println(item)
	}
}

func delete() {
	var err error
	for index, ele := range TmpArray {
		if err = es_proc.DeleteData(context.Background(), &ele); err != nil {
			fmt.Println(err, index)
		}

		break
	}
}