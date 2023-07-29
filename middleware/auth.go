package middleware

import (
	"ginStudy/config"
	"ginStudy/global"
	"ginStudy/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

const (
	AuthorizationKey = "Authorization"
	GrantTypeKey     = "Grant_type"
	SpaceKey         = " "
	BearerKey        = "Bearer"
)

/*
JWTAuthMiddleware
@author: LJR
@Description: JWT 中间件
@return func(ctx *gin.Context)
*/
func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		rtoken := ctx.Request.Header.Get(GrantTypeKey)
		zap.L().Info("rtoken来了 " + rtoken)
		authHeader := ctx.Request.Header.Get(AuthorizationKey)
		if authHeader == "" {
			utils.ResponseError(ctx, utils.FAIL_TOKEN, nil, utils.GetCodeMsg(20001))
			ctx.Abort()
			return
		}
		zap.L().Info("客户端ip:" + ctx.ClientIP() + " token:" + authHeader)

		// 按空格分隔
		parts := strings.SplitN(authHeader, SpaceKey, 2)
		if !(len(parts) == 2) && parts[0] == BearerKey {
			utils.ResponseError(ctx, utils.FAIL_TOKEN, nil, utils.GetCodeMsg(20003))
			ctx.Abort()
			return
		}

		// part[1]是获取到的 access token
		mc, err := config.ParseToken(parts[1])
		_, err1 := global.RD.ZRank("JWT_AUTH_USER:Baned", parts[1]).Result()
		if err1 == nil {
			utils.ResponseError(ctx, utils.FAIL_TOKEN, nil, utils.GetCodeMsg(20009))
			ctx.Abort()
			return
		}

		if err != nil {
			newAcToken, newReToken, err := config.RefreshToken(parts[1], rtoken)
			if err != nil {
				utils.ResponseError(ctx, utils.FAIL_TOKEN, nil, utils.GetCodeMsg(20004))
				ctx.Abort()
				return
			}
			ctx.Request.Header.Set("newAcToken", newAcToken)
			zap.L().Info("新的AcToken来了 " + newAcToken)
			ctx.Request.Header.Set("newReToken", newReToken)
			zap.L().Info("新的ReToken来了 " + newReToken)
			ctx.Next()
		}
		if mc.Type != "access" {
			utils.ResponseError(ctx, utils.FAIL_TOKEN, nil, utils.GetCodeMsg(20003))
			ctx.Abort()
			return
		}
		ctx.Set("token", parts[1])
		ctx.Set("myClaims", mc)
		ctx.Set("user_id", mc.UserId)
		ctx.Next()
	}
}
