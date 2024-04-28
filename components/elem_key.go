package components

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ElemKey struct {
	Title string
	Rh    int
}

type ElemKeys []ElemKey

func LoadKeys(fname string) (ElemKeys, error) {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	ret := make(ElemKeys, 0, 10)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), " ")
		rh, _ := strconv.ParseInt(items[1], 10, 64)
		tmp := ElemKey{
			Title: items[0],
			Rh:    int(rh),
		}

		ret = append(ret, tmp)
	}
	return ret, nil
}
