package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/veerlakshya/user-service/consts"
	svc "github.com/veerlakshya/user-service/internal/application/user"
)

type userHandler struct {
	userService svc.UserService
}

func NewUserHandler(r *gin.Engine, service svc.UserService) {
	handler := &userHandler{
		userService: service,
	}

	prefix := fmt.Sprintf("%s%s%s", "/", consts.ServiceName, "/api/v1/users")
	userRoutes := r.Group(prefix)

	userRoutes.POST("instant/signup", handler.SignUpHandler)
}

func (h *userHandler) SignUpHandler(c *gin.Context) {

}
