package config

import (
	"errors"
	"ginStudy/utils"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"time"
)

var ErrorToken = errors.New(utils.GetCodeMsg(20005))

const (
	AcTokenExpiredDuration = 5 * time.Hour
	ReTokenExpiredDuration = 10 * 24 * time.Hour
)

type MyClaims struct {
	IP        string    `json:"ip"`
	UserId    uint64    `json:"user_id"`
	UserName  string    `json:"user_name"`
	Mobile    string    `json:"mobile"`
	Email     string    `json:"email"`
	LoginTime time.Time `json:"login_time"`
	Type      string    `json:"type"`
	jwt.RegisteredClaims
}

type MyReClaims struct {
	UserId uint64 `json:"user_id"`
	Type   string `json:"type"`
	jwt.RegisteredClaims
}

/*
GenToken
@author: LJR
@Description: 颁发token access token 和 refresh token
@param ip
@param userID
@param mobile
@param email
@return acToken
@return rfToken
@return err
*/
func GenToken(ip string, userID uint64, mobile string, email string, userName string) (acToken, rfToken string, err error) {
	var mySecret = []byte(Config.JWTConfig.Secret)
	myClaims := MyClaims{
		IP:        ip,
		UserId:    userID,
		Mobile:    mobile,
		Email:     email,
		UserName:  userName,
		LoginTime: time.Now(),
		Type:      "access",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    Config.JWTConfig.Issuer,
			Subject:   Config.JWTConfig.Subject,
			ExpiresAt: GetJWTTime(AcTokenExpiredDuration), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),     // 生效时间
		},
	}
	acToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims).SignedString(mySecret)
	if err != nil {
		zap.L().Error(utils.GetCodeMsg(20006), zap.Error(err))
		return "", "", errors.New(utils.GetCodeMsg(20006))
	}
	// refresh token 不需要保存任何用户信息
	rfClaims := MyReClaims{
		UserId: userID,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    Config.JWTConfig.Issuer,
			Subject:   Config.JWTConfig.Subject,
			ExpiresAt: GetJWTTime(ReTokenExpiredDuration), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),     // 生效时间
		},
	}
	rfToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, rfClaims).SignedString(mySecret)
	if err != nil {
		zap.L().Error(utils.GetCodeMsg(20007), zap.Error(err))
		return "", "", errors.New(utils.GetCodeMsg(20007))
	}
	return
}

/*
ParseToken
@author: LJR
@Description: 解析 JWT
@param token
@return *MyClaims
@return error
*/
func ParseToken(token string) (*MyClaims, error) {
	var mySecret = []byte(Config.JWTConfig.Secret)
	var mc = new(MyClaims)
	t, err := jwt.ParseWithClaims(token, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		zap.L().Error(utils.GetCodeMsg(20008), zap.Error(err))
		return nil, err
	}
	if !t.Valid {
		err = ErrorToken
		return nil, err
	}
	return mc, nil
}

/*
RefreshToken
@author: LJR
@Description: 通过 refresh token 刷新 access token
@param acToken
@param rfToken
@return newAcToken
@return newRfToken
@return err
*/
func RefreshToken(acToken, rfToken string) (newAcToken, newRfToken string, err error) {
	var mySecret = []byte(Config.JWTConfig.Secret)
	// rfToken 无效直接返回
	if _, err = jwt.Parse(rfToken, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	}); err != nil {
		return
	}
	// 从旧 access token 中解析出claims数据
	var mc MyClaims
	_, err = jwt.ParseWithClaims(acToken, &mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	// 判断错误是不是因为 access token 正常过期导致的
	if errors.Is(err, jwt.ErrTokenExpired) {
		return GenToken(mc.IP, mc.UserId, mc.Mobile, mc.Email, mc.UserName)
	}
	return
}

func GetJWTTime(t time.Duration) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(t))
}
