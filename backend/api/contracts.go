package api

import "github.com/McCune1224/space-jam/types"

func (stc *SpaceTradersClient) ListContracts(pageOpts ...PaginationOptions) (PaginatedApiDataResponse[types.Contract], error) {
	endpointUrl := BASE_URL + "/my/contracts"
	if len(pageOpts) == 0 {
		pageOpts = append(pageOpts, PaginationOptions{Limit: 10, Page: 1})
	} else {
		if pageOpts[0].Limit > 20 {
			pageOpts[0].Limit = 20
		}
	}

	pageQueryParams := pageOpts[0].ToQueryParams()

	var contractsResponse struct {
		Data []types.Contract `json:"data"`
		Meta types.Meta       `json:"meta"`
	}
	err := stc.httpClient.Get(endpointUrl, stc.authHeader, pageQueryParams, &contractsResponse)
	if err != nil {
		return contractsResponse, err
	}

	return contractsResponse, nil
}

func (stc *SpaceTradersClient) GetContract() {
}

func (stc *SpaceTradersClient) AcceptContract() {
}

func (stc *SpaceTradersClient) DeliverCargoToContract() {
}

func (stc *SpaceTradersClient) FufillContract() {
}
