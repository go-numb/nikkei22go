package nikke22go

import (
	"net/url"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/gocolly/colly/v2"
)

type Client struct {
	c      *colly.Collector
	target *url.URL
}

func New(allowDomain ...string) *Client {
	c := colly.NewCollector(
		colly.AllowedDomains(
			allowDomain...,
		))
	return &Client{
		c:      c,
		target: &url.URL{},
	}
}

func (p *Client) SetTarget(target string, depth int) error {
	u, err := url.Parse(target)
	if err != nil {
		return err
	}

	p.target = u
	p.c.MaxDepth = depth

	return nil
}

func (p *Client) ToCSV(pathname string, v interface{}) error {
	f, err := os.Create(pathname)
	if err != nil {
		return err
	}
	defer f.Close()

	w := gocsv.DefaultCSVWriter(f)

	if err := gocsv.MarshalCSV(&v, w); err != nil {
		return err
	}
	return nil
}
