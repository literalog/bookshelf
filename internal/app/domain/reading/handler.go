package reading

import (
	"encoding/json"
	"net/http"

	"github.com/literalog/bookshelf/pkg/models"
	"github.com/literalog/cerrors"

	"github.com/gorilla/mux"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetByUserId(w http.ResponseWriter, r *http.Request)
	GetByBookId(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)

	Routes() *mux.Router
}

type handler struct {
	service Service
	router  *mux.Router
}

func NewHandler(s Service) Handler {
	h := &handler{
		service: s,
		router:  mux.NewRouter(),
	}

	h.setupRoutes()

	return h
}

func (h *handler) setupRoutes() {
	h.router.HandleFunc("/", h.Create).Methods(http.MethodPost)
	h.router.HandleFunc("/", h.Update).Methods(http.MethodPut)
	h.router.HandleFunc("/{id}", h.Delete).Methods(http.MethodDelete)
	h.router.HandleFunc("/user/{id}", h.GetByUserId).Methods(http.MethodGet)
	h.router.HandleFunc("/book/{id}", h.GetByBookId).Methods(http.MethodGet)
	h.router.HandleFunc("/", h.GetAll).Methods(http.MethodGet)
}

func (h *handler) Routes() *mux.Router {
	return h.router
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := new(models.ReadingRequest)
	json.NewDecoder(r.Body).Decode(&req)

	rd := models.NewReading(*req)
	if err := h.service.Create(ctx, rd); err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rd)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := new(models.ReadingRequest)
	json.NewDecoder(r.Body).Decode(&req)

	rd := models.NewReading(*req)
	if err := h.service.Update(ctx, rd); err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rd)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]

	if err := h.service.Delete(ctx, id); err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) GetByUserId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]

	rd, err := h.service.GetByUserId(ctx, id)
	if err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rd)
}

func (h *handler) GetByBookId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]

	rd, err := h.service.GetByBookId(ctx, id)
	if err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rd)
}

func (h *handler) GetById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]

	rd, err := h.service.GetById(ctx, id)
	if err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rd)
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	rds, err := h.service.GetAll(ctx)
	if err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rds)
}
