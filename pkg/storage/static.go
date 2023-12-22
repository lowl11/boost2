package storage

import "github.com/lowl11/boostef/ef"

func Init() {
	ef.Init(buildConnectionString())
}
