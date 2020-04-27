module github.com/lybeat/in-world-server

go 1.14

require (
	github.com/astaxie/beego v1.12.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.2
	github.com/go-ini/ini v1.55.0
	github.com/jinzhu/gorm v1.9.12
	github.com/unknwon/com v1.0.1
)

replace (
	github.com/lybeat/in-world-server/conf => ../in-world-server/conf
	github.com/lybeat/in-world-server/middleware => ../in-world-server/middleware
	github.com/lybeat/in-world-server/model => ../in-world-server/model
	github.com/lybeat/in-world-server/pkg/e => ../in-world-server/pkg/e
	github.com/lybeat/in-world-server/pkg/file => ../in-world-server/pkg/file
	github.com/lybeat/in-world-server/pkg/net => ../in-world-server/pkg/net
	github.com/lybeat/in-world-server/pkg/setting => ../in-world-server/pkg/setting
	github.com/lybeat/in-world-server/pkg/util => ../in-world-server/pkg/util
	github.com/lybeat/in-world-server/router => ../in-world-server/router
	github.com/lybeat/in-world-server/router/api => ../in-world-server/router/api
)
