package hubgo

import (
	"context"

	"github.com/clarkmcc/go-hubspot/generated/v3/contacts"
)

type contactClient struct {
	ctx    context.Context
	client *contacts.APIClient
}

func newContactClient(ctx context.Context) *contactClient {
	return &contactClient{
		ctx:    ctx,
		client: contacts.NewAPIClient(contacts.NewConfiguration()),
	}
}

func (c contactClient) Get(id string, properties []string) (*contacts.SimplePublicObjectWithAssociations, error) {
	req := c.client.BasicApi.GetByID(c.ctx, id).Properties(properties)
	obj, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return obj, nil
}
