package smartphone

import (
	"phone-reviews-api/internal/lib/logs"
	"phone-reviews-api/internal/pkg/repository/database"
)

type StorageGateway interface {
	create(cmd *CreateSmartphoneCMD) (*Smartphone, error)
}

type Storage struct {
	*database.MySqlClient
}

func (s *Storage) create(cmd *CreateSmartphoneCMD) (*Smartphone, error) {
	tx, err := s.MySqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot create transaction")
		return nil, err
	}

	res, err := tx.Exec(`insert into smartphone (name, price, country_origin, os) 
	values (?, ?, ?, ?)`, cmd.Name, cmd.Price, cmd.CountryOrigin, cmd.OS)

	if err != nil {
		logs.Log().Error("cannot execute statement")
		_ = tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		logs.Log().Error("cannot fetch last id")
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return &Smartphone{
		Id:            id,
		Name:          cmd.Name,
		Price:         cmd.Price,
		CountryOrigin: cmd.CountryOrigin,
		OS:            cmd.OS,
	}, nil
}
