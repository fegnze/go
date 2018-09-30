package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/Luxurioust/excelize"
)

func parseXcsl(rows [][]string) (types []string, names []string, data [][]string) {
	return rows[1], rows[2], rows[3:]
}

func parseSheetData(xlsx *excelize.File, sheetName string) map[string]interface{} {
	sheetData := xlsx.GetRows(sheetName)
	types, names, datas := parseXcsl(sheetData)

	mapData := map[string]interface{}{}
	for rowIndex, row := range datas {
		rowData := map[string]interface{}{}
		for colIndex, value := range row {
			switch types[colIndex] {
			case "int":
				intValue, err := strconv.Atoi(value)
				if err != nil {
					panic(fmt.Sprintf("TypeError@@ sheetName:%s,rowIndex:%d,colIndex:%d,value(%v) expect int,(%v)", sheetName, rowIndex, colIndex, value, err))
				}
				rowData[names[colIndex]] = intValue
			case "string":
				if reflect.TypeOf(value).String() != "string" {
					panic(fmt.Sprintf("TypeError@@ sheetName:%s,rowIndex:%d,colIndex:%d,value(%v) expect string", sheetName, rowIndex, colIndex, value))
				}
				rowData[names[colIndex]] = value
			case "array":
				rowData[names[colIndex]] = getArrayData(xlsx, names[colIndex], value)
			case "object":
				rowData[names[colIndex]] = getObjectData(xlsx, names[colIndex], value)
			default:
				panic(fmt.Sprintf("TypeNameError(%s)@@ sheetName:%s,rowIndex:%d,colIndex:%d", types[colIndex], sheetName, rowIndex, colIndex))
			}
		}
		mapData[row[0]] = rowData
	}
	return mapData
}

func getArrayData(xlsx *excelize.File, sheetName string, rowIndex string) []interface{} {
	sheetData := xlsx.GetRows(sheetName)
	types, names, datas := parseXcsl(sheetData)

	array := []interface{}{}
	for _, row := range datas {
		if row[0] == rowIndex {
			for colIndex, value := range row {
				if colIndex == 0 {
					continue
				}
				switch types[colIndex] {
				case "int":
					intValue, err := strconv.Atoi(value)
					if err != nil {
						panic(fmt.Sprintf("TypeError@@ sheetName:%s,rowIndex:%s,colIndex:%d,value(%v) expect int,(%v)", sheetName, rowIndex, colIndex, value, err))
					}
					array = append(array, intValue)
				case "string":
					if reflect.TypeOf(value).String() != "string" {
						panic(fmt.Sprintf("TypeError@@ sheetName:%s,rowIndex:%s,colIndex:%d,value(%v) expect string", sheetName, rowIndex, colIndex, value))
					}
					array = append(array, value)
				case "array":
					array = append(array, getArrayData(xlsx, names[colIndex], value))
				case "object":
					array = append(array, getObjectData(xlsx, names[colIndex], value))
				default:
					panic(fmt.Sprintf("TypeNameError(%s)@@ sheetName:%s,rowIndex:%s,colIndex:%d", types[colIndex], sheetName, rowIndex, colIndex))
				}
			}
		}
	}

	return array
}

func getObjectData(xlsx *excelize.File, sheetName string, rowIndex string) map[string]interface{} {
	sheetData := xlsx.GetRows(sheetName)
	types, names, datas := parseXcsl(sheetData)

	mapData := map[string]interface{}{}
	for _, row := range datas {
		if row[0] == rowIndex {
			for colIndex, value := range row {
				switch types[colIndex] {
				case "int":
					intValue, err := strconv.Atoi(value)
					if err != nil {
						panic(fmt.Sprintf("TypeError@@ sheetName:%s,rowIndex:%s,colIndex:%d,value(%v) expect int,(%v)", sheetName, rowIndex, colIndex, value, err))
					}
					mapData[names[colIndex]] = intValue
				case "string":
					if reflect.TypeOf(value).String() != "string" {
						panic(fmt.Sprintf("TypeError@@ sheetName:%s,rowIndex:%s,colIndex:%d,value(%v) expect string", sheetName, rowIndex, colIndex, value))
					}
					mapData[names[colIndex]] = value
				case "array":
					mapData[names[colIndex]] = getArrayData(xlsx, names[colIndex], value)
				case "object":
					mapData[names[colIndex]] = getObjectData(xlsx, names[colIndex], value)
				default:
					panic(fmt.Sprintf("TypeNameError(%s)@@ sheetName:%s,rowIndex:%s,colIndex:%d", types[colIndex], sheetName, rowIndex, colIndex))
				}

			}
		}
	}

	return mapData
}

func writeWithIo(file, content string) {
	fileObj, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to open the file,%s", err.Error()))
	}
	if _, err := io.WriteString(fileObj, content); err == nil {
		fmt.Println("Successful appending to the file with os.OpenFile and io.WriteString.", content)
	}
}

func main() {
	// args := []string{"1"}
	// args = append(args, "E:\\WorkSpace\\GoWorld\\src\\excel2json\\equip.xlsx")
	args := os.Args
	var file string
	defer func() {
		if err := recover(); err != nil {
			writeWithIo(file+".error", string(file+"::"+err.(string)))
		}
	}()

	//读取输出目录
	var output string
	if _, e := os.Stat("./ini.json"); !os.IsNotExist(e) {
		data, err := ioutil.ReadFile("./ini.json")
		if err != nil {
			panic(fmt.Sprintf("Failed to open the file,%s", err.Error()))
		}
		str := string(data)
		fmt.Println(str)
		type Configs struct {
			Output string
		}
		var cfg Configs
		if err = json.Unmarshal([]byte(str), &cfg); err != nil {
			panic(fmt.Sprintf("Failed to load ini,%s", err.Error()))
		}

		output = cfg.Output
		_, err0 := os.Stat(output)
		if os.IsNotExist(err0) {
			output = ""
		}
	}

	for _, v := range args[1:] {
		file = v
		names := strings.Split(file, "\\")
		tmp := names[len(names)-1]
		name := strings.Split(tmp, ".")[0]
		path := strings.Replace(file, tmp, "", -1)
		if output != "" {
			path = output
		}
		xlsx, err := excelize.OpenFile(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		datas := parseSheetData(xlsx, "main")
		b, err := json.MarshalIndent(datas, "", "  ")
		str := string(b)

		fmt.Println(str)
		writeWithIo(path+"\\"+name+".json", str)
	}
}
