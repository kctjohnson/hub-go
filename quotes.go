package hubgo

import (
	"context"

	"github.com/clarkmcc/go-hubspot/generated/v3/quotes"
)

type quoteClient struct {
	ctx    context.Context
	client *quotes.APIClient
}

func newQuoteClient(ctx context.Context) *quoteClient {
	return &quoteClient{
		ctx:    ctx,
		client: quotes.NewAPIClient(quotes.NewConfiguration()),
	}
}

func (c quoteClient) Create(properties map[string]string) (*quotes.SimplePublicObject, error) {
	req := c.client.BasicApi.Create(c.ctx).SimplePublicObjectInput(quotes.SimplePublicObjectInput{
		Properties: properties,
	})
	obj, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c quoteClient) CreateAssociation(id, toObjectType, toObjectId, associationType string) (*quotes.SimplePublicObjectWithAssociations, error) {
	req := c.client.AssociationsApi.AssociationsCreate(c.ctx, id, toObjectType, toObjectId, associationType)
	obj, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c quoteClient) Get(id string, properties []string) (*quotes.SimplePublicObjectWithAssociations, error) {
	req := c.client.BasicApi.GetByID(c.ctx, id).Properties(properties)
	obj, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c quoteClient) GetAssociations(id string, toObjectType string) ([]quotes.AssociatedId, error) {
	req := c.client.AssociationsApi.AssociationsGetAll(c.ctx, id, toObjectType)
	obj, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return obj.Results, nil
}

func (c quoteClient) Update(id string, properties map[string]string) error {
	req := c.client.BasicApi.Update(c.ctx, id).SimplePublicObjectInput(quotes.SimplePublicObjectInput{
		Properties: properties,
	})
	_, _, err := req.Execute()
	if err != nil {
		return err
	}
	return nil
}

func (c quoteClient) Archive(id string) error {
	_, err := c.client.BasicApi.Archive(c.ctx, id).Execute()
	if err != nil {
		return err
	}
	return nil
}
