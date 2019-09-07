package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"os"
	"strconv"
)

func GetClass(num int) string {
	xlsx, err := excelize.OpenFile("./excel/grade/2016级学生成绩.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	axis:="A"
	axis+=strconv.Itoa(num)
	cell, _ := xlsx.GetCellValue("Sheet1", axis)
	cell=cell+"1"
	return cell
}

func GetNum(num int) string{
	xlsx, err := excelize.OpenFile("./excel/grade/2016级学生成绩.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	axis:="B"
	axis+=strconv.Itoa(num)
	cell, _ := xlsx.GetCellValue("Sheet1", axis)
	return cell
}

func GetName(num int) string{
	xlsx, err := excelize.OpenFile("./excel/grade/2016级学生成绩.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	axis:="C"
	axis+=strconv.Itoa(num)
	cell, _ := xlsx.GetCellValue("Sheet1", axis)
	return cell
}
