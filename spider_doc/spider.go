package main

import (
	"daily/spider_doc/db"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var host = "https://gocn.vip"
var lastLatest string
var isLatest bool

func main() {
	db.Init()

	for {
		page := 1
		isLatest = false
		for {
			reader, err := getDailyBody(page)
			if err != nil {
				fmt.Println("请求获取每日文章页面失败:" + err.Error())
				time.Sleep(time.Second * 3)
				continue
			}
			doc, err := goquery.NewDocumentFromReader(reader)
			if err != nil {
				fmt.Println("解析每日文章页面失败:" + err.Error())
				time.Sleep(time.Second * 3)
				continue
			}

			doc.Find(".media-heading>a[title*=每日新闻]").EachWithBreak(handleNode)

			fmt.Println(fmt.Sprintf("第%d页完成，准备开始下一页", page))
			if !strings.Contains(doc.Find("a[rel=next]").Last().Text(), "下一页") {
				break
			}

			if page == 1 {
				latest := doc.Find(".media-heading>a[title*=每日新闻]").First().Text()
				if latest != lastLatest {
					lastLatest = latest
				}
			}

			if isLatest {
				break
			}
			page++
		}
		fmt.Println(fmt.Sprintf("本日爬取已完成,当前最新日期为%s", lastLatest))
		time.Sleep(time.Second * 3600 * 24)
	}
}

func handleNode(i int, selection *goquery.Selection) bool {
	if selection.Text() == lastLatest {
		isLatest = true
		return false
	}
	if url, ok := selection.Attr("href"); ok {
		reader, err := getEachTopic(url)
		if err != nil {
			fmt.Println(err.Error())
		}
		doc, err := goquery.NewDocumentFromReader(reader)
		if err != nil {
			panic("解析文章页面失败:" + err.Error())
		}
		date := doc.Find("span[class=title]").First().Text()
		fmt.Println(date)

		if strings.Contains(date, "(") {
			date = strings.Split(date, "(")[1]
			date = strings.Split(date, ")")[0]
		} else if strings.Contains(date, "（") {
			date = strings.Split(date, "（")[1]
			date = strings.Split(date, "）")[0]
		} else {
			return true
		}

		doc.Find("ol>li").Each(func(i int, selection *goquery.Selection) {
			l := strings.Split(selection.Text(), "http")
			title := ""
			href := ""
			if len(l) == 2 {
				title = strings.Replace(l[0], " ", "", -1)
				href = "http" + strings.Replace(l[1], " ", "", -1)
			} else if len(l) < 2 {
				if h, ok := selection.Attr("href"); ok {
					title = strings.Replace(selection.Text(), " ", "", -1)
					href = h
				}
			}
			article := db.Articles{
				Date:  date,
				Title: title,
				Href:  href,
			}
			_, err := db.DB.InsertOne(&article)
			if err != nil {
				fmt.Println("ERROR:", err.Error())
			}
		})
	}
	return true
}

func getEachTopic(url string) (io.Reader, error) {
	res, err := doGetReq(host + url)
	if err != nil {
		return nil, err
	}
	return res.Body, err
}

func getDailyBody(page int) (io.Reader, error) {
	dailyUrl := host + fmt.Sprintf("/topics/node18?page=%d", page)
	res, err := doGetReq(dailyUrl)
	if err != nil {
		return nil, err
	}
	return res.Body, err
}

func doGetReq(url string) (*http.Response, error) {
	time.Sleep(time.Second * 3)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	//req.Header.Set("User-Agent", "PostmanRuntime/7.23.0")
	//req.Header.Set("Host", "gocn.vip")
	//req.Header.Set("Cookie", "user_id=eyJfcmFpbHMiOnsibWVzc2FnZSI6ImJuVnNiQT09IiwiZXhwIjpudWxsLCJwdXIiOiJjb29raWUudXNlcl9pZCJ9fQ%3D%3D--7be7692be2b183cb596182cc0f259343111372bc; _homeland_session=TYu7GLgXqcGSflxQrDdFWsQIxVfFB4P007PjKk0CBUqe2%2Ft%2BikyJBrP0AbZRJgBiDfIYoWZbx6v4%2BQ9GgF9WcxkT1ZTfuGZ9QBpYZ4oOJi1w%2FyX1FU2jn2wdHzeEEdAgFCAt%2BXboxjrToiCiI6FxfKeYPSHaiu%2Fndkm6deMhgHBP0ArxiW0sg5KAECKNQFM7CVOa%2FkHfli%2FxjhXWrLk0FN5wka2fuz1gIgC1Eq5Yruklen1u8%2FjoM1e14OF78Tuqvz3Wb9fqKFF1Iq8z7MIEbqvKNp8q1VeGwnpm79XK%2FQ4x4t5kXYc6XLiefqrovok%3D--fjY9XfucB5iNruW2--hscittM8KwkFTuHTCO3n1g%3D%3D")
	//req.Header.Add("if-none-match", `W/"10f3befcce377a5094fa0477e771d31b"`)
	//req.Header.Add("accept-encoding", `gzip, deflate, br`)
	//req.Header.Set("Cache-Control", `no-cache`)
	//req.Header.Set("Postman-Token", `ab87ccc4-bd79-46e2-9151-db9b31b9ccf8`)
	//req.Header.Add("sec-fetch-site", `same-origin`)
	//req.Header.Add("sec-fetch-mode", `navigate`)
	//req.Header.Add("sec-fetch-dest", `document`)
	//req.Header.Add("upgrade-insecure-requests", `1`)
	//req.Header.Set("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9`)
	//req.Header.Set("Accept-Encoding", `gzip, deflate, br`)
	cli := http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	return res, err
}
