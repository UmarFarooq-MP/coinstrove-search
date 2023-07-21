package datasvc

import "coinstrove-search/internal/core/ports"

type newDataSVC struct {
}

func NewDataService() ports.DataSVC {
	return &newDataSVC{}
}

func (dataSVC *newDataSVC) GetCoinDetails()     {}
func (dataSVC *newDataSVC) GetListOfCoins()     {}
func (dataSVC *newDataSVC) GetListOfExchanges() {}
func (dataSVC *newDataSVC) UpdateDB()           {}
