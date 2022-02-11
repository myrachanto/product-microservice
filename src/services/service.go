package service

import (
	"fmt"

	httperors "github.com/myrachanto/custom-http-error"
	"github.com/myrachanto/microservice/product/src/model"
	r "github.com/myrachanto/microservice/product/src/repository"
)

//productService ...
var (
	ProductService ProductServiceInterface = &productService{}
)

type ProductServiceInterface interface {
	Create(product *model.Product) (string, httperors.HttpErr)
	GetOne(code int) (*model.Product, httperors.HttpErr)
	GetAll(search string, page, pagesize int) ([]model.Product, httperors.HttpErr)
	Update(id int, product *model.Product) (*model.Product, httperors.HttpErr)
	Delete(id int) (string, httperors.HttpErr)
	GetOnebyCode(code string) (*model.Product, httperors.HttpErr)
}

type productService struct {
	repository r.ProductRepoInterface
}

func NewproductService(repo r.ProductRepoInterface) ProductServiceInterface {
	return &productService{
		repo,
	}
}

func (service productService) Create(product *model.Product) (string, httperors.HttpErr) {
	if err := product.Validate(); err != nil {
		return "", err
	}
	s, err1 := r.Productrepo.Create(product)
	if err1 != nil {
		return "", err1
	}
	return s, nil

}
func (service productService) GetOne(code int) (*model.Product, httperors.HttpErr) {
	product, err1 := r.Productrepo.GetOne(code)
	if err1 != nil {
		return nil, err1
	}
	return product, nil
}
func (service productService) GetOnebyCode(code string) (*model.Product, httperors.HttpErr) {
	product, err1 := r.Productrepo.GetOnebyCode(code)
	if err1 != nil {
		return nil, err1
	}
	return product, nil
}

func (service productService) GetAll(search string, page, pagesize int) ([]model.Product, httperors.HttpErr) {
	results, err := r.Productrepo.GetAll(search, page, pagesize)
	return results, err
}

// func (service productService) UpdateRole(code, admin, supervisor, employee, level, productcode string) (string, *httperors.HttpError) {
// 	product, err1 := r.Productrepo.UpdateRole(code, admin, supervisor, employee, level, productcode)
// 	return product, err1
// }

func (service productService) Update(id int, product *model.Product) (*model.Product, httperors.HttpErr) {
	fmt.Println("update1-controller")
	fmt.Println(id)
	product, err1 := r.Productrepo.Update(id, product)
	if err1 != nil {
		return nil, err1
	}

	return product, nil
}
func (service productService) Delete(id int) (string, httperors.HttpErr) {
	success, failure := r.Productrepo.Delete(id)
	return success, failure
}
