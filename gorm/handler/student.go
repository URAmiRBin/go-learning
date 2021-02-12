package handler

import (
	"errors"
	"gorm/model"
	"gorm/store"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Student struct {
	store store.Student
}

type request struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	StudentID uint64 `json:"student_id"`
}

func NewStudent(store store.Student) *Student {
	return &Student{
		store: store,
	}
}

func (s *Student) Create(c echo.Context) error {
	var req request

	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	st := model.Student{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		StudentID: req.StudentID,
	}

	if st.FirstName == "" || st.LastName == "" {
		return echo.ErrBadRequest
	}

	if err := s.store.Save(st); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, "Your student is created")
}

func (s *Student) Find(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))

	if err != nil {
		return echo.ErrBadRequest
	}

	st, err := s.store.Load(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.ErrNotFound
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, st)
}
