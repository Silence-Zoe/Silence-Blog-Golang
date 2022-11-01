package global

import (
	"blog/config"
	"blog/ent"
)

var (
	DB     *ent.Client
	CONFIG config.Server
)
