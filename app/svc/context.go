package svc

import (
	"auto/app/model/house"
	"auto/app/model/user"
	"auto/pkg/conf"

	"auto/pkg/log"
	"auto/pkg/mgo"
)

type ServiceContext struct {
	Config     conf.Config
	Mongo      *mgo.MongoModel
	UserModel  *user.Model
	HouseModel *house.Model
}

func NewServiceContext(c conf.Config) *ServiceContext {
	mongoModel, err := mgo.NewMongoModel(c.Mongo, c.Redis)
	if err != nil {
		log.Errorf("NewMongoModel err: %v", err)
		panic(err)
	}

	return &ServiceContext{
		Config:     c,
		UserModel:  user.NewModel(mongoModel),
		HouseModel: house.NewModel(mongoModel),
	}
}
