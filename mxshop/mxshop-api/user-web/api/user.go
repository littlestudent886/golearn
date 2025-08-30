package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
	"user-web/global/response"
	"user-web/proto"
)

func HandleGrpcErrorToHttp(err error, ctx *gin.Context) {
	// 将grpc的状态码转换成http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "服务不可用",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误",
				})
			}
			return
		}
	}

}

func GetUserList(ctx *gin.Context) {
	ip := "127.0.0.1"
	port := 50051

	// 连接用户grpc服务
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetUserList]连接用户服务失败",
			"msg", err.Error(),
		)
	}

	// 调用接口
	userClient := proto.NewUserClient(conn)
	rsp, err := userClient.GetUserList(ctx, &proto.PageInfo{
		Pn:    1,
		PSize: 10,
	})

	if err != nil {
		zap.S().Errorw("[GetUserList]获取用户列表失败",
			"msg", err.Error(),
		)
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		//data := make(map[string]interface{})
		//data["id"] = value.Id
		//data["mobile"] = value.Mobile
		//data["nick_name"] = value.NickName
		//data["gender"] = value.Gender
		//data["birthday"] = value.Birthday
		user := response.UserResponse{
			ID:       value.Id,
			Mobile:   value.Mobile,
			NickName: value.NickName,
			Gender:   value.Gender,
			//Birthday: time.Unix(int64(value.Birthday), 0).Format("2006-01-02"),
			Birthday: response.JsonTime(time.Unix(int64(value.Birthday), 0)),
		}
		result = append(result, user)
	}

	ctx.JSON(http.StatusOK, result)
}
