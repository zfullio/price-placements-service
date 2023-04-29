package phone

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/zfullio/price-placements-service/internal/domain/entity"
	"strings"
)

type PhService interface {
	Get(spreadsheetID string, developer string) (phones []entity.Phone, err error)
}

type LotService interface {
	Get(url string) (lots []entity.Lot, err error)
}

type UseCase struct {
	phoneService PhService
	lotService   LotService
	logger       *zerolog.Logger
}

func NewPhoneUseCase(phoneService PhService, lotService LotService, logger *zerolog.Logger) *UseCase {
	useCaseLogger := logger.With().Str("useCase", "phone").Logger()

	return &UseCase{
		phoneService: phoneService,
		lotService:   lotService,
		logger:       &useCaseLogger,
	}
}

func (u UseCase) CheckNumbers(ctx context.Context, spreadsheetID string, developer string, feedUrl string,
	placement entity.Placement) ([]entity.CheckResult, error) {

	u.logger.Trace().Msg("CheckNumbers")

	result := make([]entity.CheckResult, 0)
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		developers := strings.Split(developer, "/")
		var phones []entity.Phone
		for _, dev := range developers {
			dev = strings.TrimSpace(dev)
			phonesDev, err := u.phoneService.Get(spreadsheetID, dev)
			if err != nil {
				result = append(result, entity.CheckResult{
					Developer: developer,
					Placement: placement,
					Base:      entity.UnknownPlacement,
					URL:       feedUrl,
					Message:   err.Error(),
					Status:    entity.Error,
				})
			}
			phones = append(phones, phonesDev...)
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
				res := entity.CheckResult{
					Developer: developer,
					Placement: placement,
					Base:      entity.UnknownPlacement,
					URL:       feedUrl,
					Message:   fmt.Sprintf("Не нахожу ЖК '%s' в согласованных номерах", obj),
					Status:    entity.Error,
				}
				result = append(result, res)
				continue
			}
			for _, lotObjectNum := range lotObjectNums {
				for _, phoneObjectNum := range phoneObjectNums {
					if lotObjectNum != phoneObjectNum {
						res := entity.CheckResult{
							Developer: developer,
							Placement: placement,
							Base:      entity.UnknownPlacement,
							URL:       feedUrl,
							Message:   fmt.Sprintf("Объект: %s / Ожидалось: %v. Получено: %v", obj, phoneObjectNum, lotObjectNum),
							Status:    entity.Warning,
						}
						result = append(result, res)
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
			for _, alt := range phones[i].ObjectExtended {
				_, ok := objects[alt]
				if !ok {
					objects[alt] = []int{phones[i].Number}
					continue
				}
				objects[phones[i].Object] = append(objects[alt], phones[i].Number)
			}
			continue
		}
		objects[phones[i].Object] = append(objects[phones[i].Object], phones[i].Number)
	}
	return objects
}
