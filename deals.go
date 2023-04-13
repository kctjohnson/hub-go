package hubgo

import (
	"context"

	"github.com/clarkmcc/go-hubspot/generated/v3/deals"
)

type dealClient struct {
	ctx    context.Context
	client *deals.APIClient
}

func newDealClient(ctx context.Context) *dealClient {
	return &dealClient{
		ctx:    ctx,
		client: deals.NewAPIClient(deals.NewConfiguration()),
	}
}

func (c dealClient) Get(id string, properties []string) (*deals.SimplePublicObjectWithAssociations, error) {
	req := c.client.BasicApi.GetByID(c.ctx, id).Properties(properties)
	obj, _, err := req.Execute()
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c dealClient) GetAssociations(id string, toObjectType string) ([]deals.AssociatedId, error) {
	req := c.client.AssociationsApi.AssociationsGetAll(c.ctx, id, toObjectType)
	obj, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return obj.Results, nil
}

func (c dealClient) Update(id string, properties map[string]string) error {
	req := c.client.BasicApi.Update(c.ctx, id).SimplePublicObjectInput(deals.SimplePublicObjectInput{
		Properties: properties,
	})
	_, _, err := req.Execute()
	if err != nil {
		return err
	}
	return nil
}
