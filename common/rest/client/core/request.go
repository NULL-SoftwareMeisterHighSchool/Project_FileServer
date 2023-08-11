package core_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
)

var client = &http.Client{}

type userInfoResp struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetUserIDFromToken(accessToken string) uint {
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://%s:3000/users/me/tiny", config.CORE_IP), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return 0
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result userInfoResp
	if err = json.Unmarshal(data, &result); err != nil {
		panic(err)
	}

	return uint(result.Id)
}
