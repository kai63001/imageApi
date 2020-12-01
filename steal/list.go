package steal

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type List struct {
	NowPage string
	AllPage string
	Image []map[string]interface{}
}

func ListSteal(typeWall string,path string,page string,orderby string) List{
	k := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),

	);

	k.OnRequest(func(r *colly.Request) {
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		r.Headers.Set("Referrer", "https://wall.alphacoder.com/")
		r.Headers.Set("Cookie","__cfduid=d6a0bfdcf6c048bb47daff541854c2da71606293961; wa_session=6it4th5p6lh27hic2f6kn47qtsqmbtnsrc7m33fmppijei3rhrh18and37npsmeg3ak49f8p6jl0ipbhcvtv5e22eh52jip0k8kbd61; _ga=GA1.2.627310865.1606293964; cookieconsent_status=allow; _pbjs_userid_consent_data=3524755945110770; Sorting=newest; AlphaCodersView=paged; _gid=GA1.2.453286881.1606633263; cto_bidid=4ZePIl95SmFyeXBlUUNzV3Vmb25LMXhHb3liYzNGNTk5ZiUyQnZmQWVIb1JhUyUyRmpUanFmQTAyRkFEbm5BNEtDazloT3p0Vm5KaXUzZEVoQnclMkYwdFo1blZBUE1WMUZVRlUxOTZoM3pTN252VDh5RE1meHR4Q0Q3JTJCbjU2Wm1hOXVxWVZwNlR4; cto_bundle=wP11YF9zJTJGYzJqOVUxSFU5cWdhZzEzRE1vVXVtblVnMGhkYjBEWjNLZ01RMXVSanJ6cmttNzM1a052MTZnOWNBJTJGNjg0TW43d25ZVzFQNlVPbUh5enZzNmJvM2xZN2l2UExEJTJGJTJGT1pqTkswUndZRlZLNlFJR3ROJTJCNVFPOVFlcHV0bnZzeUdTdk1SOGc4TFU3a2d3UG8lMkZNSk8lMkZZUSUzRCUzRA; __cf_bm=ba5f1c0826f4197330462a4d2f98aebd7d45e7c4-1606808153-1800-AQ+1yza5X1swQ6GrB0saF31+XEkYzXY8rWkYmTYHmEb3qHk8q3HIZvpi9MMurhIUFz7wyRozqazmH450CX6Rw+MFeexbgDlqbxg+MV8z29J0S4HKqmXpKbxjcU5N8mMyKoENX2EvG9sR4v4c92ZMb3c=; session_depth=wall.alphacoders.com%3D4%7C860732765%3D4; hbcm_sd=1%7C1606808153231; _gat_gtag_UA_281956_21=1; __gads=ID=3c2f0c6581281639:T=1606806869:S=ALNI_MYjvlAsF3Tu03DOmCK0CQyIdszS-A")
	})
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
		page = "Err"+err.Error()
	})
	k.Visit(uri)
	m := List{
		page,myPage,imageData,
	}
	return m
}
