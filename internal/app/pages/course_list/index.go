package course_list

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type ViewFunc func() (templ.Component, error)

type CourseListHandler struct {
	view ViewFunc
}

func NewCourseListHandler(viewFunc ViewFunc) *CourseListHandler {
	return &CourseListHandler{
		view: viewFunc,
	}
}

func (h *CourseListHandler) Mount(r chi.Router) {
	r.Get("/", h.index)
}

func (h *CourseListHandler) index(w http.ResponseWriter, r *http.Request) {
	view, err := h.view()
	if err != nil {
		zap.L().Error("failed to get course page view", zap.Error(err))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := view.Render(r.Context(), w); err != nil {
		zap.L().Error("failed to write response", zap.Error(err))
	}
}
