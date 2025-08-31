package api

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"strings"
	"time"
	"user-web/forms"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"user-web/global"
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

func HandleValidatorError(ctx *gin.Context, err error) {
	//如何返回错误信息
	errors, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errors.Translate(global.Trans)),
	})

}

func GetUserList(ctx *gin.Context) {
	ip := global.ServerConfig.UserSrvInfo.Host
	port := global.ServerConfig.UserSrvInfo.Port

	// 连接用户grpc服务
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetUserList]连接用户服务失败",
			"msg", err.Error(),
		)
	}

	// 调用接口
	userClient := proto.NewUserClient(conn)

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "0")
	pSizeInt, _ := strconv.Atoi(pSize)

	rsp, err := userClient.GetUserList(ctx, &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
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

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func PassWordLogin(ctx *gin.Context) {
	// 表单验证
	passWordLoginForm := forms.PassWordLoginForm{}
	if err := ctx.ShouldBind(&passWordLoginForm); err != nil {
		//如何返回错误信息
		HandleValidatorError(ctx, err)
	}
}
