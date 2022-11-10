package utils

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func FormatCurrency(s string) string {
	reg, err := regexp.Compile(`[^\d_.]`)
	if err != nil {
		log.Fatal(err)
	}

	cur, err := strconv.ParseFloat(reg.ReplaceAllString(strings.ReplaceAll(s, ",", "."), ""), 64)
	if err != nil {
		log.Panic(err)
	}

	return fmt.Sprintf("%.2f", cur)
}
