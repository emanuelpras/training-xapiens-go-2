package controller

import (
	"log"
	"os"
	"time"
	"xapiens-day9/models"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func (StrDB *StrDB) MiddleWare() (mw *jwt.GinJWTMiddleware) {

	if err := godotenv.Load(".env"); err != nil {
		log.Println("file .env tidak ditemukan")
	}
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(os.Getenv("SECRET_KEY")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.Users); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
					"role":      v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &login{
				Username: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var (
				loginVals login
				userDB    models.Users
			)

			// baca dari json di raw data
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			// check di database, apakah ada data yang dimaksud -> (inputan dari raw data)
			StrDB.DB.Where("username = ? ", loginVals.Username).First(&userDB)

			// compare password yang dimasukan di raw data apakah sama dengan yang ada di database
			if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(loginVals.Password)); err != nil {
				log.Println("Password tidak match")
			} else { // ketika passwordnya match maka hit code dibawahnya.
				return &userDB, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},

		// menentukan role nya
		Authorizator: func(data interface{}, c *gin.Context) bool {
			claims := jwt.ExtractClaims(c)
			method := c.Request.Method
			var result bool

			if claims["role"] == "admin" {
				result = true
			} else if claims["role"] == "guest" && method != "GET" {
				result = false
			} else {
				result = true
			}
			return result
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}) // close authMidlleware

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	// fmt.Println("ini param valuenya : ", mw)
	// fmt.Println("ini returnya : ", authMiddleware)
	return authMiddleware
}
