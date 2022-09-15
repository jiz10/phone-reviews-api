package smartphone

import (
	"phone-reviews-api/internal/pkg/repository/database"
)

type CreateGateway interface {
	Create(cmd *CreateSmartphoneCMD) (*Smartphone, error)
}

type CreateGtw struct {
	StorageGateway StorageGateway
}

func NewSmartphoneCreateGateway(client *database.MySqlClient) CreateGateway {
	return &CreateGtw{StorageGateway: &Storage{client}}
}

func (c CreateGtw) Create(cmd *CreateSmartphoneCMD) (*Smartphone, error) {
	return c.StorageGateway.create(cmd)
}
