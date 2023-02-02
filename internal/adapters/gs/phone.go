package gs

import (
	"fmt"
	"google.golang.org/api/sheets/v4"
	"price-placements-service/internal/domain/entity"
	"strconv"
	"strings"
)

type phoneRepository struct {
	client sheets.Service
}

func NewPhoneRepository(client sheets.Service) *phoneRepository {
	return &phoneRepository{client: client}
}

func (pr phoneRepository) Get(spreadsheetID string) (phones []entity.Phone, err error) {
	readRange := "// Номера!A2:D"
	resp, err := pr.client.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("api error: %w", err)
	}
	for _, row := range resp.Values {
		object, ok := row[0].(string)
		if !ok {
			return nil, fmt.Errorf("не могу обработать объект в gsheets")
		}

		placement, ok := row[1].(string)
		if !ok {
			return nil, fmt.Errorf("не могу обработать площадку в gsheets")
		}

		number, ok := row[2].(string)
		if !ok {
			return nil, fmt.Errorf("не могу получить номер в gsheets")
		}
		convNumber, err := strconv.Atoi(number)
		if err != nil {
			return nil, fmt.Errorf("не могу сконвертировать номер: '%s' в число в gsheets", number)

		}

		developer, ok := row[3].(string)
		if !ok {
			return nil, fmt.Errorf("не могу обработать застройщика в gsheets")
		}

		phone := entity.Phone{
			Number:    convNumber,
			Developer: developer,
			Object:    optimizeObject(object),
			Placement: entity.Placement(strings.ToLower(placement)),
		}
		phones = append(phones, phone)
	}
	return phones, err
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
