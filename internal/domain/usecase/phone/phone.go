package phone

import (
	"checkphones/internal/domain/entity"
	"fmt"
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
}

func NewPhoneUseCase(phoneService PhService, lotService LotService) *phoneUseCase {
	return &phoneUseCase{phoneService: phoneService, lotService: lotService}
}

func (u phoneUseCase) CheckNumbers(spreadsheetID string, feedUrl string, placement entity.Placement) (result []string, err error) {
	phones, err := u.phoneService.Get(spreadsheetID)
	if err != nil {
		return result, err
	}
	lots, err := u.lotService.Get(feedUrl)
	if err != nil {
		return result, err
	}
	objectsFromPhones := matchObjectFromPhones(phones, placement)
	if objectsFromPhones == nil {
		return result, fmt.Errorf("не могу найти согласованные номера для площадки: %s", placement)
	}
	objectsFromLots := matchObjectFromLots(lots)
	for obj, v := range objectsFromLots {
		lotObjectNums := v
		phoneObjectNums, ok := objectsFromPhones[obj]
		if !ok {
			fmt.Printf("Не нахожу ЖК %s в согласованных номерах\n", obj)
			continue
		}
		for _, lotObjectNum := range lotObjectNums {
			for _, phoneObjectNum := range phoneObjectNums {
				if lotObjectNum != phoneObjectNum {
					result = append(result, fmt.Sprintf("Объект: %s. Ожидалось: %v. Получено: %v", obj, phoneObjectNum, lotObjectNum))
				}
			}
		}
	}
	return result, err
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
