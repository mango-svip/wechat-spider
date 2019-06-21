package jwt

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var jwtSecret = []byte("xxxxx")

type Claims struct {
    Username string `json:"username"`
    Password string `json:"password"`
    jwt.StandardClaims
}


func GenerateToken(username string , password string) (string, error) {
    now := time.Now()
    expire := now.Add(3 * time.Hour)

    standardClaims := jwt.StandardClaims{
        ExpiresAt: expire.Unix(),
        Issuer:    "spider",
    }

    claims := Claims{
        Username: username,
        Password: password,
        StandardClaims: standardClaims,

    }
    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    s, e := tokenClaims.SignedString(jwtSecret)
    return s,e
}

func ParseToken(token string) (*Claims, error) {
    tokenClaims, e := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if tokenClaims != nil {
        if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
            return claims, nil
        }
    }
    return nil, e
}
