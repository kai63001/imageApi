package steal

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type List struct {
	NowPage string
	AllPage string
	Image []map[string]interface{}
}

func ListSteal(typeWall string,path string,page string,orderby string) List{
	k := colly.NewCollector(
		colly.AllowedDomains(typeWall+".alphacoders.com"),
	);
	extensions.RandomUserAgent(k)
    extensions.Referer(k)
	uri := "https://"+typeWall+".alphacoders.com"+"/"+path+"&page="+page
	_ = uri
	switch orderby{
		case "none" :
			uri = "https://"+typeWall+".alphacoders.com"+"/"+path+"?page="+page
		case "new" :
			uri = "https://"+typeWall+".alphacoders.com"+"/"+path+"&page="+page
	}
	if(orderby == "new"){
		err := k.Post(uri, map[string]string{"view":"paged","min_resolution":"0x0","resolution_equals":"%3E%3D","sort":"newest"})
		if err != nil {
			log.Fatal(err)
		}
	}
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
	var i = 0
	myPage := ""
	k.OnHTML(".form-control",func(e *colly.HTMLElement) {
		link := e.Attr("placeholder")
		if !strings.Contains(link,"Page") {
			return
		}
		if i > 0 {
			return
		}
		link = strings.ReplaceAll(link,"Page # / ","")
		myPage = link
		i++
	})

	k.OnError(func(_ *colly.Response, err error) {
		myPage = "error"
	})
	k.Visit(uri)
	m := List{
		page,myPage,imageData,
	}
	return m
}
