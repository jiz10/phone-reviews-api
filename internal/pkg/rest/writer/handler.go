package writer

import (
	"encoding/json"
	"net/http"
	"phone-reviews-api/internal/pkg/domain/smartphone"
	"phone-reviews-api/internal/pkg/repository/database"
)

func (h *CreateSmartphoneHandler) SaveSmartphoneHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cmd := parseRequest(r)
	res, err := h.Create(cmd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error in create smartphone"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}

type CreateSmartphoneHandler struct {
	smartphone.CreateGateway
}

func NewCreateSmartphoneHandler(client *database.MySqlClient) *CreateSmartphoneHandler {
	return &CreateSmartphoneHandler{smartphone.NewSmartphoneCreateGateway(client)}
}

func parseRequest(r *http.Request) *smartphone.CreateSmartphoneCMD {
	body := r.Body
	defer body.Close()
	var cmd smartphone.CreateSmartphoneCMD

	_ = json.NewDecoder(body).Decode(&cmd)

	return &cmd
}
