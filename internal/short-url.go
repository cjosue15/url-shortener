package shortenurl

import (
	"fmt"
	"math/rand"

	"gorm.io/gorm"
)

type ShortUrl struct {
	gorm.Model
	Url   string
	Short string
}

type ShortenUrl struct {
	Database *gorm.DB
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func NewShortenUrl(database *gorm.DB) *ShortenUrl {
	return &ShortenUrl{Database: database}
}

func (s *ShortenUrl) CreateShortUrl(originalUrl string) (*string, error) {
	short := randSeq(5)

	url := ShortUrl{Url: originalUrl, Short: short}

	result := s.Database.Create(&url)

	if result.Error != nil {
		return nil, fmt.Errorf("Error creating short URL")
	}

	return &short, nil
}

func (s *ShortenUrl) GetOriginalUrl(short string) (*string, error) {
	var data ShortUrl

	result := s.Database.Where("short = ?", short).First(&data)

	if result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("URL not found")
		}

		return nil, fmt.Errorf("Error getting the URL")
	}

	return &data.Url, nil
}
