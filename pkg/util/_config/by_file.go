package _config

import (
	"adminbg/pkg/util"
	"gopkg.in/yaml.v2"
	"os"
)

func MustLoadByFile(fpath string, conf interface{}) {
	f, err := os.Open(fpath)
	util.PanicIfErr(err, nil)

	err = yaml.NewDecoder(f).Decode(conf)
	util.PanicIfErr(err, nil)
}
