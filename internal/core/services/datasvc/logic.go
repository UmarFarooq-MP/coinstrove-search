package datasvc

import "coinstrove-search/internal/core/ports"

type newDataSVC struct {
	dataServiceRepo ports.DataSVCRepository
}

func NewDataService(dataServiceRepo ports.DataSVCRepository) ports.DataSVC {
	return &newDataSVC{dataServiceRepo: dataServiceRepo}
}

func (dataSVC *newDataSVC) GetCoinDetails()     {}
func (dataSVC *newDataSVC) GetListOfCoins()     {}
func (dataSVC *newDataSVC) GetListOfExchanges() {}
func (dataSVC *newDataSVC) UpdateDB()           {}
