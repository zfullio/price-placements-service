package gs

import (
	"fmt"
	"github.com/zfullio/price-placements-service/internal/domain/entity"
	"strings"
)

const msgNotFound = "Not Found"

func matchPlacement(placement string) (entity.Placement, error) {
	placement = strings.ToLower(placement)
	placement = strings.TrimSpace(placement)

	for _, item := range entity.PlacementsAll {
		if placement == string(item) {
			return item, nil
		}
	}

	return msgNotFound, fmt.Errorf("unknown placement: %s", placement)
}

func matchAreaType(areaType string) (entity.AreaType, error) {
	areaType = strings.ToLower(areaType)
	areaType = strings.TrimSpace(areaType)
	switch areaType {
	case "жилая":
		return entity.LIVING, nil
	case "нежилая":
		return entity.COMMERCE, nil
	default:
		return msgNotFound, fmt.Errorf("unknown area type: %s", areaType)
	}
}

func matchStatus(status string) (entity.StatusFeed, error) {
	status = strings.ToLower(status)
	status = strings.TrimSpace(status)
	switch status {
	case "активен":
		return entity.Active, nil
	case "неактивен":
		return entity.Inactive, nil
	case "неизвестно":
		return entity.Unknown, nil
	case "удалено":
		return entity.Deleted, nil
	default:
		return msgNotFound, fmt.Errorf("unknown status: %s", status)
	}
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
