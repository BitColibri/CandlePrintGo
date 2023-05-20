package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/muesli/termenv"

	"github.com/bitcolibri/candlePrintGo"
)

func readData(path string) ([]*candlePrintGo.Candle, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scan := bufio.NewScanner(f)
	scan.Split(bufio.ScanLines)
	candles := make([]*candlePrintGo.Candle, 0)
	for scan.Scan() {
		line := scan.Text()
		ohcl := strings.Split(line, ",")
		o, _ := strconv.ParseFloat(ohcl[0], 64)
		h, _ := strconv.ParseFloat(ohcl[1], 64)
		l, _ := strconv.ParseFloat(ohcl[2], 64)
		c, _ := strconv.ParseFloat(ohcl[3], 64)
		tick := candlePrintGo.NewCandle(o, h, l, c)
		candles = append(candles, tick)
	}
	return candles, nil
}

func main() {
	data, err := readData("./examples/data.csv")
	if err != nil {
		panic(err)
	}
	chart := candlePrintGo.NewCandleChart(data, 50, termenv.ColorProfile())
	fmt.Println(chart.Render())
}
