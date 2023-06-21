package api

import "github.com/McCune1224/space-jam/types"

func (stc *SpaceTradersClient) GetAgent() (types.Agent, error) {
	var agentResponse types.Agent
	url := BASE_URL + "/my/agent"
	err := stc.httpClient.Get(url, stc.authHeader, nil, &agentResponse)
	if err != nil {
		return agentResponse, err
	}
	return agentResponse, nil
}
