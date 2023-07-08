package middleware

import (
	"auto/app/svc"
	"auto/pkg/ctx"
	"auto/pkg/e"
	"auto/pkg/jwt"
	"auto/pkg/log"
	"time"

	"github.com/google/uuid"
)

const Token = "g-token"

func JWT(sv *svc.ServiceContext) ctx.HandlerFunc {
	conf := sv.Config
	return func(c *ctx.Context) {

		//

		token := c.Request.Header.Get(Token)
		if token == "" {
			c.Abort(e.TokenError.Add("token is empty"))
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			c.Abort(e.TokenError.Add(err.Error()))
			return
		}

		if time.Unix(claims.ExpiresAt, 0).Add(conf.Token.Expiration).Before(time.Now()) {
			c.Abort(e.TokenError.Add("token expired"))
			return
		}
		// ok set user id  and phone
		log.L.Debugf("user phone %s, userid: %d, level %s", claims.Phone, claims.Id, claims.Role)
		// todo 设置用户信息
		// c.Set("user", &api.User{Id: claims.Id, Phone: claims.Phone, Level: claims.Level})

		if claims.ExpiresAt < time.Now().Unix() {
			// 每次更新token
			LoadToken(c, claims, conf.Token.Expiration)
		}

		c.Set("request_id", uuid.New().String())

		c.Next()

	}
}

func LoadToken(c *ctx.Context, claims *jwt.Claims, expiration time.Duration) {
	claims.ExpiresAt = time.Now().Add(expiration).Unix()
	c.Header(Token, jwt.GenToken(*claims))
}
