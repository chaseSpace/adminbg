package pkg

import (
	"log"
)

/*
pkg目录：存放可以被重用的代码
-	不过请注意, 它区别于util目录，这个目录存放的是具有一定业务相关性的代码
-	可以包含子目录
*/

func PublicFunc() {
	log.Println("pkg/pkg_example.go: this is a publicFunc")
}
