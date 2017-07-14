package main

import (
	"database/sql"
	"fmt"
	"github.com/swgloomy/gutil"
)

var (
	model = gutil.MySqlDBStruct{
		DbUser: "root",
		DbHost: "127.0..0.1",
		DbPort: 3309,
		DbPass: "xqsxqsxqs",
		DbName: "",
	}
	dbs *sql.DB
)

func main() {
	filePath := getXmlData()
	if len(filePath) != "" {
		loadXml(filePath)
	}
	gutil.MySqlClose(dbs)
	fmt.Println("run successfully")
}

//将xml文件录入数据库
func loadXml(filePathStr string) {
	sqlStr := fmt.Sprintf("load XML local infile '%s' into table pdc_store.m_business_order_201708 ROWS IDENTIFIED BY '<loadFile>'", filePathStr)
	_, err := gutil.MySqlSqlExec(dbs, model, sqlStr)
	if err != nil {
		fmt.Println("MySqlSqlExec ", err.Error())
		return
	}
}

// XML文件内容获取生成
func getXmlData() string {
	var (
		modelMap    map[string]string = make(map[string]string)
		fileContent []byte
		contentByte []byte
		pathStr     = "./xmlsdfasdf"
	)
	columnArrayIn, dataArrayIn, err := gutil.MysqlSelectUnknowColumn(dbs, model, "select * from pdc_store.m_business_order_201707")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	cloumnArray := *columnArrayIn

	for _, array := range *dataArrayIn {
		fileContent = append(fileContent, []byte("<loadFile>")...)
		modelMap = make(map[string]string)
		for index, item := range array {
			fileContent = append(fileContent, []byte(fmt.Sprintf("<%s>%s</%s>", cloumnArray[index], item, cloumnArray[index]))...)
			modelMap[cloumnArray[index]] = item
		}
		fileContent = append(fileContent, contentByte...)
		fileContent = append(fileContent, []byte("</loadFile>")...)
	}
	fmt.Println(len(*dataArrayIn))
	err = gutil.FileCreateAndWrite(&fileContent, pathStr, false)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return pathStr
}
