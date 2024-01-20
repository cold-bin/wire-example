package service

import (
	"github.com/IBM/sarama"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"wire-example/dal"
	"wire-example/entity"
)

type ItemSvc struct {
	Rdb    *redis.Client
	Kmq    sarama.SyncProducer
	Logger *zap.SugaredLogger

	Itemdal *dal.ItemDal
}

func (is *ItemSvc) Update(e *entity.Item) (err error) {
	return is.Itemdal.Update(e)
}
