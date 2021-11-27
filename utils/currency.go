package utils

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func FormatCurrency(s string) float32 {
	reg, err := regexp.Compile(`[^\d_.]`)
	if err != nil {
		log.Fatal(err)
	}

	cur, e := strconv.ParseFloat(reg.ReplaceAllString(strings.ReplaceAll(s, ",", "."), ""), 64)
	if e != nil {
		log.Panic(e)
	}

	return float32(cur)
}
