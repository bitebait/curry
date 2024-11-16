package helpers

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var nonDigitRegex = regexp.MustCompile(`[^\d_.]`)

func FormatCurrency(currencyString string) string {
	cleanedCurrency := cleanCurrencyString(currencyString)
	currencyValue, err := strconv.ParseFloat(cleanedCurrency, 64)
	if err != nil {
		log.Println(err)
		return ""
	}

	currencyFormatted := fmt.Sprintf("%.2f", currencyValue)
	return currencyFormatted
}

func cleanCurrencyString(currencyString string) string {
	cleaned := strings.TrimSpace(strings.ReplaceAll(currencyString, ",", "."))
	return nonDigitRegex.ReplaceAllString(cleaned, "")
}
