package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bitcolibri/candlePrintGo"
)

func readData(path string) ([]candlePrintGo.Candle, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scan := bufio.NewScanner(f)
	scan.Split(bufio.ScanLines)
	candles := make([]candlePrintGo.Candle, 0)
	i := 0
	for scan.Scan() {
		line := scan.Text()
		ohcl := strings.Split(line, ",")
		o, _ := strconv.ParseFloat(ohcl[1], 64)
		h, _ := strconv.ParseFloat(ohcl[2], 64)
		l, _ := strconv.ParseFloat(ohcl[3], 64)
		c, _ := strconv.ParseFloat(ohcl[4], 64)
		tick := candlePrintGo.NewCandleBar(o, h, l, c)
		candles = append(candles, tick)
		i++
	}
	return candles, nil
}

func main() {
	data, err := readData("./examples/TSLA.csv")
	if err != nil {
		panic(err)
	}
	chart := candlePrintGo.NewCandleChart(data[:200], 25)
	fmt.Println(chart.Render())
}
