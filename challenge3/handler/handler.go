package handler

import (
	"net/http"

	"countmeat/fetch"
	"countmeat/wordcount"

	"github.com/labstack/echo/v4"
)

// beefSummary รองรับ handles endpoint /beef/summary
func BeefSummary(c echo.Context) error {
	text, err := fetch.FetchMeatText()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	counts := wordcount.CountMeats(text)

	return c.JSON(http.StatusOK, map[string]interface{}{"beef": counts})
}
