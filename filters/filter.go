package filters

import (
	"fmt"
	"github.com/astaxie/beego/context"
)

func FilterUser(ctx *context.Context) {
	fmt.Println(ctx.Input.RequestBody)
}
