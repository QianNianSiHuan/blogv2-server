package jwts

import (
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

// JwtPayLoad jwt中payload数据
type JwtPayLoad struct {
	UserID   uint          `json:"user_id"`
	Username string        `json:"username"` // 用户名
	Role     enum.RoleType `json:"role"`     // 权限  1 管理员  2 用户 3 游客
}

type CustomClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}

func (m CustomClaims) GetUser() (user models.UserModel, err error) {
	err = global.DB.Take(&user, m.UserID).Error
	return
}

// GenToken 创建 Token
func GenToken(user JwtPayLoad) (string, error) {
	claim := CustomClaims{
		JwtPayLoad: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expire))),
			Issuer:    global.Config.Jwt.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(global.Config.Jwt.Secret))
}

// ParseToken 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	if tokenStr == "" {
		return nil, errors.New("token不能为空")
	}
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.Secret), nil
	})
	if err != nil {
		fmt.Println("token err : ", err)
		if strings.Contains(err.Error(), "token is expired") {
			return nil, errors.New("token过期")
		}
		if strings.Contains(err.Error(), "token is malformed") {
			return nil, errors.New("token格式错误")
		}
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
func ParseTokenByGin(c *gin.Context) (*CustomClaims, error) {
	var err error
	token := c.GetHeader("token")
	if token == "" {
		token = c.Query("token")
		if token == "" {
			token, err = c.Cookie("token")
			if err != nil || token == "" {
				return nil, errors.New("token不存在")
			}
		}
	}
	return ParseToken(token)
}

func GetClaims(c *gin.Context) (claims *CustomClaims) {
	_claims, ok := c.Get("claims")
	if !ok {
		return
	}
	claims, ok = _claims.(*CustomClaims)
	if !ok {
		return
	}
	return
}
