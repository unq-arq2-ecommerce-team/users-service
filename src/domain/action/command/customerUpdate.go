package command

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/model"
)

type UpdateCustomer struct {
	customerRepo          model.CustomerRepository
	findCustomerByIdQuery query.FindCustomerById
}

func NewUpdateCustomer(customerRepo model.CustomerRepository, findCustomer query.FindCustomerById) *UpdateCustomer {
	return &UpdateCustomer{
		customerRepo:          customerRepo,
		findCustomerByIdQuery: findCustomer,
	}
}

func (c UpdateCustomer) Do(ctx context.Context, customerId int64, updateCustomer model.UpdateCustomer) error {
	customer, err := c.findCustomerByIdQuery.Do(ctx, customerId)
	if err != nil {
		return err
	}
	customer.Merge(updateCustomer)
	if _, err := c.customerRepo.Update(ctx, *customer); err != nil {
		return err
	}
	return nil
}
