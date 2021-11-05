package user

import (
	"github.com/gin-gonic/gin"

	"go-web-backend/pkg/core"
	"go-web-backend/pkg/errors"
	"go-web-backend/pkg/log"
	metav1 "go-web-backend/pkg/meta/v1"

	"go-web-backend/internal/pkg/code"
	v1 "go-web-backend/internal/pkg/entity/apiserver/v1"
)

// Create add new user to the storage.
func (u *Controller) Create(c *gin.Context) {
	log.C(c).Info("user create function called.")

	var r v1.User

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	if errs := r.Validate(); len(errs) != 0 {
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	// Insert the user to the storage.
	if err := u.srv.Users().Create(c, &r, metav1.CreateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, r)
}
