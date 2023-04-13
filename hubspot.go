package hubgo

import (
	"context"

	"github.com/clarkmcc/go-hubspot"
)

type HubspotClient struct {
	ctx context.Context

	Deals    dealApi
	Contacts contactApi
	Owners   ownerApi
	Objects  objectApi
	Quotes   quoteApi
}

func NewHubspotClient(token string) *HubspotClient {
	auth := hubspot.NewTokenAuthorizer(token)
	ctx := hubspot.WithAuthorizer(context.Background(), auth)

	return &HubspotClient{
		ctx:      ctx,
		Deals:    newDealClient(ctx),
		Contacts: newContactClient(ctx),
		Owners:   newOwnerClient(ctx),
		Objects:  newObjectClient(ctx),
		Quotes:   newQuoteClient(ctx),
	}
}
