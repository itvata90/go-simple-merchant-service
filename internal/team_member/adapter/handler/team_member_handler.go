package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	sv "github.com/core-go/core"

	domain "merchant-service/internal/team_member/entity"
	"merchant-service/internal/team_member/service"
)

func NewTeamMemberHandler(service service.TeamMemberService, validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), logError func(context.Context, string, ...map[string]interface{})) *HttpTeamMemberHandler {
	return &HttpTeamMemberHandler{service: service, validate: validate, logError: logError}
}

type HttpTeamMemberHandler struct {
	service  service.TeamMemberService
	validate func(context.Context, interface{}) ([]sv.ErrorMessage, error)
	logError func(context.Context, string, ...map[string]interface{})
}

func (h *HttpTeamMemberHandler) GetTeamMembers(w http.ResponseWriter, r *http.Request) {
	pageIndexParam := r.URL.Query().Get("pageIndex")
	pageSizeParam := r.URL.Query().Get("pageSize")
	pageIndex, err := strconv.Atoi(pageIndexParam)
	if err != nil {
		http.Error(w, "pageIndex must be an integer", http.StatusBadRequest)
		return
	}
	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil {
		http.Error(w, "pageSize must be an integer", http.StatusBadRequest)
		return
	}
	res, err := h.service.GetTeamMembers(r.Context(), pageSize, pageIndex)
	sv.Respond(w, r, http.StatusOK, res, err, h.logError, nil)

}

func (h *HttpTeamMemberHandler) GetTeamMemberById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "id cannot be empty", http.StatusBadRequest)
		return
	}
	teamMember, err := h.service.GetTeamMemberById(r.Context(), id)
	sv.Respond(w, r, http.StatusOK, teamMember, err, h.logError, nil)
}

func (h *HttpTeamMemberHandler) GetTeamMemberByMerchantCode(w http.ResponseWriter, r *http.Request) {
	merchantCode := mux.Vars(r)["merchantCode"]
	if len(merchantCode) == 0 {
		http.Error(w, "merchantCode cannot be empty", http.StatusBadRequest)
		return
	}
	pageIndexParam := r.URL.Query().Get("pageIndex")
	pageSizeParam := r.URL.Query().Get("pageSize")
	pageIndex, err := strconv.Atoi(pageIndexParam)
	if err != nil {
		http.Error(w, "pageIndex must be an integer", http.StatusBadRequest)
		return
	}
	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil {
		http.Error(w, "pageSize must be an integer", http.StatusBadRequest)
		return
	}
	if len(merchantCode) > 0 {
		res, err := h.service.GetTeamMemberByMerchantCode(r.Context(), merchantCode, pageSize, pageIndex)
		if err != nil {
			h.logError(r.Context(), err.Error())
			http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
			return
		}
		if res == nil && len(res.Data) == 0 {
			JSON(w, http.StatusNotFound, nil)
			return
		}
		JSON(w, http.StatusOK, res)
	}
}

func (h *HttpTeamMemberHandler) CreateTeamMember(w http.ResponseWriter, r *http.Request) {
	var teamMember domain.TeamMember
	er1 := sv.Decode(w, r, &teamMember)
	if er1 == nil {
		errors, er2 := h.validate(r.Context(), &teamMember)
		if er2 != nil {
			h.logError(r.Context(), er2.Error())
			http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
			return
		}
		if len(errors) > 0 {
			h.logError(r.Context(), er2.Error())
			JSON(w, http.StatusUnprocessableEntity, errors)
			return
		}
		res, er3 := h.service.CreateTeamMember(r.Context(), &teamMember)
		sv.Respond(w, r, http.StatusCreated, res, er3, h.logError, nil)
		return
	} else {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
}

func (h *HttpTeamMemberHandler) UpdateTeamMember(w http.ResponseWriter, r *http.Request) {
	var teamMember domain.TeamMember
	er1 := json.NewDecoder(r.Body).Decode(&teamMember)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	if len(teamMember.Id) == 0 {
		teamMember.Id = id
	} else if id != teamMember.Id {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}
	errors, er2 := h.validate(r.Context(), &teamMember)
	if er2 != nil {
		h.logError(r.Context(), er2.Error())
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		h.logError(r.Context(), er2.Error())
		JSON(w, http.StatusUnprocessableEntity, errors)
		return
	}
	res, er3 := h.service.UpdateTeamMember(r.Context(), &teamMember)
	sv.Respond(w, r, http.StatusOK, res, er3, h.logError, nil)

}

func (h *HttpTeamMemberHandler) DeleteTeamMember(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "code cannot be empty", http.StatusBadRequest)
		return
	}
	res, err := h.service.DeleteTeamMember(r.Context(), id)
	sv.Respond(w, r, http.StatusOK, res, err, h.logError, nil)

}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}
