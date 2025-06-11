package middleware

import (
	"net/http"
	"school-backend/pkg/jwt"
	"school-backend/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

// Middleware JWT + validasi IP & User-Agent
func AuthRequired() gin.HandlerFunc {
	jwtManager := jwt.NewManager()

	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.Error(c, http.StatusUnauthorized, "Token tidak ditemukan", nil)
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwtManager.ParseToken(tokenStr)
		//fmt.Println("[DEBUG] claims.Peran:", claims.Peran)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Token tidak valid", err.Error())
			c.Abort()
			return
		}

		requestIP := c.ClientIP()
		requestUA := c.GetHeader("User-Agent")

		if claims.IP != "" && claims.IP != requestIP {
			response.Error(c, http.StatusUnauthorized, "IP address tidak sesuai", gin.H{
				"expected": claims.IP,
				"got":      requestIP,
			})
			c.Abort()
			return
		}

		if claims.UserAgent != "" && claims.UserAgent != requestUA {
			response.Error(c, http.StatusUnauthorized, "Perangkat tidak sesuai", gin.H{
				"expected": claims.UserAgent,
				"got":      requestUA,
			})
			c.Abort()
			return
		}

		c.Set("pengguna_id", claims.PenggunaID)
		c.Set("peran_id_str", claims.Peran)
		c.Next()
	}
}

// Middleware untuk validasi role tertentu
func RequireRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get("peran_id_str")
		role, ok := roleVal.(string)

		//fmt.Println("[DEBUG] Context role:", role)
		if !exists || !ok {
			response.Error(c, http.StatusForbidden, "Peran tidak ditemukan atau tidak valid", nil)
			c.Abort()
			return
		}

		for _, allowed := range allowedRoles {
			if role == allowed {
				c.Next()
				return
			}
		}

		response.Error(c, http.StatusForbidden, "Akses ditolak: role tidak diizinkan", nil)
		c.Abort()
	}
}
