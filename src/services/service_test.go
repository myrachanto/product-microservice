package service

import (
	"testing"

	httperors "github.com/myrachanto/custom-http-error"
	"github.com/myrachanto/microservice/product/src/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

var (
	Product = &model.Product{
		Name:        "product one",
		Description: "Product one description",
		Picture:     "img/product.png",
		Productcode: "prod1234567",
		Usercode:    "user1234567",
	}
)

type ProductMockInterface interface {
	Create(Product *model.Product) (*model.Product, httperors.HttpErr)
	GetOne(id string) (*model.Product, httperors.HttpErr)
	GetAll() ([]*model.Product, httperors.HttpErr)
	Update(code string, Product *model.Product) httperors.HttpErr
}

func (mock MockRepository) Create(Product *model.Product) (*model.Product, httperors.HttpErr) {
	args := mock.Called()
	result := args.Get(0)
	Product, err := result.(*model.Product), args.Error(1)
	if err != nil {
		return nil, httperors.NewNotFoundError("Something went wrong creating the resourse")
	}
	return Product, nil
}
func (mock MockRepository) GetOne() (*model.Product, httperors.HttpErr) {
	args := mock.Called()
	result := args.Get(0)
	Product, err := result.(*model.Product), args.Error(1)
	if err != nil {
		return nil, httperors.NewNotFoundError("Something went wrong getting the resourse")
	}
	return Product, nil
}
func (mock MockRepository) GetAll() ([]*model.Product, httperors.HttpErr) {
	args := mock.Called()
	result := args.Get(0)
	Products, err := result.([]*model.Product), args.Error(1)
	if err != nil {
		return nil, httperors.NewNotFoundError("Something went wrong getting the resourses")
	}
	return Products, nil
}
func (mock MockRepository) Update(code string, Product *model.Product) httperors.HttpErr {
	args := mock.Called()
	result := args.Get(0)
	_, err := result.(*model.Product), args.Error(1)
	if err != nil {
		return httperors.NewNotFoundError("Something went wrong updating the resourse")
	}
	return nil
}
func TestGetAll(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("Create").Return(Product, nil)
	_, _ = mockRepo.Create(Product)
	mockRepo.On("GetAll").Return([]*model.Product{Product}, nil)
	results, _ := mockRepo.GetAll()
	//mock assertion: behavioral
	mockRepo.AssertExpectations(t)
	//data assertion
	assert.Equal(t, Product.Name, results[0].Name)
	assert.Equal(t, Product.Description, results[0].Description)
	assert.Equal(t, Product.Picture, results[0].Picture)

}
func TestGetOne(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("Create").Return(Product, nil)
	_, _ = mockRepo.Create(Product)
	mockRepo.On("GetOne").Return(Product, nil)
	results, _ := mockRepo.GetOne()
	//mock assertion: behavioral
	mockRepo.AssertExpectations(t)
	//data assertion
	assert.Equal(t, Product.Name, results.Name)
	assert.Equal(t, Product.Description, results.Description)
	assert.Equal(t, Product.Picture, results.Picture)

}
func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("Create").Return(Product, nil)
	result, err := mockRepo.Create(Product)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, Product.Name, result.Name)
	assert.Equal(t, Product.Description, result.Description)
	assert.Equal(t, Product.Picture, result.Picture)
	assert.Nil(t, err)

}
