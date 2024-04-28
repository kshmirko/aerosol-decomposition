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

// Measurement - структура хранит запись измерений
type Measurement struct {
	Title string       // Название или метка высоты или дня
	Data  utlis.Vector // измерения в формате b355, b532, b1064, a355, a532, delta
}

func NewMeasurement(title string) Measurement {
	return Measurement{
		Title: title,
		Data:  make(utlis.Vector, MAX_SIZE),
	}
}

// Dep - возвращает значение деполяризации
func (m Measurement) Dep() float64 {
	return m.Data[MAX_SIZE-1]
}

func (m Measurement) B11() float64 {
	return m.Data[1]
}

func (m Measurement) B22() float64 {
	return (1 - m.Data[MAX_SIZE-1]) * m.Data[1] / (1 + m.Data[MAX_SIZE-1])
}

func (m Measurement) Print() {
	fmt.Printf("%10s\t%.2e\n", m.Title, m.Data)
}

func (m Measurement) Print1() {
	fmt.Printf("%10s\t%.2e\n", m.Title, m.Data)
	fmt.Printf("-----------------------------------------------------------------------\n\n")
}

// Measurements - массив измерений
type Measurements []Measurement

func (m Measurements) Len() int {
	return len(m)
}

// LoadFromFile - загружает из файла данные, формируя массив
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
		tmp := NewMeasurement(items[0])
		for i := 1; i < len(items); i++ {
			tmp.Data[i-1], _ = strconv.ParseFloat(items[i], 64)
		}
		tmp.Data[MAX_SIZE-1] /= 100.0
		ret = append(ret, tmp)
	}
	return ret, nil
}

func (m Measurements) MakeAverage() Measurement {

	total := make(utlis.Vector, MAX_SIZE)
	for _, mi := range m {
		for j := 0; j < MAX_SIZE-1; j++ {
			total[j] += mi.Data[j]
		}
		total[MAX_SIZE-1] = total[MAX_SIZE-1] + mi.B22()
	}
	for i := range total {
		total[i] /= float64(len(m))
	}
	total[MAX_SIZE-1] = (total[1] - total[MAX_SIZE-1]) / (total[1] + total[MAX_SIZE-1])
	return Measurement{
		Title: "Average",
		Data:  total,
	}
}

func (m Measurements) Print() {
	fmt.Printf("  Title 	[ b355     b532     b1064    e355     e532     d532   ]\n")
	fmt.Printf("-----------------------------------------------------------------------\n")
	for _, mi := range m {
		mi.Print()
	}
	fmt.Printf("-----------------------------------------------------------------------\n\n")

}
