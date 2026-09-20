package main

import "embed"

//go:embed all:frontend/dist
var Assets embed.FS

//go:embed bin/weed
var SeaweedFS []byte
