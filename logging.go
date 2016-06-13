package plalib

import (
	"github.com/op/go-logging"
)

var log *logging.Logger

func RegisterLogger(logger *logging.Logger) {
	log = logger
}
