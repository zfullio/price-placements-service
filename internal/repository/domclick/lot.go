package domclick

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
	client placementsFeeds.DomclickFeed
	logger *zerolog.Logger
}

func NewLotRepository(logger *zerolog.Logger) *lotRepository {
	repoLogger := logger.With().Str("repo", "lot").Str("type", "domclick").Logger()

	return &lotRepository{
		client: placementsFeeds.DomclickFeed{},
		logger: &repoLogger,
	}
}

func (lr lotRepository) Get(url string) (lots []entity.Lot, err error) {
	lr.logger.Trace().Str("url", url).Msg("Get")

	err = lr.client.Get(url)
	if err != nil {
		return lots, fmt.Errorf("не могу разобрать фид: %w", err)
	}

	onlyDigitInt, err := phoneNumberToInt(lr.client.Complex.SalesInfo.SalesPhone)
	lot := entity.Lot{
		ID:     lr.client.Complex.ID,
		Object: optimizeObject(lr.client.Complex.Name),
		Phone:  onlyDigitInt,
	}
	lots = append(lots, lot)

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
	lr.logger.Trace().Str("url", url).Msg("Get")

	err = lr.client.Get(url)
	if err != nil {
		return results, fmt.Errorf("не могу разобрать фид: %w", err)
	}

	results = append(results, lr.client.Check()...)

	return results, err
}
