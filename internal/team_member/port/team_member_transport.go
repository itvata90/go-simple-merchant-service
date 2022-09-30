package port

import "net/http"

type TeamMemberHandler interface {
	GetTeamMembers(w http.ResponseWriter, r *http.Request)
	GetTeamMemberById(w http.ResponseWriter, r *http.Request)
	GetTeamMemberByMerchantCode(w http.ResponseWriter, r *http.Request)
	CreateTeamMember(w http.ResponseWriter, r *http.Request)
	UpdateTeamMember(w http.ResponseWriter, r *http.Request)
	DeleteTeamMember(w http.ResponseWriter, r *http.Request)
}
