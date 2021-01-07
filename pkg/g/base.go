package g

/*
base.go 管理 ·g· 下面所有对象的init和stop
*/
func MustInit() {
	mustInitDB()
}

func Stop() {
	if Mysql != nil {
		db, _ := Mysql.DB()
		_ = db.Close()
	}
}
