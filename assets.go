//go:build !dev

package evorsio

import "embed"

const ViteURL = ""

//go:embed all:web/dist
var Assets embed.FS
