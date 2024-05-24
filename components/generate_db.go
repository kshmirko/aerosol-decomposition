package components

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	utils "gitflic.ru/project/physicist2018/aerosol-decomposition/utlis"
)

// GenerateDB - создает массив с данными, используется если не указан файл на диске
func GenerateDB() OpticalDB {
	db := OpticalDB{
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "WASO",
				Rm:    0.0212,
				Sigma: 2.24,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{6.892719160564049e-09, 4.10101443479323e-09, 1.20229676157226e-09},
			Bck: utils.Vector{1.8669515593808272e-10, 9.776810580616591e-11, 2.7354073988952202e-11},
			B22: utils.Vector{1.8669515593808272e-10, 9.776810580616591e-11, 2.7354073988952202e-11},
			MRe: utils.Vector{1.530e+00, 1.530e+00, 1.520e+00},
			MIm: utils.Vector{5.000e-03, 6.000e-03, 1.550e-02},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "WASO",
				Rm:    2.62e-2,
				Sigma: 2.24,
			},
			Rh:  50,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{1.12e-8, 6.37e-09, 2.3e-09},
			Bck: utils.Vector{2.02e-10, 1.12e-10, 4.53e-11},
			B22: utils.Vector{2.02e-10, 1.12e-10, 4.53e-11},
			MRe: utils.Vector{1.442e+00, 1.437e+00, 1.429e+00},
			MIm: utils.Vector{2.640e-03, 3.180e-03, 8.210e-03},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "WASO",
				Rm:    2.85e-2,
				Sigma: 2.24,
			},
			Rh:  70,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{1.37e-8, 7.91e-09, 2.91e-09},
			Bck: utils.Vector{2.24e-10, 1.28e-10, 5.44e-11},
			B22: utils.Vector{2.24e-10, 1.28e-10, 5.44e-11},
			MRe: utils.Vector{1.420e+00, 1.414e+00, 1.406e+00},
			MIm: utils.Vector{2.058e-03, 2.470e-03, 6.381e-03},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "INSO",
				Rm:    0.471,
				Sigma: 2.51,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{8.21706357179527e-06, 8.47062744445488e-06, 9.11222892100765e-06},
			Bck: utils.Vector{1.4829530319167258e-07, 2.3615464999053918e-07, 3.237251781086033e-07},
			B22: utils.Vector{1.4829530319403866e-07, 2.361546499907535e-07, 3.237251781085007e-07},
			MRe: utils.Vector{1.530e+00, 1.530e+00, 1.520e+00},
			MIm: utils.Vector{8.000e-03, 8.000e-03, 8.000e-03},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "MINM",
				Rm:    0.07,
				Sigma: 1.95,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{9.80951273643021e-08, 7.321185257067201e-08, 2.84232816142845e-08},
			Bck: utils.Vector{1.1210728310687036e-09, 1.0038883582999582e-09, 5.035038138703176e-10},
			B22: utils.Vector{6.969252732364542e-10, 6.998117339069417e-10, 4.396763232133248e-10},
			MRe: utils.Vector{1.530e+00, 1.530e+00, 1.530e+00},
			MIm: utils.Vector{1.700e-02, 5.500e-03, 4.000e-03},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "MIAM",
				Rm:    0.39,
				Sigma: 2.0,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{3.2668690903309903e-06, 3.4755169389825797e-06, 3.7850200944718e-06},
			Bck: utils.Vector{2.9716028014374995e-08, 7.312997432454559e-08, 8.035918757283128e-08},
			B22: utils.Vector{1.783620295632392e-08, 4.092387257005516e-08, 4.372849145610061e-08},
			MRe: utils.Vector{1.530e+00, 1.530e+00, 1.530e+00},
			MIm: utils.Vector{1.700e-02, 5.500e-03, 4.000e-03},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "MICM",
				Rm:    1.9,
				Sigma: 2.15,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{8.05440222321512e-05, 8.17412643347408e-05, 8.49259891589777e-05},
			Bck: utils.Vector{1.5226889285703908e-07, 4.219001969465608e-07, 1.24313446620421e-06},
			B22: utils.Vector{1.4099053669634103e-07, 2.886504832852378e-07, 7.38775280012451e-07},
			MRe: utils.Vector{1.530e+00, 1.530e+00, 1.530e+00},
			MIm: utils.Vector{1.700e-02, 5.500e-03, 4.000e-03},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "SUSO",
				Rm:    0.0695,
				Sigma: 2.03,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{1.00106312645026e-07, 7.345165842958101e-08, 2.8828393431438298e-08},
			Bck: utils.Vector{2.568796457461932e-09, 1.3333850427949276e-09, 5.115661532689058e-10},
			B22: utils.Vector{2.568796457461932e-09, 1.3333850427949276e-09, 5.115661532689058e-10},
			MRe: utils.Vector{1.452e+00, 1.430e+00, 1.422e+00},
			MIm: utils.Vector{1.000e-08, 1.000e-08, 1.530e-06},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "SOOT",
				Rm:    0.0118,
				Sigma: 2.00,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{9.975630412681489e-10, 5.80274504547982e-10, 2.2764951041774198e-10},
			Bck: utils.Vector{1.0543332492524975e-11, 6.071737389433575e-12, 1.4283833931974681e-12},
			B22: utils.Vector{1.0543332492524975e-11, 6.071737389433575e-12, 1.4283833931974681e-12},
			MRe: utils.Vector{1.750e+00, 1.750e+00, 1.760e+00},
			MIm: utils.Vector{4.650e-01, 4.400e-01, 4.400e-01},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "SSAM",
				Rm:    0.209,
				Sigma: 2.03,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{9.629396726874278e-07, 1.02289491178967e-06, 9.255513590713411e-07},
			Bck: utils.Vector{8.053018450126309e-08, 6.763479500856055e-08, 2.800542853915713e-08},
			B22: utils.Vector{8.053018450126309e-08, 6.763479500856055e-08, 2.800542853915713e-08},
			MRe: utils.Vector{1.510e+00, 1.500e+00, 1.470e+00},
			MIm: utils.Vector{3.240e-07, 1.000e-08, 1.410e-04},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "MITR",
				Rm:    0.5,
				Sigma: 2.0,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{6.2253090221335104e-06, 6.50055396244829e-06, 7.1514130149097395e-06},
			Bck: utils.Vector{3.784107368851875e-08, 1.1660585796578328e-07, 1.741722708238256e-07},
			B22: utils.Vector{2.456134038720592e-08, 6.733795204552347e-08, 9.652350849459192e-08},
			MRe: utils.Vector{1.530e+00, 1.530e+00, 1.530e+00},
			MIm: utils.Vector{1.700e-02, 5.500e-03, 4.000e-03},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "MIN1",
				Rm:    0.18,
				Sigma: 1.28,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{5.39e-6, 3.95e-6, 1.53e-6},
			Bck: utils.Vector{1.34e-7, 7.14e-8, 2.83e-8},
			B22: utils.Vector{1.32e-7, 7.09e-8, 2.82e-8},
			MRe: utils.Vector{1.53, 1.53, 1.53},
			MIm: utils.Vector{1.7e-2, 5.5e-3, 4.0e-3},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "MIA1",
				Rm:    0.59,
				Sigma: 1.23,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{8.11e-6, 7.14e-6, 6.03e-6},
			Bck: utils.Vector{1.77e-7, 1.66e-7, 8.22e-8},
			B22: utils.Vector{1.58e-7, 1.18e-7, 5.65e-8},
			MRe: utils.Vector{1.53, 1.53, 1.53},
			MIm: utils.Vector{1.7e-2, 5.3e-3, 4.0e-3},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "MIC1",
				Rm:    1.04,
				Sigma: 1.45,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{1.61e-5, 1.52e-5, 1.48e-5},
			Bck: utils.Vector{1.90e-7, 3.19e-7, 4.01e-7},
			B22: utils.Vector{1.71e-7, 2.83e-7, 2.24e-7},
			MRe: utils.Vector{1.53, 1.53, 1.53},
			MIm: utils.Vector{1.7e-2, 5.3e-3, 4.0e-3},
		},
	}
	// fmt.Printf("Размер БД = %d\n\n", len(db))
	// content, err := json.MarshalIndent(db, "", "  ")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	os.Stdout.Write(content)
	// }
	return db
}

func GenerateDBAggl() OpticalDB {
	db := OpticalDB{
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "WASO",
				Rm:    0.0212,
				Sigma: 2.24,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{6.892719160564049e-09, 4.10101443479323e-09, 1.20229676157226e-09},
			Bck: utils.Vector{1.8669515593808272e-10, 9.776810580616591e-11, 2.7354073988952202e-11},
			B22: utils.Vector{1.8669515593808272e-10, 9.776810580616591e-11, 2.7354073988952202e-11},
			MRe: utils.Vector{1.530e+00, 1.530e+00, 1.520e+00},
			MIm: utils.Vector{5.000e-03, 6.000e-03, 1.550e-02},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "WASO",
				Rm:    2.62e-2,
				Sigma: 2.24,
			},
			Rh:  50,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{1.12e-8, 6.37e-09, 2.3e-09},
			Bck: utils.Vector{2.02e-10, 1.12e-10, 4.53e-11},
			B22: utils.Vector{2.02e-10, 1.12e-10, 4.53e-11},
			MRe: utils.Vector{1.442e+00, 1.437e+00, 1.429e+00},
			MIm: utils.Vector{2.640e-03, 3.180e-03, 8.210e-03},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "WASO",
				Rm:    2.85e-2,
				Sigma: 2.24,
			},
			Rh:  70,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{1.37e-8, 7.91e-09, 2.91e-09},
			Bck: utils.Vector{2.24e-10, 1.28e-10, 5.44e-11},
			B22: utils.Vector{2.24e-10, 1.28e-10, 5.44e-11},
			MRe: utils.Vector{1.420e+00, 1.414e+00, 1.406e+00},
			MIm: utils.Vector{2.058e-03, 2.470e-03, 6.381e-03},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "INSO",
				Rm:    0.471,
				Sigma: 2.51,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{8.21706357179527e-06, 8.47062744445488e-06, 9.11222892100765e-06},
			Bck: utils.Vector{1.4829530319167258e-07, 2.3615464999053918e-07, 3.237251781086033e-07},
			B22: utils.Vector{1.4829530319403866e-07, 2.361546499907535e-07, 3.237251781085007e-07},
			MRe: utils.Vector{1.530e+00, 1.530e+00, 1.520e+00},
			MIm: utils.Vector{8.000e-03, 8.000e-03, 8.000e-03},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "MINM",
				Rm:    0.07,
				Sigma: 1.95,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{3.01e-09, 1.30e-09, 2.24e-10},
			Bck: utils.Vector{4.78e-11, 3.04e-11, 9.17e-12},
			B22: utils.Vector{4.4e-11, 2.92e-11, 8.99e-12},
			MRe: utils.Vector{1.530e+00, 1.530e+00, 1.530e+00},
			MIm: utils.Vector{1.700e-02, 5.500e-03, 4.000e-03},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "MIAM",
				Rm:    0.39,
				Sigma: 2.0,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{4.55e-07, 4.16e-07, 2.59e-07},
			Bck: utils.Vector{5.83e-09, 6.85e-09, 3.80e-09},
			B22: utils.Vector{3.15e-9, 4.04e-9, 2.85e-9},
			MRe: utils.Vector{1.530e+00, 1.530e+00, 1.530e+00},
			MIm: utils.Vector{1.700e-02, 5.500e-03, 4.000e-03},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "MICM",
				Rm:    1.9,
				Sigma: 2.15,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{5.66e-06, 8.55e-06, 1.23e-05},
			Bck: utils.Vector{5.94e-08, 2.02e-07, 2.98e-07},
			B22: utils.Vector{2.44e-8, 0.93e-7, 1.28e-7},
			MRe: utils.Vector{1.530e+00, 1.530e+00, 1.530e+00},
			MIm: utils.Vector{1.700e-02, 5.500e-03, 4.000e-03},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "SUSO",
				Rm:    0.0695,
				Sigma: 2.03,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{1.00106312645026e-07, 7.345165842958101e-08, 2.8828393431438298e-08},
			Bck: utils.Vector{2.568796457461932e-09, 1.3333850427949276e-09, 5.115661532689058e-10},
			B22: utils.Vector{2.568796457461932e-09, 1.3333850427949276e-09, 5.115661532689058e-10},
			MRe: utils.Vector{1.452e+00, 1.430e+00, 1.422e+00},
			MIm: utils.Vector{1.000e-08, 1.000e-08, 1.530e-06},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "SOOT",
				Rm:    0.0118,
				Sigma: 2.00,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{9.975630412681489e-10, 5.80274504547982e-10, 2.2764951041774198e-10},
			Bck: utils.Vector{1.0543332492524975e-11, 6.071737389433575e-12, 1.4283833931974681e-12},
			B22: utils.Vector{1.0543332492524975e-11, 6.071737389433575e-12, 1.4283833931974681e-12},
			MRe: utils.Vector{1.750e+00, 1.750e+00, 1.760e+00},
			MIm: utils.Vector{4.650e-01, 4.400e-01, 4.400e-01},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "SSAM",
				Rm:    0.209,
				Sigma: 2.03,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{9.629396726874278e-07, 1.02289491178967e-06, 9.255513590713411e-07},
			Bck: utils.Vector{8.053018450126309e-08, 6.763479500856055e-08, 2.800542853915713e-08},
			B22: utils.Vector{8.053018450126309e-08, 6.763479500856055e-08, 2.800542853915713e-08},
			MRe: utils.Vector{1.510e+00, 1.500e+00, 1.470e+00},
			MIm: utils.Vector{3.240e-07, 1.000e-08, 1.410e-04},
		},
		OpticalCoefs{
			AerosolMode: AerosolMode{
				Title: "MITR",
				Rm:    0.5,
				Sigma: 2.0,
			},
			Rh:  0,
			Wvl: utils.Vector{0.355, 0.532, 1.064},
			Ext: utils.Vector{6.2253090221335104e-06, 6.50055396244829e-06, 7.1514130149097395e-06},
			Bck: utils.Vector{3.784107368851875e-08, 1.1660585796578328e-07, 1.741722708238256e-07},
			B22: utils.Vector{2.456134038720592e-08, 6.733795204552347e-08, 9.652350849459192e-08},
			MRe: utils.Vector{1.530e+00, 1.530e+00, 1.530e+00},
			MIm: utils.Vector{1.700e-02, 5.500e-03, 4.000e-03},
		},
	}
	// fmt.Printf("Размер БД = %d\n\n", len(db))
	// content, err := json.MarshalIndent(db, "", "  ")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	os.Stdout.Write(content)
	// }
	return db
}

// StoreDB - сохраняет БД в файл на диске
func StoreDB(fname string, db *OpticalDB) {
	content, err := json.MarshalIndent(*db, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(fname, content, 0666)
}

func LoadDB(fname string) (OpticalDB, error) {
	var ret OpticalDB
	content, err := os.ReadFile(fname)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = json.Unmarshal(content, &ret)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return ret, nil
}
