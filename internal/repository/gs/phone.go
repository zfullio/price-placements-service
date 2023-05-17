package gs

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/zfullio/price-placements-service/internal/domain/entity"
	"google.golang.org/api/sheets/v4"
	"strconv"
	"strings"
)

type phoneRepository struct {
	client sheets.Service
	logger *zerolog.Logger
}

func NewPhoneRepository(client sheets.Service, logger *zerolog.Logger) *phoneRepository {
	repoLogger := logger.With().Str("repo", "phone").Str("type", "sheets").Logger()

	return &phoneRepository{
		client: client,
		logger: &repoLogger,
	}
}

func (pr phoneRepository) Get(spreadsheetID string, developer string) (phones []entity.Phone, err error) {
	pr.logger.Trace().Msg("Get")

	readRange := "// Номера!A2:E"

	resp, err := pr.client.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("api error: %w", err)
	}

	for _, row := range resp.Values {
		client, ok := row[0].(string)
		if !ok {
			return nil, fmt.Errorf("не могу получить: 'застройщик' в gsheets")
		}

		client = strings.ToLower(client)

		object, ok := row[1].(string)
		if !ok {
			return nil, fmt.Errorf("не могу получить: 'жк' в gsheets")
		}

		placementStr, ok := row[2].(string)
		if !ok {
			return nil, fmt.Errorf("не могу получить: 'Портал' в gsheets")
		}

		placement, err := matchPlacement(placementStr)
		if err != nil {
			return nil, fmt.Errorf("не могу обработать: 'Портал' в gsheets: %w", err)
		}

		number, ok := row[3].(string)
		if !ok {
			return nil, fmt.Errorf("не могу получить: 'выделенный номер' в gsheets")
		}

		convNumber, err := strconv.Atoi(number)
		if err != nil {
			return nil, fmt.Errorf("не могу сконвертировать номер: '%s' в число в gsheets", number)
		}

		variationsClear := make([]string, 0)

		if len(row) > 4 {
			variationsRaw := strings.Split(row[4].(string), ",")
			for _, item := range variationsRaw {
				variationsClear = append(variationsClear, optimizeObject(item))
			}
		}

		if client == strings.ToLower(developer) {
			phone := entity.Phone{
				Number:         convNumber,
				Developer:      client,
				Object:         optimizeObject(object),
				ObjectExtended: variationsClear,
				Placement:      placement,
			}
			phones = append(phones, phone)
		}
	}

	return phones, err
}
