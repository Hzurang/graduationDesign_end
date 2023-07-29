package v1

import (
	"ginStudy/config"
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

/*
RefreshDbHandler
@author: LJR
@Description:
@param ctx
@Router /admin/v1/ljr/delete/empty/db [post]
*/
func RefreshDbHandler(ctx *gin.Context) {
	p := new(model.ParamRefreshDb)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("RefreshDb with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 参数缺漏
			utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, err)
			return
		}
		utils.ResponseParamTypeError(ctx, utils.FAIL_BUSINESS, nil, utils.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}
	if p.UserName == config.Config.RefreshDbConfig.UserName && p.Password == config.Config.RefreshDbConfig.Password {
		global.Db.Migrator().DropTable(&model.EssayWord{})
		global.Db.AutoMigrate(&model.EssayWord{})
		global.Db.Migrator().DropTable(&model.Essay{})
		global.Db.AutoMigrate(&model.Essay{})
		global.Db.Migrator().DropTable(&model.EssayCollect{})
		global.Db.AutoMigrate(&model.EssayCollect{})
		global.Db.Migrator().DropTable(&model.User{})
		global.Db.AutoMigrate(&model.User{})
		global.Db.Migrator().DropTable(&model.UserInfo{})
		global.Db.AutoMigrate(&model.UserInfo{})
		global.Db.Migrator().DropTable(&model.UserDate{})
		global.Db.AutoMigrate(&model.UserDate{})
		global.Db.Migrator().DropTable(&model.Sentence{})
		global.Db.AutoMigrate(&model.Sentence{})
		global.Db.Migrator().DropTable(&model.SentenceCollect{})
		global.Db.AutoMigrate(&model.SentenceCollect{})
		global.Db.Migrator().DropTable(&model.Listen{})
		global.Db.AutoMigrate(&model.Listen{})
		global.Db.Migrator().DropTable(&model.ListenCollect{})
		global.Db.AutoMigrate(&model.ListenCollect{})
		global.Db.Migrator().DropTable(&model.ListenWord{})
		global.Db.AutoMigrate(&model.ListenWord{})
		global.Db.Migrator().DropTable(&model.Admin{})
		global.Db.AutoMigrate(&model.Admin{})
		global.Db.Migrator().DropTable(&model.Word{})
		global.Db.AutoMigrate(&model.Word{})
		global.Db.Migrator().DropTable(&model.WordCollect{})
		global.Db.AutoMigrate(&model.WordCollect{})
		global.Db.Migrator().DropTable(&model.UserFeedback{})
		global.Db.AutoMigrate(&model.UserFeedback{})
		utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "操作成功")
		return
	}
	utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, "操作失败")
}
