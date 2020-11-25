package steal

import (
	"strings"

	"github.com/gocolly/colly"
)

func ListSteal(url string) []map[string]interface{}{
	k := colly.NewCollector(
		colly.AllowedDomains("wall.alphacoders.com"),
	);
	imageData := make([]map[string]interface{}, 0, 0)
	k.OnHTML("img[src]",func(e *colly.HTMLElement) {
		link := e.Attr("src")
		if !strings.Contains(link,"static.") {
			getName ,_ := e.DOM.Parent().Attr("title")
			href_link ,_ := e.DOM.Parent().Attr("href")
			var singleMap = make(map[string]interface{})
			singleMap["url"] = link
			singleMap["name"] = getName
			singleMap["href"] = strings.Replace(href_link,"big.php?i=","", -1)
			imageData = append(imageData, singleMap)
		}
	})
	k.Visit(url)
	return imageData
}
