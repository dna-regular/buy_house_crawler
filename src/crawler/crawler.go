package crawler

import "log"

type House struct {
	url      string
	price    int
	location string
	roomType string // 两室一厅/三室一厅
	date     string
}

type crawler interface {
	Init()
	NextHouse() (house House, done bool, err error)
}

var crawls []crawler

func isHouseValid(house House) bool {
	return false
}

func CrawlRegister(crawl crawler) {
	crawls = append(crawls, crawl)
}

func GetHouses() []House {
	var houses []House
	for _, crawl := range crawls {
		crawl.Init()
		for {
			house, done, err := crawl.NextHouse()
			if err != nil {
				log.Println(err)
			}
			if done {
				break
			}
			if isHouseValid(house) {
				houses = append(houses, house)
			}
		}
	}
	return houses
}
