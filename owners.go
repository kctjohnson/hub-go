package hubgo

import (
	"context"

	"github.com/clarkmcc/go-hubspot/generated/v3/owners"
)

type ownerClient struct {
	ctx    context.Context
	client *owners.APIClient
}

func newOwnerClient(ctx context.Context) *ownerClient {
	return &ownerClient{
		ctx:    ctx,
		client: owners.NewAPIClient(owners.NewConfiguration()),
	}
}

func (c ownerClient) Get(ownerId int) (*owners.PublicOwner, error) {
	req := c.client.OwnersApi.GetByID(c.ctx, int32(ownerId))
	obj, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return obj, nil
}
