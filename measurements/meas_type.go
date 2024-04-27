package measurements

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gitflic.ru/project/physicist2018/aerosol-decomposition/utlis"
)

const MAX_SIZE int = 6

type Measurement struct {
	Title string
	Data  utlis.Vector
}

func (m Measurement) Dep() float64 {
	return m.Data[MAX_SIZE-1]
}

type Measurements []Measurement

func LoadFromFile(fname string) (Measurements, error) {

	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	ret := make(Measurements, 0, 10)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), "\t")
		tmp := Measurement{Title: items[0]}
		for i := 1; i < len(items); i++ {
			tmp.Data[i-1], _ = strconv.ParseFloat(items[i], 64)
		}
		tmp.Data[MAX_SIZE-1] /= 100.0
		ret = append(ret, tmp)
	}
	return ret, nil
}
