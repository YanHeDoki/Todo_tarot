package tarots

import "embed"

//go:embed tarotCards
var TarotCards embed.FS

//go:embed json
var TarotJson embed.FS

//go:embed fontLibrary
var Font embed.FS
