package hubgo

import (
	"context"

	"github.com/clarkmcc/go-hubspot/generated/v3/objects"
)

type objectClient struct {
	ctx    context.Context
	client *objects.APIClient
}

func newObjectClient(ctx context.Context) *objectClient {
	return &objectClient{
		ctx:    ctx,
		client: objects.NewAPIClient(objects.NewConfiguration()),
	}
}

func (c objectClient) Create(objectType string, properties map[string]string) (*objects.SimplePublicObject, error) {
	req := c.client.BasicApi.Create(c.ctx, objectType).SimplePublicObjectInput(objects.SimplePublicObjectInput{
		Properties: properties,
	})
	obj, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c objectClient) CreateMany(objectType string, objs []objects.SimplePublicObjectInput) (*objects.BatchResponseSimplePublicObject, error) {
	req := c.client.BatchApi.BatchCreate(c.ctx, objectType).BatchInputSimplePublicObjectInput(
		objects.BatchInputSimplePublicObjectInput{
			Inputs: objs,
		},
	)
	obj, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return obj, nil
}
