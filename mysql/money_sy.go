package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"math"
	"strconv"
)

func main() {
	f, err := excelize.OpenFile("mysql/二次判定提报-东南区.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
	rows := f.GetRows("Sheet4")
	for _, row := range rows {
		hid := row[0]
		city_name := row[1]
		city_id := row[2]
		start_date := row[3]
		end_date := row[4]
		pid := row[5]

		row6, _ := strconv.ParseFloat(row[6], 64)
		no_login := math.Round(row6*100) / 100

		row7, _ := strconv.ParseFloat(row[7], 32)
		diff_city_click := math.Round(row7*100) / 100

		fmt.Sprintf("%.3f", no_login)
		fmt.Sprintf("%f", diff_city_click)
		return

		print(hid, city_id, city_name, start_date, end_date, pid)

		//print(no_login)
		//print(diff_city_click)
		return
	}

}
