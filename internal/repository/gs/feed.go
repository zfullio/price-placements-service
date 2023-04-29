package gs

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/zfullio/price-placements-service/internal/domain/entity"
	"google.golang.org/api/sheets/v4"
	"strings"
)

type feedRepository struct {
	client sheets.Service
	logger *zerolog.Logger
}

func NewFeedRepository(client sheets.Service, logger *zerolog.Logger) *feedRepository {
	repoLogger := logger.With().Str("repo", "feed").Str("type", "sheets").Logger()

	return &feedRepository{
		client: client,
		logger: &repoLogger,
	}
}

// Get Получает все фиды которые надо проверить
func (fr feedRepository) Get(spreadsheetID string, developer string) (feeds []entity.Feed, err error) {
	fr.logger.Trace().Msg("Get")

	readRange := "// Фиды!A2:F"
	resp, err := fr.client.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("api error: %w", err)
	}

	bases, err := fr.GetBaseFeed(spreadsheetID)
	if err != nil {
		return nil, fmt.Errorf("не могу получить настройки для проверки фидов: %w", err)
	}

	for _, row := range resp.Values {
		client, ok := row[0].(string)
		if !ok {
			return nil, fmt.Errorf("не могу получить: 'застройщик' в gsheets")
		}
		client = strings.ToLower(client)

		placementStr, ok := row[1].(string)
		if !ok {
			return nil, fmt.Errorf("не могу получить: 'портал' в gsheets")
		}
		placement, err := matchPlacement(placementStr)
		if err != nil {
			return nil, fmt.Errorf("не могу обработать: 'портал' в gsheets: %w", err)
		}

		// Пропуск если фид клиентский
		if placement == entity.ClientFeed {
			continue
		}

		baseFeed, ok := bases[placement]
		if !ok {
			return nil, fmt.Errorf("не могу понять базовый тип фида для площадки %s", placementStr)
		}

		areaTypeStr, ok := row[2].(string)
		if !ok {
			return nil, fmt.Errorf("не могу получить: 'тип' в gsheets")
		}

		areaType, err := matchAreaType(areaTypeStr)
		if err != nil {
			return nil, fmt.Errorf("не могу обработать: 'тип' в gsheets: %w", err)
		}

		// Пропуск если фид для коммерции
		if areaType == entity.COMMERCE {
			continue
		}

		statusStr, ok := row[4].(string)
		if !ok {
			return nil, fmt.Errorf("не могу получить: 'статус' в gsheets")
		}

		status, err := matchStatus(statusStr)
		if err != nil {
			return nil, fmt.Errorf("не могу обработать: 'статус' в gsheets: %w", err)
		}

		// Пропуск если фид не активен
		if status != entity.Active {
			continue
		}

		url, ok := row[5].(string)
		if !ok {
			return nil, fmt.Errorf("не могу обработать url в gsheets")
		}
		url = strings.TrimSpace(url)

		if client == strings.ToLower(developer) {
			feed := entity.Feed{
				Developer: client,
				Placement: placement,
				AreaType:  areaType,
				Status:    status,
				URL:       url,
				BaseFeed:  baseFeed,
			}
			feeds = append(feeds, feed)
		}
	}

	return feeds, err
}

// GetBaseFeed Получает правила обработки фидов различных типов
func (fr feedRepository) GetBaseFeed(spreadsheetID string) (map[entity.Placement]entity.Placement, error) {
	fr.logger.Trace().Msg("GetBaseFeed")

	readRange := "// Как проверять!A2:B"
	resp, err := fr.client.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("api error: %w", err)
	}

	bases := make(map[entity.Placement]entity.Placement, len(resp.Values))
	for _, row := range resp.Values {
		placementStr, ok := row[0].(string)
		if !ok {
			return nil, fmt.Errorf("не могу получить: 'Портал' в gsheets")
		}

		placement, err := matchPlacement(placementStr)
		if err != nil {
			return nil, fmt.Errorf("не могу обработать: 'Портал' в gsheets: %w", err)
		}

		baseStr, ok := row[0].(string)
		if !ok {
			return nil, fmt.Errorf("не могу получить: 'Использует фид' в gsheets")
		}

		base, err := matchPlacement(baseStr)
		if err != nil {
			return nil, fmt.Errorf("не могу обработать: 'Использует фид' в gsheets: %w", err)
		}

		if _, ok := bases[placement]; !ok {
			bases[placement] = base
		} else {
			return nil, fmt.Errorf("дублирующиеся значения '%s' 'Использует фид' в gsheets", placementStr)
		}
	}

	return bases, err
}
