/*
@Time : 2020/11/3 下午9:49
@Author : hoastar
@File : auth
@Software: GoLand
*/

package middlerware

import (
	jwt "github.com/hoastar/orange/pkg/jwtauth"
	"github.com/hoastar/orange/tools/config"
	"time"
)

func AuthInit() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm: "test zone",
		Key: []byte(config.ApplicationConfig.JwtSecret),
		Timeout: time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: handler.PayloadFunc,
		IdentityHandler: handler.IdentityHandler,
		Authenticator: handler.Authenticator,
		Authorizator: handler.Authorizator,
		Unauthorized: handler.Unauthorized,
		TokenLookup: "header: Authorizaiton, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})
}
