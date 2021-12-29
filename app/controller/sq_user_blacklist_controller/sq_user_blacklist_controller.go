package sq_user_blacklist_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jtzjtz/kit/convert"
	ctext "github.com/jtzjtz/kit/ctx"
	"github.com/jtzjtz/ys_api/app"
	gRPCPool "github.com/jtzjtz/ys_api/app/conn/grpc"
	"github.com/jtzjtz/ys_pack/client/grpc/sq_user_blacklist_grpc"
	"github.com/jtzjtz/ys_pack/entity"
	"github.com/jtzjtz/ys_pack/proto/sq_user_blacklist_proto"
	"net/http"
	"time"
)

func CreateUserBlacklist(ctx *gin.Context) {
	var result app.Result

	service := &sq_user_blacklist_grpc.SqUserBlacklistService{}
	conn, err := gRPCPool.FrameworkGRPC.Get()
	if err != nil {
		result.SetCode(app.CODE_ERROR).SetMessage("参数错误")
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer gRPCPool.FrameworkGRPC.Dispose(conn)
	grpcCtx, _ := ctext.CreateContext()
	userData := &sq_user_blacklist_proto.SqUserBlacklist{}
	userData.Userid = 1
	userData.Username = "zhansang"
	userData.Createtime = time.Now().Format("2006-01-02 15:04")
	resultCreate := service.CreateSqUserBlacklist(grpcCtx, userData, conn)
	ctx.JSON(200, resultCreate)

}
func CreateUserWithBind(ctx *gin.Context) {
	var result app.Result

	userRequest := entity.SqUserBlacklist{}
	ctx.ShouldBind(&userRequest)
	service := &sq_user_blacklist_grpc.SqUserBlacklistService{}

	conn, err := gRPCPool.FrameworkGRPC.Get()
	if err != nil {
		result.SetCode(app.CODE_ERROR).SetMessage("参数错误")
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer gRPCPool.FrameworkGRPC.Dispose(conn)

	grpcCtx, _ := ctext.CreateContextWithTime(time.Second * 60)
	userData := &sq_user_blacklist_proto.SqUserBlacklist{}
	convert.EntityToEntity(userRequest, &userData)
	resultCreate := service.CreateSqUserBlacklist(grpcCtx, userData, conn)
	ctx.JSON(200, resultCreate)

}

//
//func UpdateUser(ctx *gin.Context) {
//	var result app.Result
//
//	service := &user_service.UserService{}
//
//	conn, err := gRPCPool.FrameworkGRPC.Get()
//	if err != nil {
//		result.SetCode(app.CODE_ERROR).SetMessage("参数错误")
//		ctx.JSON(http.StatusOK, result)
//		return
//	}
//	defer gRPCPool.FrameworkGRPC.Dispose(conn)
//
//	userEntity := entity.User{}
//	ctx.ShouldBind(&userEntity)
//
//	userData := user_proto.User{}
//	util.EntityToEntity(userEntity, &userData)
//	//userData.Email = "jtzjtz@163.com"
//	//userData.Fullname = "测试"
//	//userData.Mastername = "test"
//	updateCondition := user_proto.UpdateAndCondition{}
//	updateCondition.Query = &user_proto.Query{} //此处一定先给Query 初始化
//	//更新条件
//	//updateCondition.Query.SqlQuery = "masterid=1 and employeeid='1' " //sql 语句的where条件
//	updateCondition.Query.EntityQuery = &user_proto.User{Masterid: 5, Mastername: "test"}
//	//更新的数据
//	updateCondition.Entity = &userData
//	//如果更新的字段是该数据类型的默认值，请在此处填写
//	updateCondition.UpdateEmptyFields = []string{"Status", "Deptname"}
//	grpcCtx, _ := util.CreateContextWithTime(time.Second * 60)
//
//	resultCreate := service.UpdateUser(grpcCtx, &updateCondition, conn)
//	ctx.JSON(200, resultCreate)
//}
//func DelUser(ctx *gin.Context) {
//	var result app.Result
//
//	service := &user_service.UserService{}
//
//	conn, err := gRPCPool.FrameworkGRPC.Get()
//	if err != nil {
//		result.SetCode(app.CODE_ERROR).SetMessage("参数错误")
//		ctx.JSON(http.StatusOK, result)
//		return
//	}
//	defer gRPCPool.FrameworkGRPC.Dispose(conn)
//
//	grpcCtx, _ := util.CreateContextWithTime(time.Second * 60)
//	query := user_proto.Query{}
//	//删除条件
//	//query.SqlQuery = "masterid=1 and employeeid='1' " //sql 语句的where条件
//	query.EntityQuery = &user_proto.User{Masterid: 6, Mastername: "test"}
//	resultCreate := service.DeleteUser(grpcCtx, &query, conn)
//	ctx.JSON(200, resultCreate)
//}
func GetUserBlackList(ctx *gin.Context) {
	var result app.Result

	service := &sq_user_blacklist_grpc.SqUserBlacklistService{}

	conn, err := gRPCPool.FrameworkGRPC.Get()
	if err != nil {
		result.SetCode(app.CODE_ERROR).SetMessage("参数错误")
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer gRPCPool.FrameworkGRPC.Dispose(conn)

	grpcCtx, _ := ctext.CreateContextWithTime(time.Second * 60)

	query := sq_user_blacklist_proto.Query{}
	//删除条件
	//query.SqlQuery = "masterid=1 and employeeid='1' " //sql 语句的where条件
	query.EntityQuery = &sq_user_blacklist_proto.SqUserBlacklist{}
	query.Db = sq_user_blacklist_proto.DB_WRITE //强制主库查询 默认从库
	//排序
	query.OrderBy = map[string]string{"id": "desc"}
	resultCreate := service.GetSqUserBlacklistList(grpcCtx, &query, conn)
	ctx.JSON(200, resultCreate)
}

//
//func GetUserPageList(ctx *gin.Context) {
//	var result app.Result
//
//	var pageEntity entity2.PageForm
//	ctx.ShouldBind(&pageEntity)
//	service := &user_service.UserService{}
//
//	conn, err := gRPCPool.FrameworkGRPC.Get()
//	if err != nil {
//		result.SetCode(app.CODE_ERROR).SetMessage("参数错误")
//		ctx.JSON(http.StatusOK, result)
//		return
//	}
//	defer gRPCPool.FrameworkGRPC.Dispose(conn)
//
//	grpcCtx, _ := util.CreateContextWithTime(time.Second * 60)
//
//	pageQuery := user_proto.PageQuery{}
//
//	query := user_proto.Query{}
//	pageQuery.Query = &query
//	//删除条件
//	//query.SqlQuery = "masterid=1 and employeeid='1' " //sql 语句的where条件
//	query.EntityQuery = &user_proto.User{}
//	query.Db = user_proto.DB_WRITE //强制主库查询 默认从库
//	//排序
//	query.OrderBy = map[string]string{"masterid": "desc", "updatetime": "asc"}
//	pageQuery.Page = pageEntity.Page
//	pageQuery.PageNum = pageEntity.PageNum
//	resultCreate := service.GetUserPageList(grpcCtx, &pageQuery, conn)
//	ctx.JSON(200, resultCreate)
//}
//func GetUser(ctx *gin.Context) {
//	var result app.Result
//
//	service := &user_service.UserService{}
//	conn, err := gRPCPool.FrameworkGRPC.Get()
//	if err != nil {
//		result.SetCode(app.CODE_ERROR).SetMessage("参数错误")
//		ctx.JSON(http.StatusOK, result)
//		return
//	}
//	defer gRPCPool.FrameworkGRPC.Dispose(conn)
//
//	grpcCtx, _ := util.CreateContextWithTime(time.Second * 60)
//
//	query := user_proto.Query{}
//	//删除条件
//	//query.SqlQuery = "masterid=1 and employeeid='1' " //sql 语句的where条件
//	query.EntityQuery = &user_proto.User{Mastername: "test"}
//	//query.Db = user_proto.DB_WRITE //强制主库查询 默认从库
//	//排序
//	query.OrderBy = map[string]string{"masterid": "desc", "updatetime": "asc"}
//	resultCreate := service.GetUser(grpcCtx, &query, conn)
//	ctx.JSON(200, resultCreate)
//}
//
//func GetUserStatus(ctx *gin.Context) {
//	pool := redisPool.FrameworkRedis
//	redis := pool.Get()
//
//	r, err := redis.Do("ping")
//	if err != nil {
//		log.Fatal("[ERROR] ping redis fail", err)
//	}
//
//	ctx.JSON(200, r)
//}
