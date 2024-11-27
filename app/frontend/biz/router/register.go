// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	auth "gomall/app/frontend/biz/router/auth"
	category "gomall/app/frontend/biz/router/category"
	home "gomall/app/frontend/biz/router/home"
	product "gomall/app/frontend/biz/router/product"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	product.Register(r)

	category.Register(r)

	auth.Register(r)

	home.Register(r)
}
