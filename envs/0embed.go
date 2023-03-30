//go:build env || cli

package envs

import "embed"

//go:embed .env.development
//go:embed .env.test
//go:embed .env.production
var envs embed.FS

func init() {
	Envs = &envs
}
