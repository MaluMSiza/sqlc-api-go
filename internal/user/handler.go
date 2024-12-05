package user

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/jackc/pgx/v5/pgtype"
	"strconv"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUserHandler(c echo.Context) error {
	var req struct {
		Name      string       `json:"name"`
		BirthDate pgtype.Date `json:"birth_date"`
		Cpf       string      `json:"cpf"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	user, err := h.service.CreateUser(c.Request().Context(), req.Name, req.BirthDate, req.Cpf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUserHandler(c echo.Context) error {
	id := c.Param("id")
	parsedID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	user, err := h.service.GetUser(c.Request().Context(), int32(parsedID))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetAllUsersHandler(c echo.Context) error {

	user, err := h.service.GetAllUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Users not found"})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUserAgeHandler(c echo.Context) error {
	id := c.Param("id")
	parsedID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	age, err := h.service.GetUserAge(c.Request().Context(), int32(parsedID))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, age)
}


func (h *UserHandler) UpdateUserHandler(c echo.Context) error {
	id := c.Param("id")

	parsedID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var req struct {
		Name      string       `json:"name"`
		BirthDate pgtype.Date `json:"birth_date"`
		Cpf       string      `json:"cpf"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	user, err := h.service.UpdateUser(c.Request().Context(), int32(parsedID), req.Name, req.BirthDate, req.Cpf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUserHandler(c echo.Context) error {
	id := c.Param("id")

	parsedID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	user, err := h.service.DeleteUser(c.Request().Context(), int32(parsedID))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}
