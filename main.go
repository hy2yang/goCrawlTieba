package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	var target = `https://tieba.baidu.com/f?kw=`
	bar := []thread{}
	crawlTieba(target, &bar)
	fmt.Println("-------------------------------------------")
}

func crawlTieba(url string, repo *[]thread) {
	c := colly.NewCollector()

	c.OnHTML("#frs_list_pager", func(pagination *colly.HTMLElement) {
		var nextPage = pagination.ChildAttr(".next.pagination-item ", "href")
		if nextPage != "" {
			nextPage = "https:" + nextPage
			fmt.Println(nextPage)
			c.Visit(nextPage)
		}
	})

	c.OnHTML("#thread_list", func(e *colly.HTMLElement) {

		e.ForEach(".j_thread_list", func(_ int, t *colly.HTMLElement) {

			if !strings.Contains(t.Attr("class"), "thread_top") {

				tempThread := thread{}
				var stringTID = t.Attr("data-tid")
				tempThread.tid, _ = strconv.ParseInt(stringTID, 10, 64)
				tempThread.URL = "https://tieba.baidu.com/p/" + stringTID
				tempThread.title = t.ChildAttr("a", "title")
				tempThread.authorID = extractUID(t.ChildAttr(".tb_icon_author", "data-field"))

				*repo = append(*repo, tempThread)
			}

		})

	})

	c.Visit(url)

}

func extractUID(input string) int64 {
	var temp authorIconDataField
	input = `{"user_id":215209552}`
	json.Unmarshal([]byte(input), &temp)
	return temp.UID
}

// func crawlThread(c *colly.Collector, e *colly.HTMLElement) {

// 	c.OnHTML(".left_section", func(e *colly.HTMLElement) {
// 		tempThread := thread{}
// 		tempThread.URL = e.Request.URL.String()
// 		tempThread.title = e.ChildText(".core_title_txt")

// 		bar = append(bar, tempThread)
// 	})

// }
