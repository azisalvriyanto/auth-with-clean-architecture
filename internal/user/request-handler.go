package user

import (
	"auth-with-clean-architecture/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandler struct {
	ctrl *Controller
}

func NewRequestHandler(ctrl *Controller) *RequestHandler {
	return &RequestHandler{
		ctrl: ctrl,
	}
}

func DefaultRequestHandler(db *gorm.DB) *RequestHandler {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}

type CreateRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h RequestHandler) ShowAll(c *gin.Context) {
	res, err := h.ctrl.ShowAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Meta: dto.MetaResponse{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Meta: dto.MetaResponse{
			Success: true,
			Message: "",
		},
		Data: res,
	})
}

func (h RequestHandler) Create(c *gin.Context) {
	var req CreateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Meta: dto.MetaResponse{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	res, err := h.ctrl.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Meta: dto.MetaResponse{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Meta: dto.MetaResponse{
			Success: true,
			Message: res.Message,
		},
		Data: res.Data,
	})
}

func (h RequestHandler) Show(c *gin.Context) {
	ID := c.Param("ID")
	res, err := h.ctrl.Show(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Meta: dto.MetaResponse{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Meta: dto.MetaResponse{
			Success: true,
			Message: "",
		},
		Data: res,
	})
}

func (h RequestHandler) Update(c *gin.Context) {
	user := User{}
	ID := c.Param("ID")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Meta: dto.MetaResponse{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	res, err := h.ctrl.Update(ID, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Meta: dto.MetaResponse{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Meta: dto.MetaResponse{
			Success: true,
			Message: res.Message,
		},
		Data: res.Data,
	})
}

func (h RequestHandler) Destroy(c *gin.Context) {
	ID := c.Param("ID")
	res, err := h.ctrl.Destroy(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Meta: dto.MetaResponse{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Meta: dto.MetaResponse{
			Success: true,
			Message: res.Message,
		},
		Data: res.Data,
	})
}
