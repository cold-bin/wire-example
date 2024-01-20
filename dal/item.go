package dal

import (
	"wire-example/entity"
)

type ItemDal struct {
	*Base
}

func (orderdal *ItemDal) Update(e *entity.Item) error {
	c := &entity.Item{}
	return orderdal.gdb.Model(c).Save(e).Error
}
