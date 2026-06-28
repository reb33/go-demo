package stat

type GetStatResponse struct {
	Period string `json:"period"`
	Sum    int    `json:"sum"`
}

func ConvertRepoStatsToPayload(repoStats []RepoGetStatsResponse) []GetStatResponse {
	payload := make([]GetStatResponse, len(repoStats))
	for i, s := range repoStats {
		payload[i] = GetStatResponse{
			Period: s.Period,
			Sum:    s.Sum,
		}
	}
	return payload
}