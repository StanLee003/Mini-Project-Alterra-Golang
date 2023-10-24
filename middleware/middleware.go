package middleware

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/dgrijalva/jwt-go"
    "bikrent/models"
    "gorm.io/gorm"
    "strings"
)

type CustomClaims struct {
    UserID uint
    Role   int
    jwt.StandardClaims
}

var JWTMiddleware echo.MiddlewareFunc

func InitJWTMiddleware(db *gorm.DB) echo.MiddlewareFunc {
    JWTMiddleware = createJWTMiddleware(db)
    return JWTMiddleware
}

func createJWTMiddleware(db *gorm.DB) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            authorizationHeader := c.Request().Header.Get("Authorization")

            if !strings.HasPrefix(authorizationHeader, "Bearer ") {
                return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing or invalid JWT token"})
            }

            token := strings.TrimPrefix(authorizationHeader, "Bearer ")

            claims := &CustomClaims{}
            t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
                return []byte("your-secret-key"), nil
            })

            if err != nil {
                return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid JWT token"})
            }

            if !t.Valid {
                return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid JWT token"})
            }

            var user models.User
            if err := db.Where("id = ?", claims.UserID).First(&user).Error; err != nil {
                return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User not found"})
            }

            if user.Role != claims.Role {
                return c.JSON(http.StatusForbidden, map[string]string{"error": "Permission denied"})
            }

            c.Set("userID", claims.UserID)
            c.Set("userRole", claims.Role)

            return next(c)
        }
    }
}

