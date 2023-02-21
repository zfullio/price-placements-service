package phone

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/zfullio/price-placements-service/internal/domain/entity"
)

type PhService interface {
	Get(spreadsheetID string) (phones []entity.Phone, err error)
}

type LotService interface {
	Get(url string) (lots []entity.Lot, err error)
}

type phoneUseCase struct {
	phoneService PhService
	lotService   LotService
	logger       *zerolog.Logger
}

func NewPhoneUseCase(phoneService PhService, lotService LotService, logger *zerolog.Logger) *phoneUseCase {
	useCaseLogger := logger.With().Str("usecase", "phone").Logger()

	return &phoneUseCase{
		phoneService: phoneService,
		lotService:   lotService,
		logger:       &useCaseLogger,
	}
}

func (u phoneUseCase) CheckNumbers(ctx context.Context, spreadsheetID string, feedUrl string, placement entity.Placement) ([]string, error) {
	u.logger.Trace().Msg("CheckNumbers")
	result := make([]string, 0)
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		phones, err := u.phoneService.Get(spreadsheetID)
		if err != nil {
			return nil, err
		}

		lots, err := u.lotService.Get(feedUrl)
		if err != nil {
			return nil, err
		}

		objectsFromPhones := matchObjectFromPhones(phones, placement)
		if objectsFromPhones == nil {
			return nil, fmt.Errorf("не могу найти согласованные номера для площадки: %s", placement)
		}

		objectsFromLots := matchObjectFromLots(lots)
		for obj, v := range objectsFromLots {
			lotObjectNums := v
			phoneObjectNums, ok := objectsFromPhones[obj]
			if !ok {
				result = append(result, fmt.Sprintf("%s: %s Не нахожу ЖК '%s' в согласованных номерах\n", placement, feedUrl, obj))
				continue
			}
			for _, lotObjectNum := range lotObjectNums {
				for _, phoneObjectNum := range phoneObjectNums {
					if lotObjectNum != phoneObjectNum {
						result = append(result, fmt.Sprintf("Площадка %s. Объект: %s. Ожидалось: %v. Получено: %v", placement, obj, phoneObjectNum, lotObjectNum))
					}
				}
			}
		}
	}
	return result, nil
}

func matchObjectFromLots(lots []entity.Lot) map[string][]int {
	objects := make(map[string][]int)
	for i := 0; i < len(lots); i++ {
		_, ok := objects[lots[i].Object]
		if !ok {
			objects[lots[i].Object] = []int{lots[i].Phone}
			continue
		}
		objects[lots[i].Object] = append(objects[lots[i].Object], lots[i].Phone)
	}
	return objects
}

func matchObjectFromPhones(phones []entity.Phone, placement entity.Placement) map[string][]int {
	objects := make(map[string][]int)
	for i := 0; i < len(phones); i++ {
		if phones[i].Placement != placement {
			continue
		}
		_, ok := objects[phones[i].Object]
		if !ok {
			objects[phones[i].Object] = []int{phones[i].Number}
			continue
		}
		objects[phones[i].Object] = append(objects[phones[i].Object], phones[i].Number)
	}
	return objects
}
