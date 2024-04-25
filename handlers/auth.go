package handlers

import (
	"crypto/rand"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	secretKey = []byte("your-secret-key")
	tokens    = make(map[string]time.Time)
	mu        sync.Mutex
)

// CustomClaims 自定义声明结构体，用于生成 Token
type CustomClaims struct {
	Username string `json:"username"`
	UserSalt string `json:"userSalt"`
	jwt.StandardClaims
}

// 生成指定长度的随机字符串
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// charsetLen := len(charset)

	// 为了避免 modulo bias，我们将随机值映射到 charset 的长度范围内
	charsetLen := big.NewInt(int64(len(charset)))

	randomString := make([]byte, length)
	for i := range randomString {
		randomIndex, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			panic(err)
		}
		randomString[i] = charset[randomIndex.Int64()]
	}

	return string(randomString)
}

// 生成 Token
func generateToken(username string) (string, error) {
	// 设置 Token 的过期时间为 24 小时
	expirationTime := time.Now().Add(12 * time.Hour)
	// 随机生成8位作为盐,区分不同会话
	userSalt := generateRandomString(8)
	claims := CustomClaims{
		Username: username,
		UserSalt: userSalt,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "your-website",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	// 将 Token 存储起来
	mu.Lock()
	defer mu.Unlock()
	tokens[tokenString] = expirationTime

	return tokenString, nil
}

// AuthMiddleware Token 验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		// 去除 Bearer 前缀
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		fmt.Println(tokenString)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "not login"})
			c.Abort()
			return
		}

		// 解析 Token
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		// Token 解析失败或过期
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is invalid"})
			c.Abort()
			return
		}

		// 检查 Token 是否过期
		expirationTime, ok := tokens[tokenString]
		if !ok || time.Now().After(expirationTime) {
			delete(tokens, tokenString) // 删除过期的 Token
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}
		// 如果快要过期了，则重新生成一个token替代原本的token todo

		// 获取到token中的内容
		claims, _ := token.Claims.(*CustomClaims)
		c.Set("username", claims.Username)
		c.Next()
	}
}
