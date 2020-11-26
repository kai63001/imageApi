package steal

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type ImageDetail struct {
	image string
}

func ImageDetailSteal(typeWall string,id string) ImageDetail{
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
    extensions.Referer(c)

	var imageDetail ImageDetail
	_ = imageDetail

	c.OnHTML("body",func(e *colly.HTMLElement){
		fmt.Println("fuck u")
	})
	c.Visit("https://"+typeWall+".alphacoders.com"+"/big.php?i="+id)
	return imageDetail
}
