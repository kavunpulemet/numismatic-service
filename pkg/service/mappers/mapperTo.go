package mappers

import (
	"NumismaticClubApi/models"
	dbmodels "NumismaticClubApi/pkg/database/models"
)

func MapToUpdateCoin(serviceCoin models.Coin) dbmodels.UpdateCoin {
	return dbmodels.UpdateCoin{
		Name:           serviceCoin.Name,
		Country:        serviceCoin.Country,
		Year:           serviceCoin.Year,
		Denomination:   serviceCoin.Denomination,
		Material:       serviceCoin.Material,
		Weight:         serviceCoin.Weight,
		Diameter:       serviceCoin.Diameter,
		Thickness:      serviceCoin.Thickness,
		Condition:      serviceCoin.Condition,
		MintMark:       serviceCoin.MintMark,
		HistoricalInfo: serviceCoin.HistoricalInfo,
		Value:          serviceCoin.Value,
	}
}
