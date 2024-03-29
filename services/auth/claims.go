package auth

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UpdateDur *jwt.NumericDate `json:"upd,omitempty"` // Token 更新时限
	jwt.RegisteredClaims
}

// Set 设置数据
func (c *Claims) Set(args ...any) *Claims {
	for _, arg := range args {
		c.decideType(arg)
	}
	return c
}

func (c *Claims) decideType(arg any) {
	switch arg.(type) {
	case string:
		str := arg.(string)
		if c.Issuer == "" {
			c.Issuer = str
			break
		}
		if c.Subject == "" {
			c.Subject = str
			break
		}
		if c.ID == "" {
			c.ID = str
		}
	case jwt.ClaimStrings:
		c.Audience = arg.(jwt.ClaimStrings)
	case []string:
		c.Audience = arg.([]string)
	case *jwt.NumericDate:
		time := arg.(*jwt.NumericDate)
		if c.UpdateDur == nil {
			c.UpdateDur = time
			break
		}
		if c.ExpiresAt == nil {
			c.ExpiresAt = time
			break
		}
		if c.NotBefore == nil {
			c.NotBefore = time
			break
		}
		if c.IssuedAt == nil {
			c.IssuedAt = time
		}
	}
}
