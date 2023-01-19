package nikke22go

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type NIkkei225Code struct {
	Name     string `csv:"name"`
	Code     string `csv:"code"`
	Category string `csv:"category"`
}

// 社名などを得て、検索、ウェブサイトを訪問し、サイト内の「@」付き情報を取得する
func (p *Client) FindNikkei225Code() []NIkkei225Code {
	var (
		data = make([]NIkkei225Code, 0)
	)

	p.c.OnHTML(".container", func(e *colly.HTMLElement) {
		var category string
		// e.Request.Visit(e.Attr("href"))
		e.ForEach(".row", func(i int, h *colly.HTMLElement) {
			if h.ChildText(".col-sm-11") != "" {
				category = h.ChildText(".col-sm-11")
				fmt.Println(h.Text)
			}

			if h.ChildText(".col-sm-1_5") != "" && h.ChildText(".col-sm-1_5") != "コード" {
				data = append(data, NIkkei225Code{
					Name:     h.ChildText(".col-sm-8"),
					Code:     h.ChildText(".col-sm-1_5"),
					Category: category,
				})
			}

		})
	})

	p.c.Visit(p.target.String())

	return data
}
