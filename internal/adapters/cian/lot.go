package cian

import (
	"fmt"
	placementsFeeds "github.com/zfullio/price-placements"
	"price-placements-service/internal/domain/entity"
	"regexp"
	"strconv"
	"strings"
)

type lotRepository struct {
	client placementsFeeds.CianFeed
}

func NewLotRepository() *lotRepository {
	return &lotRepository{client: placementsFeeds.CianFeed{}}
}

func (lr lotRepository) Get(url string) (lots []entity.Lot, err error) {
	err = lr.client.Get(url)
	if err != nil {
		return lots, fmt.Errorf("не могу разобрать фид: %w", err)
	}
	for i := 0; i < len(lr.client.Object); i++ {
		onlyDigitInt, err := phoneNumberToInt(lr.client.Object[i].Phones.PhoneSchema.CountryCode + lr.client.Object[i].Phones.PhoneSchema.Number)
		if err != nil {
			return nil, fmt.Errorf("не могу сконвертировать номер в число")
		}

		lot := entity.Lot{
			ID:     lr.client.Object[i].ExternalId,
			Object: optimizeObject(lr.client.Object[i].JKSchema.Name),
			Phone:  onlyDigitInt,
		}
		lots = append(lots, lot)
	}
	return lots, err
}

func optimizeObject(str string) string {
	result := strings.ToLower(str)
	result = strings.ReplaceAll(result, "\"", "")
	result = strings.ReplaceAll(result, "«", "")
	result = strings.ReplaceAll(result, "»", "")
	result = strings.TrimPrefix(result, "жк ")
	result = strings.TrimPrefix(result, "ск ")
	result = strings.ReplaceAll(result, "апарт-комплекс", "")
	result = strings.ReplaceAll(result, "сити-комплекс", "")
	result = strings.TrimSpace(result)
	return result
}

func phoneNumberToInt(str string) (phone int, err error) {
	re := regexp.MustCompile(`\D+`)
	res := re.ReplaceAllString(str, "")
	re = regexp.MustCompile(`([0-9]{11})`)
	onlyDigitStr := string(re.Find([]byte(res)))
	phone, err = strconv.Atoi(onlyDigitStr)
	if err != nil {
		return phone, fmt.Errorf("не могу сконвертировать номер в число")
	}
	return phone, err
}

func (lr lotRepository) Validate(url string) (results []string, err error) {
	err = lr.client.Get(url)
	if err != nil {
		return results, fmt.Errorf("не могу разобрать фид: %w", err)
	}
	results = append(results, lr.client.Check()...)
	return results, err
}
