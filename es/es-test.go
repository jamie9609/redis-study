package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)


var EsClient *elastic.Client

var host = "http://127.0.0.1:9200/"

func  main(){
	//es 配置
	errorlog := log.New(os.Stdout, "[INFO]: ", log.LstdFlags)
	errorlog.Output(0,"jamie")
	//新建一个Logger， out设置输出，一般是文件，或者是os.Sdtout，prefix设置日记打印前缀，比如[info] [error]等信息，flag设置时间，日期等信息，可选设置如下：
	var err error

	EsClient, err := elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(host))

	if err != nil {
		panic(err)
	}

	info, code, err := EsClient.Ping(host).Do(context.Background())

	if err != nil {
		panic(err)
	}

	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := EsClient.ElasticsearchVersion(host)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
}
