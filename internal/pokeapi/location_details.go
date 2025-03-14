package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationDetails(locationName string) (RespLocation, error) {
	url := baseURL + "/location-area/" + locationName

	val, ok := c.cache.Get(url)
	if ok {
		locationResp := RespLocation{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespLocation{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocation{}, err
	}

	locationResp := RespLocation{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespLocation{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
