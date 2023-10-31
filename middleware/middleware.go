package middleware

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/dgrijalva/jwt-go"
    "strings"
)

type CustomClaims struct {
    UserID uint
    Role   int
    jwt.StandardClaims
}

var secretKey = []byte("your-secret-key")

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        tokenString := c.Request().Header.Get("Authorization")
        if tokenString == "" {
            return c.String(http.StatusUnauthorized, "Unauthorized")
        }

        tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrSignatureInvalid
            }
            return secretKey, nil
        })

        if err != nil || !token.Valid {
            return c.String(http.StatusUnauthorized, "Unauthorized")
        }

        claims, ok := token.Claims.(*CustomClaims)
        if !ok || claims.Role != 1 {
            return c.String(http.StatusUnauthorized, "Unauthorized")
        }

        c.Set("user", claims)

        return next(c)
    }
}

func SuperAdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        tokenString := c.Request().Header.Get("Authorization")
        if tokenString == "" {
            return c.String(http.StatusUnauthorized, "Unauthorized")
        }

        tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrSignatureInvalid
            }
            return secretKey, nil
        })

        if err != nil || !token.Valid {
            return c.String(http.StatusUnauthorized, "Unauthorized")
        }

        claims, ok := token.Claims.(*CustomClaims)
        if !ok || claims.Role != 2 {
            return c.String(http.StatusUnauthorized, "Unauthorized")
        }

        c.Set("user", claims)

        return next(c)
    }
}