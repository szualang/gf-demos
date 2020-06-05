package boot

import (
	"github.com/gogf/gf-swagger/swagger"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	gf_gdb_driver_taosSql "github.com/szualang/gf-gdb-driver-taosSql"
)

// 用于应用初始化。
func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	if err := gdb.Register("taosSql", &gf_gdb_driver_taosSql.DriverTaosSQL{}); err != nil {
		panic(err)
	}
}
