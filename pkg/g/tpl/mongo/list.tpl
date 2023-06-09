package {{.package}}

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)


// 列表
func (m *Model) {{.funcName}}(ctx context.Context, req *{{.reqName}}) (items []*{{.StructName}}, err error) {
	// Define the find options
	opt := options.Find()
	opt.SetSort(bson.D{{"{{"}}Key: "updateAt", Value: -1{{"}}"}})

	{{if .hasPage}}
	// 如果不是读取全部数据，就分页
	if req.Size != 0 {
	  opt.SetSkip((req.Page - 1) * req.Size)
    opt.SetLimit(req.Size)
  }
	{{end}}

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


func (m *Model) Count(ctx context.Context, req *{{.reqName}}) (int64, error) {
	return m.coll.CountDocuments(ctx, m.genFilter(req))
}


func (m *Model) genFilter(req *{{.reqName}}) (filter bson.M) {
	filter = bson.M{}
	{{.filter}}
	return
}
