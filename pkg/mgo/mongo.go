package mgo

import (
	"auto/pkg/cache"
	"auto/pkg/conf"
	"auto/pkg/log"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoModel struct {
	*mongo.Database
	*cache.Cache
}

// 创建 collection
func (m *MongoModel) InitCollection(name string) *mongo.Collection {
	collectionOptions := options.CreateCollection().SetCapped(true).SetSizeInBytes(1024 * 1024)
	// 如果存在则忽略
	log.Infof("CreateCollection: %s before", name)

	err := m.CreateCollection(context.Background(), name, collectionOptions)
	log.Infof("CreateCollection: %s", name)
	if err != nil {
		log.Errorf("CreateCollection err: %v", err)
	}
	return m.Collection(name)
}

// 默认的model 都带缓存
func NewMongoModel(mc conf.MongoConfig, cc conf.CacheConfig) (md *MongoModel, err error) {
	md = &MongoModel{}
	log.Infof("Start connect mongo: %s", mc.Url)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mc.Url))
	if err != nil {
		return
	}

	md.Database = client.Database(mc.DB)
	log.Infof("Get mongo database: %s", mc.DB)
	md.Cache = cache.NewCache(cc)
	log.Infof("Get cache")
	return
}

func (m *MongoModel) AddUniKey(coll *mongo.Collection, keys ...string) {

	var compoundKeys bson.D

	for _, key := range keys {
		compoundKeys = append(compoundKeys, bson.E{Key: key, Value: 1})
	}
	idxModel := mongo.IndexModel{
		Keys:    compoundKeys,
		Options: options.Index().SetUnique(true),
	}
	coll.Indexes().CreateOne(context.Background(), idxModel)
}
