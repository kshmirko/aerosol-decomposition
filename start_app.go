package main

import (
	"flag"
	"fmt"

	"gitflic.ru/project/physicist2018/aerosol-decomposition/components"
	"gitflic.ru/project/physicist2018/aerosol-decomposition/utils"
)

func main() {
	dbfile := flag.String("dbfile", "db.json", "Имя json файла с базой данных оптических свойств частиц")
	flag.Parse()

	var db components.OpticalDB
	var err error

	if db, err = components.LoadDB(*dbfile); err != nil {
		db = components.GenerateDB()
	}

	db.PrintTable()

	amix := components.AerosolModeMix{
		components.AerosolModeMixItem{
			OpticalCoefs: db[0],
			N:            1000.0,
		},
		components.AerosolModeMixItem{
			OpticalCoefs: db[2],
			N:            10.0,
		},
		components.AerosolModeMixItem{
			OpticalCoefs: db[1],
			N:            0.4,
		},
	}

	fmt.Printf("Mean radius      : %.3e\n", amix.MeanRadius())
	fmt.Printf("Effective radius : %.3e\n", amix.EffectiveRadius())
	fmt.Printf("Area             : %.3e\n", amix.Area())
	fmt.Printf("Volume           : %.3e\n", amix.Volume())

	fmt.Printf("%.3f\n", amix.RefrReIdx())
	fmt.Printf("%.4f\n", amix.RefrImIdx())

	fmt.Printf("%.2e\n", amix.F(utils.Vector{1000.0, 0.1, 100}))

	//db := components.GenerateDB()
	//components.StoreDB("test.json", &db)

}
