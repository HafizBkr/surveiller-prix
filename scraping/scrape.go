package scraping

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Fonction pour effectuer le scraping des noms et des prix des ordinateurs sur Amazon
func ScrapeAmazonComputers(category string) (map[string]string, error) {
	results := make(map[string]string)

	url := fmt.Sprintf("https://www.amazon.com/s?k=%s", strings.ReplaceAll(category, " ", "+"))
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	doc.Find(".s-result-item").Each(func(i int, s *goquery.Selection) {
		name := s.Find(".a-size-medium").Text()
		price := s.Find(".a-price .a-offscreen").Text()

		results[name] = price
	})

	return results, nil
}

// Fonction pour effectuer le scraping des noms et des prix des ordinateurs sur Coin Afrique
func ScrapeCoinAfriqueComputers(category string) (map[string]string, error) {
	results := make(map[string]string)

	url := fmt.Sprintf("https://www.coinafrique.com/q/%s", strings.ReplaceAll(category, " ", "-"))
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	doc.Find(".listing-card").Each(func(i int, s *goquery.Selection) {
		name := s.Find(".product-title").Text()
		price := s.Find(".product-price").Text()

		results[name] = price
	})

	return results, nil
}
