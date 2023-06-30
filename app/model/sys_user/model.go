package sys_user

import (
	"auto/pkg/e"
	"auto/pkg/mgo"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Model struct {
	name       string
	primaryKey string
	secondKey  string
	coll       *mongo.Collection
	d          *mgo.MongoModel
}

func NewModel(md *mgo.MongoModel) *Model {
	name := "sys_user"
	// 创建collection

	m := &Model{
		name:       name,
		d:          md,
		coll:       md.InitCollection(name),
		primaryKey: fmt.Sprintf("cache:%s:id:", name),
		secondKey:  fmt.Sprintf("cache:%s:phone:", name),
	}
	// 创建唯一索引
	m.d.AddUniKey(m.coll, "phone")

	return m
}

// 加载key
func (m *Model) loadKeys(data *SystemUser) (err error) {
	if err = m.d.Set(m.d.GetKey(m.primaryKey, data.Id), data); err != nil {
		return
	}
	if err = m.d.Set(m.d.GetKey(m.secondKey, data.Phone), data.Id); err != nil {
		return
	}
	return
}

// 清空Key
func (m *Model) clearKeys(data *SystemUser) (err error) {
	if err = m.d.Del(m.d.GetKey(m.primaryKey, data.Id)); err != nil {
		return
	}
	if err = m.d.Del(m.d.GetKey(m.secondKey, data.Phone)); err != nil {
		return
	}
	return
}

// 获取
func (m *Model) FindOne(id string) (data *SystemUser, err error) {
	data = &SystemUser{}
	key := m.d.GetKey(m.primaryKey, id)

	err = m.d.Take(
		key,
		data,
		// 查询
		func(d interface{}) (err error) {
			mid, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return
			}
			err = m.coll.FindOne(context.Background(), bson.M{"_id": mid}).Decode(d)
			return
		},
		// 加载
		func(d interface{}) (err error) {
			err = m.loadKeys(d.(*SystemUser))
			return
		},
	)

	return
}

// 通过唯一key查找
func (m *Model) FindByPhone(phone string) (data *SystemUser, err error) {
	data = &SystemUser{}

	key := m.d.GetKey(m.secondKey, phone)

	err = m.d.Take(
		key,
		data,
		// 查询
		func(d interface{}) (err error) {
			var id string
			if err = m.d.Get(key, &id); err != nil {
				d, err = m.FindOne(id)
				return
			}
			err = m.coll.FindOne(context.Background(), bson.M{"phone": phone}).Decode(d)
			return
		},
		// 加载
		func(d interface{}) (err error) {
			err = m.loadKeys(d.(*SystemUser))
			return
		},
	)

	return
}

// 创建
func (m *Model) Create(data *SystemUser) (err error) {
	od, _ := m.FindByPhone(data.Phone)
	if od.Id != "" {
		// 已经存在数据，直接返回
		err = e.AlreadyExists
		return
	}

	now := time.Now().Unix()
	data.CreateAt = now
	data.UpdateAt = now

	result, err := m.coll.InsertOne(context.Background(), data)
	if result.InsertedID != nil {
		data.Id = result.InsertedID.(primitive.ObjectID).Hex()
		m.clearKeys(data)
		return
	}
	return
}

// 更新产品
func (m *Model) Update(data *SystemUser) (err error) {
	mid, _ := primitive.ObjectIDFromHex(data.Id)
	data.UpdateAt = time.Now().Unix()
	data.Id = ""
	_, err = m.coll.UpdateByID(context.Background(), mid, bson.M{"$set": data})
	if err == nil {
		m.clearKeys(data)
	}
	return
}

// 删除
func (m *Model) Delete(id string) (err error) {
	data, err := m.FindOne(id)
	if err != nil {
		return
	}

	mid, _ := primitive.ObjectIDFromHex(id)
	_, err = m.coll.DeleteOne(context.Background(), bson.M{"_id": mid})
	if err == nil {
		m.clearKeys(data)
	}
	return
}
