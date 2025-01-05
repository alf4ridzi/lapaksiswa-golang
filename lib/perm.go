package lib

import (
	"slices"

	"github.com/alf4ridzi/lapaksiswa-golang/config"
)

func IsExtAllowed(Ext string) bool {
	extAllowed := config.EXT_ALLOWED_IMAGE
	return slices.Contains(extAllowed, Ext)
}
