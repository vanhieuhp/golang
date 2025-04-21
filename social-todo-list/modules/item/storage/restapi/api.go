package restapi

import (
	"context"
	"errors"
	"fmt"
	"github.com/200Lab-Education/go-sdk/logger"
	"github.com/go-resty/resty/v2"
)

type itemService struct {
	client     *resty.Client
	serviceURL string
	logger     logger.Logger
}

func New(serviceURL string, logger logger.Logger) *itemService {
	return &itemService{
		client:     resty.New(),
		serviceURL: serviceURL,
		logger:     logger,
	}
}

func (s *itemService) GetItemLikes(ctx context.Context, ids []int) (map[int]int, error) {
	type requestBody struct {
		Ids []int `json:"ids"`
	}

	var response struct {
		Data map[int]int `json:"data"`
	}

	resp, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody{Ids: ids}).
		SetResult(&response).
		Post(fmt.Sprintf("%s/%s", s.serviceURL, "rpc/get_item_likes"))

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		// log.Println(resp.RawResponse)
		s.logger.Error(resp.RawResponse)
		return nil, errors.New("cannot call api get item likes")
	}

	return response.Data, nil

}
