package steal

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func ListSteal(typeWall string,path string,page int) []map[string]interface{}{
	k := colly.NewCollector(
		colly.AllowedDomains(typeWall+".alphacoders.com"),
	);
	extensions.RandomUserAgent(k)
    extensions.Referer(k)
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
	k.Visit("https://"+typeWall+".alphacoders.com"+"/"+path+"?page="+strconv.Itoa(page))
	return imageData
}
