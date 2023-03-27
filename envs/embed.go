//go:build env

package envs

import "embed"

//go:embed .env
//go:embed .env.example
var fs embed.FS

func init() {
	FS = &fs
}
