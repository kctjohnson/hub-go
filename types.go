package hubgo

import (
	"github.com/clarkmcc/go-hubspot/generated/v3/contacts"
	"github.com/clarkmcc/go-hubspot/generated/v3/deals"
	"github.com/clarkmcc/go-hubspot/generated/v3/objects"
	"github.com/clarkmcc/go-hubspot/generated/v3/owners"
	"github.com/clarkmcc/go-hubspot/generated/v3/quotes"
)

type dealApi interface {
	Get(id string, properties []string) (*deals.SimplePublicObjectWithAssociations, error)
	GetAssociations(id string, toObjectType string) ([]deals.AssociatedId, error)
	Update(id string, properties map[string]string) error
}

type contactApi interface {
	Get(id string, properties []string) (*contacts.SimplePublicObjectWithAssociations, error)
}

type ownerApi interface {
	Get(ownerId int) (*owners.PublicOwner, error)
}

type objectApi interface {
	Create(objectType string, properties map[string]string) (*objects.SimplePublicObject, error)
	CreateMany(objectType string, objs []objects.SimplePublicObjectInput) (*objects.BatchResponseSimplePublicObject, error)
}

type quoteApi interface {
	Create(properties map[string]string) (*quotes.SimplePublicObject, error)
	CreateAssociation(id, toObjectType, toObjectId, associationType string) (*quotes.SimplePublicObjectWithAssociations, error)
	Get(id string, properties []string) (*quotes.SimplePublicObjectWithAssociations, error)
	GetAssociations(id string, toObjectType string) ([]quotes.AssociatedId, error)
	Update(id string, properties map[string]string) error
	Archive(id string) error
}
