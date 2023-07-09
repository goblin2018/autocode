package g

func GenModels(baseDir string, name string, model *M) {
	if model.DB == "mongo" {
		GenMongo(baseDir, name, model)
	} else if model.DB == "mysql" {
		GenMysql(baseDir, name, model)
	}
}
