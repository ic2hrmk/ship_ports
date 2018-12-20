package repository

import (
	"errors"
	"github.com/ic2hrmk/ship_ports/app/services/port/persistence/model"
)

var (
	ErrPortNotFound = errors.New("ErrPortNotFound")
)

type PortRepository interface {
	Save(*model.Port) (*model.Port, error)
	SaveBulk([]*model.Port) error

	FindAll(limit, offset uint64) ([]*model.Port, error)
}
