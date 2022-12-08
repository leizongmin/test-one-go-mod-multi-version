package core

import (
	"log"

	"github.com/leizongmin/test-one-go-mod-multi-version/common"
)

func Init() {
	log.Printf("common version: %s", common.Version())
}
