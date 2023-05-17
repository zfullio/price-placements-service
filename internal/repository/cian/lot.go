package cian

import (
	"fmt"
	"github.com/rs/zerolog"
	placementsFeeds "github.com/zfullio/price-placements"
	"github.com/zfullio/price-placements-service/internal/domain/entity"
	"regexp"
	"strconv"
	"strings"
)

type lotRepository struct {
	client placementsFeeds.CianFeed
	logger *zerolog.Logger
}

func NewLotRepository(logger *zerolog.Logger) *lotRepository {
	repoLogger := logger.With().Str("repo", "lot").Str("type", "cian").Logger()

	return &lotRepository{
		client: placementsFeeds.CianFeed{},
		logger: &repoLogger,
	}
}

func (lr lotRepository) Get(url string) (lots []entity.Lot, err error) {
	lr.logger.Trace().Str("url", url).Msg("Get")

	err = lr.client.Get(url)
	if err != nil {
		return lots, fmt.Errorf("не могу разобрать фид: %w", err)
	}

	for i := 0; i < len(lr.client.Object); i++ {
		phone := lr.client.Object[i].Phones.PhoneSchema.CountryCode + lr.client.Object[i].Phones.PhoneSchema.Number

		onlyDigitInt, err := phoneNumberToInt(phone)
		if err != nil {
			return nil, fmt.Errorf("не могу сконвертировать номер из фида '%s' в число", phone)
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
	lr.logger.Trace().Str("url", url).Msg("Validate")

	err = lr.client.Get(url)
	if err != nil {
		return results, fmt.Errorf("не могу разобрать фид: %w", err)
	}

	results = append(results, lr.client.Check()...)

	return results, err
}
