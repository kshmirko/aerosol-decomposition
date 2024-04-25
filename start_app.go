package main

import (
	"fmt"

	"gitflic.ru/project/physicist2018/aerosol-decomposition/components"
)

func main() {

	am := components.AerosolMode{
		Rm:    0.07,
		Sigma: 1.95,
	}

	fmt.Printf("Mean radius      : %.3e\n", am.MeanRadius())
	fmt.Printf("Effective radius : %.3e\n", am.EffectiveRadius())
	fmt.Printf("Area             : %.3e\n", am.Area())
	fmt.Printf("Volume           : %.3e\n", am.Volume())

	db, _ := components.LoadDB("test.json")
	fmt.Println(db)
	//db := components.GenerateDB()
	//components.StoreDB("test.json", &db)

}
