package helpers

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func FormatCurrency(currencyString string) string {
	reg, err := regexp.Compile(`[^\d_.]`)
	if err != nil {
		log.Println(err)
	}

	cleanedString := strings.TrimSpace(strings.ReplaceAll(currencyString, ",", "."))
	currencyValue, err := strconv.ParseFloat(reg.ReplaceAllString(cleanedString, ""), 64)
	if err != nil {
		log.Println(err)
	}

	formattedCurrency := fmt.Sprintf("%.2f", currencyValue)

	return formattedCurrency
}
