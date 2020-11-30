package steal

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type Detail struct {
    Name string
	Resolution []map[string]interface{}
	Author []map[string]interface{}
    Image []map[string]interface{}
	Categoly []map[string]interface{}
	Tags []map[string]interface{}
}

func ImageDetailSteal(typeWall string,id string) Detail{
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
    extensions.Referer(c)
	imageData := make([]map[string]interface{},0,0)
	c.OnHTML("source",func(e *colly.HTMLElement){
		image := e.Attr("srcset")
		if !strings.Contains(image,"big") {
			var singleMap = make(map[string]interface{})
			singleMap["url"] = image
			imageData = append(imageData, singleMap)
			fmt.Println(image)
		}
	})
	categolyData := make([]map[string]interface{},0,0)
	c.OnHTML("strong",func(e *colly.HTMLElement){
		categoly := e.Text
		id, _ := e.DOM.Parent().Attr("href")
		id = strings.Replace(id,"by_category.php?id=","",-1)
		id = strings.Replace(id,"by_sub_category.php?id=","",-1)
		cutId := strings.Split(id,"&")
		id = cutId[0]
		var singleMap = make(map[string]interface{})
		singleMap["id"] = id;
		singleMap["name"] = categoly;
		categolyData = append(categolyData, singleMap)
	})
	tagData := make([]map[string]interface{},0,0)
	c.OnHTML(".tag-element",func(e *colly.HTMLElement){
		name := e.Text
		name = strings.Trim(name," ")
		name = strings.ReplaceAll(name,"\n","")
		name = strings.Trim(name," ")
		id := e.ChildAttr("a[href]","href")
		id = strings.Replace(id,"tags.php?tid=","",-1)
		var singleMap = make(map[string]interface{})
		singleMap["id"] = id
		singleMap["name"] = name;
		tagData = append(tagData, singleMap)
	})
	authorData := make([]map[string]interface{},0,0)
	c.OnHTML(".author-container",func(e *colly.HTMLElement){
		name := e.Text
		name = strings.Trim(name," ")
		name = strings.ReplaceAll(name,"\n","")
		name = strings.Trim(name," ")
		name = strings.ReplaceAll(name,"Author: ","")
		stirngid := e.ChildAttr("a[href]","href")
		id := strings.Split(stirngid,".php?id=")
		var singleMap = make(map[string]interface{})
		singleMap["id"] = id[1]
		singleMap["type"] = id[0]
		singleMap["name"] = name;
		authorData = append(authorData, singleMap)
	})
	resolutionData := make([]map[string]interface{},0,0)
	c.OnHTML(".download-button",func(e *colly.HTMLElement){
		id := e.Attr("data-id")
		typeImage := e.Attr("data-type")
		server := e.Attr("data-server")
		userid := e.Attr("data-user-id")
		resolution := e.Text
		resolution = strings.Trim(resolution," ")
		resolution = strings.ReplaceAll(resolution,"\n","")
		resolution = strings.ReplaceAll(resolution,"Download Original ","")
		resolution = strings.Trim(resolution," ")
		var singleMap = make(map[string]interface{})
		singleMap["id"] = id
		singleMap["type"] = typeImage
		singleMap["server"] = server
		singleMap["userid"] = userid
		singleMap["resolution"] = resolution
		resolutionData = append(resolutionData, singleMap)
	})
	nameWallpaepr := ""
	c.OnHTML(".main-content",func(e *colly.HTMLElement){
		name := e.Attr("title")
		nameWallpaepr = name
	})

	c.Visit("https://"+typeWall+".alphacoders.com"+"/big.php?i="+id)
	m := Detail{
		nameWallpaepr,resolutionData,authorData,imageData,categolyData,tagData,
	}
	return m
}
