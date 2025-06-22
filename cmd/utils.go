package cmd

func getAllCollectibles(userid uint64) (collectibles []collectiblesAPIResponseData, err error) {
	var nextCursor string
	var response *collectiblesAPIResponse
	for {
		response, err = collectiblesAPI(userid, nextCursor)
		if err != nil {
			return nil, err
		}

		collectibles = append(collectibles, response.Data...)
		if response.NextPageCursor == "" {
			break
		}

		response.NextPageCursor = nextCursor
	}

	return
}
