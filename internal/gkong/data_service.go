package gkong

import (
	"context"
)

type DataService[T, Y any] interface {
	Create(ctx context.Context, nameOrID *string, item *T) (*T, error)
	Get(ctx context.Context, nameOrID *string) (*T, error)
	Update(ctx context.Context, nameOrID *string, item *T) (*T, error)
	Delete(ctx context.Context, nameOrID *string) error
	List(ctx context.Context, opt *Y) ([]*T, *T, error)
	ListAll(ctx context.Context) ([]*T, error)
}

type dataService[T, Y any] struct {
}

func (pkg *dataService[T, Y]) Create(ctx context.Context, nameOrID *string, item *T) (*T, error) {
	return nil, nil
}

func (pkg *dataService[T, Y]) Get(ctx context.Context, nameOrID, groupOrID *string) (*T, error) {
	return nil, nil
}

func (pkg *dataService[T, Y]) Update(ctx context.Context, nameOrID *string, aclGroup *T) (*T, error) {
	return nil, nil
}

func (pkg *dataService[T, Y]) Delete(ctx context.Context, nameOrID, groupOrID *string) error {
	return nil
}

func (pkg *dataService[T, Y]) List(ctx context.Context, opt *Y) ([]*T, *T, error) {
	return nil, nil, nil
}

func (pkg *dataService[T, Y]) ListAll(ctx context.Context) ([]*T, error) {
	return nil, nil
}
