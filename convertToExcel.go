package main

import (
	"bufio"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io"
	"os"
	"strings"
)

func main() {
	var fileName string
	fmt.Println("Please input the file name")
	if _, err := fmt.Scan(&fileName); err != nil {
		fmt.Println(err)
	}

	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Filed to open file")
	}
	defer file.Close()
	f := excelize.NewFile()
	index := f.NewSheet("query_sets")
	reader := bufio.NewReader(file)
	for row := 1; ; row++ {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		cols := strings.Split(str, "|||")
		for col := 1; col < len(cols)+1; col++ {
			val, err := excelize.CoordinatesToCellName(col, row)
			if err != nil {
				fmt.Println(err)
			}
			if err := f.SetCellValue("query_sets", val, cols[col-1]); err != nil {
				fmt.Println(err)
			}
		}
	}
	f.DeleteSheet("Sheet1")
	f.SetActiveSheet(index)
	if Exists("./terms.xlsx") {
		if err := os.Remove("./terms.xlsx"); err != nil {
			fmt.Println("Delete terms.xlsx failed")
		}
	}
	if err := f.SaveAs("./terms.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
