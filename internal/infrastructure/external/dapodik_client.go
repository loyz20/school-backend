package external

import (
	"encoding/json"
	"fmt"
	"net/http"
	"school-backend/internal/dto"
)

type DapodikClient struct {
	BaseURL string
}

func NewDapodikClient(baseURL string) *DapodikClient {
	return &DapodikClient{BaseURL: baseURL}
}

func (c *DapodikClient) GetFullPesertaDidik(sekolahID string) (*struct {
	Rows []dto.SiswaDapodikFull `json:"rows"`
}, error) {
	url := fmt.Sprintf("%s/peserta-didik-full?sekolah_id=%s", c.BaseURL, sekolahID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Rows []dto.SiswaDapodikFull `json:"rows"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
