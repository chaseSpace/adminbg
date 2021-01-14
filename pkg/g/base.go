package g

/*
base.go Usage:  Maintain all global objects init and stop operations under the pkg `g`.
*/
func MustInit() {
	initDB()
}

func Stop() {
	if MySQL != nil {
		db, _ := MySQL.DB()
		_ = db.Close()
	}
}
