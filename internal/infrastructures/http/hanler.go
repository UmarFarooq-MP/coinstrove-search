package http

import "coinstrove-search/internal/core/ports"

type DataSVCHandler struct {
	dataService *ports.DataSVC
}

func NewDataSVCHandler(dataService *ports.DataSVC) *DataSVCHandler {
	return &DataSVCHandler{dataService: dataService}
}
