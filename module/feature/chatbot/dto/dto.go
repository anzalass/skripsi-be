package dto

import "testskripsi/module/entities"

type RequestDataset struct {
	Request []entities.DatasetAi `json:"request"`
}
