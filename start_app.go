package main

import (
	"flag"
	"fmt"

	"gitflic.ru/project/physicist2018/aerosol-decomposition/components"
	"gitflic.ru/project/physicist2018/aerosol-decomposition/solver"
	"gitflic.ru/project/physicist2018/aerosol-decomposition/utlis"
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

	y := utlis.Vector{3.66e-06, 1.69e-06, 3.26e-07, 8.31e-05, 4.68e-05, 0.0543}
	ret := solver.FindSolutionDE(&db, y)

	fmt.Printf("Resulting error: %.2f%%\n", ret.Err*100.0)
	fmt.Printf("X = %.2f\n", ret.Xsol)
	ret.Mix.PrintComponents()

	fmt.Printf("Rmean = %.3f, Reff = %.3f, Mre = %7.2f, Mim=%8.3f\n",
		ret.Mix.MeanRadius(),
		ret.Mix.EffectiveRadius(),
		ret.Mix.RefrReIdx()[1],
		ret.Mix.RefrImIdx()[1])

}
