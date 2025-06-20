package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"school-backend/internal/dto"
	"school-backend/pkg/utils"
)

type DapodikClient struct {
	BaseURL string
}

func NewDapodikClient(baseURL string) *DapodikClient {
	return &DapodikClient{BaseURL: baseURL}
}

func (c *DapodikClient) GetFullPesertaDidik(npsn string) (*struct {
	Rows []dto.SiswaDapodikFull `json:"rows"`
}, error) {
	url := fmt.Sprintf("%s/getPesertaDidik?npsn=%s", c.BaseURL, npsn)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	dapoToken := utils.GetEnv("DAPO_TOKEN", "")

	// Set header Authorization
	req.Header.Set("Authorization", "Bearer "+dapoToken)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Rows []dto.SiswaDapodikFull `json:"rows"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *DapodikClient) GetFullPengguna(npsn string) (*struct {
	Rows []dto.PenggunaDapodikFull `json:"rows"`
}, error) {
	url := fmt.Sprintf("%s/getPengguna?npsn=%s", c.BaseURL, npsn)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	dapoToken := utils.GetEnv("DAPO_TOKEN", "")

	// Set header Authorization
	req.Header.Set("Authorization", "Bearer "+dapoToken)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Rows []dto.PenggunaDapodikFull `json:"rows"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *DapodikClient) GetFullRombel(npsn string) (*dto.FullRombelResponse, error) {
	url := fmt.Sprintf("%s/getRombonganBelajar?npsn=%s", c.BaseURL, npsn)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	dapoToken := utils.GetEnv("DAPO_TOKEN", "")
	req.Header.Set("Authorization", "Bearer "+dapoToken)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
	}

	var result dto.FullRombelResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *DapodikClient) GetFullSekolah(baseURL string, npsn string, token string) (*struct {
	Rows []dto.SekolahDapodikFull `json:"rows"`
}, error) {
	url := fmt.Sprintf("%s/getSekolah?npsn=%s", baseURL, npsn)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	dapoToken := token

	// Set header Authorization
	req.Header.Set("Authorization", "Bearer "+dapoToken)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Rows []dto.SekolahDapodikFull `json:"rows"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
