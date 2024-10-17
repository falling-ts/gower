//go:build cli

package envs

import "os"

func init() {
	_ = os.Setenv("APP_MODE", "production")
}
