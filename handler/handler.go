package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"padiplace_ijs/app/usecase/crud_po"
	"padiplace_ijs/app/usecase/crud_product"
	"padiplace_ijs/entity"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type restHandler struct {
	crud_product_uc crud_product.UseCase
	crud_po_uc      crud_po.UseCase
}

func ProductHandler(crud_product_uc crud_product.UseCase) RestHandler {
	return &restHandler{crud_product_uc: crud_product_uc}
}

func POHandler(crud_po_uc crud_po.UseCase) RestHandler {
	return &restHandler{crud_po_uc: crud_po_uc}
}

//product
func (h *restHandler) GetProducts(c *gin.Context) {
	data, err := h.crud_product_uc.GetAll()
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: data})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) GetProduct(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	data, err := h.crud_product_uc.Get(id)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: data})
	} else {
		if errors.Is(err, crud_product.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}
	}
}

func (h *restHandler) CreateProduct(c *gin.Context) {
	param := &entity.Product{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	result, err := h.crud_product_uc.Create(*param)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: result})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) UpdateProduct(c *gin.Context) {
	param := &entity.Product{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	paramId := c.Param("id")
	param.IdProduct, _ = strconv.Atoi(paramId)

	result, err := h.crud_product_uc.Update(*param)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: result})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) DeleteProduct(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	err := h.crud_product_uc.Delete(id)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: fmt.Sprintf("id:%d. successfully deleted", id)})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

//po
// func (h *restHandler) GetPOS(c *gin.Context) {
// 	data, err := h.crud_po_uc.GetAll()
// 	if err == nil {
// 		c.JSON(http.StatusOK, SuccessResponse{Data: data})
// 	} else {
// 		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
// 	}
// }

func (h *restHandler) GetPO(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	data, err := h.crud_po_uc.Get(id)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: data})
	} else {
		if errors.Is(err, crud_product.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}
	}
}

func (h *restHandler) CreatePO(c *gin.Context) {
	param := &entity.PO{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	result, err := h.crud_po_uc.Create(*param)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: result})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) UpdatePO(c *gin.Context) {
	param := &entity.PO{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	paramId := c.Param("id")
	param.IdPO, _ = strconv.Atoi(paramId)

	result, err := h.crud_po_uc.Update(*param)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: result})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) DeletePO(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	err := h.crud_po_uc.Delete(id)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: fmt.Sprintf("id:%d. successfully deleted", id)})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
}
