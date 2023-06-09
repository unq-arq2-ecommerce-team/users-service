package command

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/model"
)

type DeleteCustomer struct {
	customerRepo          model.CustomerRepository
	findCustomerByIdQuery query.FindCustomerById
}

func NewDeleteCustomer(customerRepo model.CustomerRepository, findCustomer query.FindCustomerById) *DeleteCustomer {
	return &DeleteCustomer{
		customerRepo:          customerRepo,
		findCustomerByIdQuery: findCustomer,
	}
}

func (c DeleteCustomer) Do(ctx context.Context, id int64) error {
	_, err := c.findCustomerByIdQuery.Do(ctx, id)
	if err != nil {
		return err
	}
	if _, err := c.customerRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
