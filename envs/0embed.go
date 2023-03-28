//go:build env

package envs

import "embed"

//go:embed .env.development
//go:embed .env.test
//go:embed .env.production
var fs embed.FS

func init() {
	FS = &fs
}
