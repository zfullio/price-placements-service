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

func (fr feedRepository) Get(spreadsheetID string) (feeds []entity.Feed, err error) {
	fr.logger.Trace().Msg("Get")

	readRange := "// Фиды!A2:C"
	resp, err := fr.client.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("api error: %w", err)
	}

	for _, row := range resp.Values {
		placement, ok := row[0].(string)
		if !ok {
			return nil, fmt.Errorf("не могу обработать площадку в gsheets")
		}

		url, ok := row[1].(string)
		if !ok {
			return nil, fmt.Errorf("не могу обработать url в gsheets")
		}

		developer, ok := row[2].(string)
		if !ok {
			return nil, fmt.Errorf("не могу обработать застройщика в gsheets")
		}

		feed := entity.Feed{
			Url:       url,
			Placement: entity.Placement(strings.ToLower(placement)),
			Developer: developer,
		}
		feeds = append(feeds, feed)
	}

	return feeds, err
}
