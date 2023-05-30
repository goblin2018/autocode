package house

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 列表
func (m *Model) List(ctx context.Context, req *ListHouseReq) (items []*House, err error) {
	// Define the find options
	opt := options.Find()
	opt.SetSort(bson.D{{Key: "updateAt", Value: -1}})

	// 如果不是读取全部数据，就分页
	if req.Size != 0 {
		opt.SetSkip((req.Page - 1) * req.Size)
		opt.SetLimit(req.Size)
	}

	// Execute the find operation
	cursor, err := m.coll.Find(ctx, m.genFilter(req), opt)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)

	// Decode the results into a slice
	if err = cursor.All(ctx, &items); err != nil {
		return
	}

	return
}

func (m *Model) Count(ctx context.Context, req *ListHouseReq) (int64, error) {
	return m.coll.CountDocuments(ctx, m.genFilter(req))
}

func (m *Model) genFilter(req *ListHouseReq) (filter bson.M) {
	filter = bson.M{}
	if req.OrgId != "" {
		filter["orgId"] = req.OrgId
	}

	return
}
