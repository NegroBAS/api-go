package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Claim la solictud para el jwt
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
