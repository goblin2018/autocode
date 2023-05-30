package {{.package}}

import (
	"fmt"
	"context"
	"auto/pkg/mgo"
	"auto/pkg/e"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"


)

type Model struct {
	name      string
	primaryKey string
	secondKey string
	coll      *mongo.Collection
	*mgo.MongoModel
}

func NewModel(md *mgo.MongoModel) *Model {
	name := "{{.package}}"
	// 创建collection
	

	m := &Model{
		name:      name,
		MongoModel: md,
		coll:      md.InitCollection(name),
		primaryKey: fmt.Sprintf("cache:%s:id:", name),
		secondKey: fmt.Sprintf("cache:%s{{.UniKeysWithColon}}:", name),
	}
	// 创建唯一索引
	m.AddUniKey(m.coll{{.UniKeysWithComma}})

	return m
}

// 加载key
func (m *Model) loadKeys(data *{{.StructName}}) (err error) {
	if err = m.Set(m.GetKey(m.primaryKey, data.Id), data); err != nil {
		return
	}
	if err = m.Set(m.GetKey(m.secondKey, {{.UniKeysWithDataComma}}), data.Id); err != nil {
		return
	}
	return
}

// 清空Key
func (m *Model) clearKeys(data *{{.StructName}}) (err error) {
	if err = m.Del(m.GetKey(m.primaryKey, data.Id)); err != nil {
		return
	}
	if err = m.Del(m.GetKey(m.secondKey, {{.UniKeysWithDataComma}})); err !=nil {
		return
	}
	return
}


// 获取
func (m *Model) FindOne(id string) (data *{{.StructName}},err error) {
	data = &{{.StructName}}{}
	key := m.GetKey(m.primaryKey, id)
	
	err = m.Take(
		key,
		data,
		// 查询
		func(d interface{}) (err error) {
			mid,err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return
			}
			err = m.coll.FindOne(context.Background(), bson.M{"_id": mid}).Decode(d)
			return
		},
		// 加载
		func(d interface{}) (err error) {
			err = m.loadKeys(d.(*{{.StructName}}))
			return
		},
	)

	return
}


// 通过唯一key查找
func (m *Model) FindBy{{.UniKeysWithAnd}}({{.UniKeysWithType}}) (data *{{.StructName}}, err error) {
	data = &{{.StructName}}{}

	key := m.GetKey(m.secondKey, {{.UniKeysWithoutType}})
	
	err = m.Take(
		key,
		data,
		// 查询
		func(d interface{}) (err error) {
			var id string
			if err = m.Get(key, &id); err != nil {
				d, err = m.FindOne(id)
				return
			}
			err = m.coll.FindOne(context.Background(), bson.M{ {{.UniKeysBsonFilter}} }).Decode(d)
			return
		},
		// 加载
		func(d interface{}) (err error) {
			err = m.loadKeys(d.(*{{.StructName}}))
			return
		},
	)

	return
}

// 创建
func (m *Model) Create(data *{{.StructName}}) (err error) {
	od, _ := m.FindBy{{.UniKeysWithAnd}}({{.UniKeysWithDataComma}})
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
func (m *Model) Update(data *{{.StructName}}) (err error) {
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

