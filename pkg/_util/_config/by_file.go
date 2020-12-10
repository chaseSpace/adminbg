package _config

import (
	"adminbg/pkg/_util"
	"gopkg.in/yaml.v2"
	"os"
)

func MustLoadByFile(fpath string, conf interface{}) {
	f, err := os.Open(fpath)
	_util.PanicIfErr(err, nil)

	err = yaml.NewDecoder(f).Decode(conf)
	_util.PanicIfErr(err, nil)
}
