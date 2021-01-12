package _config

import (
	"adminbg/util"
	"gopkg.in/yaml.v2"
	"os"
)

func MustLoadByFile(path string, conf interface{}) {
	f, err := os.Open(path)
	util.PanicIfErr(err, nil)

	err = yaml.NewDecoder(f).Decode(conf)
	util.PanicIfErr(err, nil)
}
