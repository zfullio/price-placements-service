package avito

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
	client placementsFeeds.AvitoFeed
	logger *zerolog.Logger
}

func NewLotRepository(logger *zerolog.Logger) *lotRepository {
	repoLogger := logger.With().Str("repo", "lot").Str("type", "avito").Logger()
	return &lotRepository{
		client: placementsFeeds.AvitoFeed{},
		logger: &repoLogger,
	}
}

func (lr lotRepository) Get(url string) (lots []entity.Lot, err error) {
	lr.logger.Trace().Str("url", url).Msg("Get")

	err = lr.client.Get(url)
	if err != nil {
		return lots, fmt.Errorf("не могу разобрать фид: %w", err)
	}

	developments, err := lr.getObjectMap()
	if err != nil {
		return lots, fmt.Errorf("не могу получить словарь DevelopmentID: %w", err)
	}

	for i := 0; i < len(lr.client.Ad); i++ {
		phone := lr.client.Ad[i].ContactPhone
		onlyDigitInt, err := phoneNumberToInt(phone)
		if err != nil {
			return nil, fmt.Errorf("не могу сконвертировать номер из фида '%s' в число", phone)
		}
		name, ok := developments[strings.TrimSpace(lr.client.Ad[i].NewDevelopmentId)]
		if !ok {
			return lots, fmt.Errorf("не могу сопоставить DevelopmentID: '%s'. AD: '%s'", lr.client.Ad[i].NewDevelopmentId, lr.client.Ad[i].ID)
		}
		lot := entity.Lot{
			ID:     lr.client.Ad[i].ID,
			Object: optimizeObject(name),
			Phone:  onlyDigitInt,
		}
		lots = append(lots, lot)
	}

	return lots, err
}

// Получает словарь объектов
func (lr lotRepository) getObjectMap() (map[string]string, error) {
	lr.logger.Trace().Msg("getObjectMap")

	objectMap := make(map[string]string, 0)
	developments, err := lr.client.GetDevelopments()
	if err != nil {
		return objectMap, err
	}

	for _, region := range developments.Region {
		for _, city := range region.City {
			for _, object := range city.Object {
				if len(object.Housing) == 0 {
					if _, ok := objectMap[strings.TrimSpace(object.ID)]; ok != true {
						objectMap[object.ID] = object.Name
					}
				} else {
					for _, housing := range object.Housing {
						if _, ok := objectMap[housing.ID]; ok != true {
							objectMap[housing.ID] = object.Name
						}
					}
				}
			}
		}
	}
	return objectMap, nil
}

// Приводит имя объекта к нужному виду
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

// Преобразует телефонный номер в число
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
