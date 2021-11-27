package utils

import (
	"curry/config"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/joiggama/money"
)

func FormatCurrency(s string) string {
	reg, err := regexp.Compile(`[^\d_.]`)
	if err != nil {
		log.Fatal(err)
	}

	cur, e := strconv.ParseFloat(reg.ReplaceAllString(strings.ReplaceAll(s, ",", "."), ""), 64)
	if e != nil {
		log.Panic(e)
	}

	cfg := config.GetConfig()

	return money.Format(cur, money.Options{
		"currency":          cfg.Currency.Symbol,
		"with_symbol_space": cfg.Currency.SymbolSpace,
		"with_currency":     cfg.Currency.ShowCurrency,
	})
}
