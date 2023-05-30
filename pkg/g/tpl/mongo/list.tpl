

// 列表
func (m *{{.StructName}}Model) List(ctx context.Context, req *{{.ListStructName}}) (items []*{{.StructName}}, err error) {
	// Define the find options
	opt := options.Find()
	opt.SetSort(bson.D{{"{{"}}Key: "updateAt", Value: -1{{"}}"}})

	{{if .UsePage}}
	// 如果不是读取全部数据，就分页
	if req.Size != 0 {
	  opt.SetSkip((req.Page - 1) * req.Size)
    opt.SetLimit(req.Size)
  }
	{{end}}

	// Execute the find operation
	cursor, err := m.coll.Find(ctx, {{.ListFilter}}, opt)
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


func (m *Model) Count(ctx context.Context{{if .UseCountFilter}}, req *{{.ListStructName}}{{end}}) (int64, error) {
	return m.coll.CountDocuments(ctx, {{.ListFilter}})
}
