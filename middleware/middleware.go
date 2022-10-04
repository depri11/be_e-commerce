package middleware

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/depri11/be_e-commerce/helpers"
	"github.com/depri11/be_e-commerce/models"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kataras/iris/v12"
)

type middleware struct {
	Redis  *redis.Client
	JwtKey string
}

func NewMiddleware(redis *redis.Client, JwtKey string) *middleware {
	return &middleware{
		Redis:  redis,
		JwtKey: JwtKey,
	}
}

func (m *middleware) CORS() iris.Handler {
	return func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", ctx.GetHeader("Origin"))
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")

		if ctx.Request().Method == "OPTIONS" {
			ctx.StatusCode(204)
			return
		}

		ctx.Next()
	}
}

func (m *middleware) GetUser(ctx iris.Context) map[string]interface{} {
	user := ctx.Values().Get("User").(jwt.MapClaims)
	return user
}

func (m *middleware) Customize() iris.Handler {
	return func(ctx iris.Context) {
		log.Println(ctx.Request().Host, ctx.Request().URL.Path)
		header := ctx.Request().Header.Get("Authorization")

		if !strings.Contains(header, "Bearer ") {
			m.response(ctx, 400, "Bad Request Jwt", nil)
			ctx.StatusCode(401)
			return
		}

		var token string
		arrToken := strings.Split(header, " ")
		if len(arrToken) == 2 {
			token = arrToken[1]
		}

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.JwtKey), nil
		})

		if claims["exp"] == nil {
			m.response(ctx, 401, "invalid token expires", nil)
			return
		}

		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			m.response(ctx, 401, "token expired", nil)
			return
		}

		sessionID := claims["session_id"].(string)
		userJson, err := m.Redis.Get(sessionID).Result()
		if err != nil {
			log.Println(err)
			m.response(ctx, 401, "unknown token", nil)
			return
		}

		user := models.User{}
		if claims["type"] == nil {
			if err := json.Unmarshal([]byte(userJson), &user); err != nil {
				log.Println(err)
				m.response(ctx, 401, "token error", nil)
				return
			}
		}

		ctx.Values().Set("User", user)
		ctx.Values().Set("user_id", user.ID)
		ctx.Values().Set("user_role", user.Roles)
		ctx.Values().Set("user_username", user.Username)
		ctx.Values().Set("session_id", sessionID)
	}
}

func (m *middleware) response(ctx iris.Context, status int, message string, data interface{}) {
	// ctx.StatusCode(status)
	ctx.JSON(helpers.Response{Status: status, Message: message, Data: data})
}
