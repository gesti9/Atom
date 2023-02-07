package exel

import (
	"atom/pkg"
	"fmt"
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func ChangeExel(kod string, bin string, name string, volume int, price int) {
	f, err := excelize.OpenFile("schet.xlsx")

	if err != nil {
		log.Fatal(err)
	}
	// kod1 := f.GetCellValue("TDSheet", "AB13")
	// fmt.Println(kod1)

	// bin1 := f.GetCellValue("TDSheet", "F22")
	// fmt.Println(bin1)

	// name1 := f.GetCellValue("TDSheet", "G26")
	// fmt.Println(name1)

	// volume1 := f.GetCellValue("TDSheet", "R26")
	// fmt.Println(volume1)

	// price1 := f.GetCellValue("TDSheet", "Y26")
	// // fmt.Println(price1)
	// num1, _ := strconv.Atoi(volume)
	// fmt.Println(volume)
	// num2, _ := strconv.Atoi(price)
	// fmt.Println(price)
	sum := volume * price
	fmt.Println(sum)

	f.SetCellValue("TDSheet", pkg.KOD, kod)
	f.SetCellValue("TDSheet", pkg.BIN, bin)
	f.SetCellValue("TDSheet", pkg.NAME, name)
	f.SetCellValue("TDSheet", pkg.VOLUME, volume)
	f.SetCellValue("TDSheet", pkg.PRICE, price)
	f.SetCellValue("TDSheet", "AE26", sum)
	f.SetCellValue("TDSheet", "AE28", sum)
	f.SetCellValue("TDSheet", "B30", "Всего наименований "+strconv.Itoa(sum)+" на сумму "+strconv.Itoa(sum)+",00 KZT") //меняю нужные поля

	if err := f.SaveAs("schet.xlsx"); err != nil { //сохранение файла
		log.Fatal(err)
	}
}
