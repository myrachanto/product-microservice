package repository

import (
	"strconv"

	httperors "github.com/myrachanto/custom-http-error"
	"github.com/myrachanto/microservice/product/src/model"
)

//productrepo ...
var (
	Productrepo ProductRepoInterface = &productrepo{}
)

type Key struct {
	EncryptionKey string `mapstructure:"EncryptionKey"`
}

type productrepo struct{}

type ProductRepoInterface interface {
	Create(product *model.Product) (string, httperors.HttpErr)
	all() (t []model.Product, r httperors.HttpErr)
	GetOne(id int) (*model.Product, httperors.HttpErr)
	GetOnebyCode(id string) (*model.Product, httperors.HttpErr)
	productExistbycode(code string) bool
	productbycode(code string) *model.Product
	GetAll(search string, page, pagesize int) ([]model.Product, httperors.HttpErr)
	Update(id int, product *model.Product) (*model.Product, httperors.HttpErr)
	Delete(id int) (string, httperors.HttpErr)
	geneCode() (string, httperors.HttpErr)
	productExist(email string) bool
	productExistByid(id int) bool
}

func NewproductRepo() *productrepo {
	return &productrepo{}
}
func (productRepo productrepo) Create(product *model.Product) (string, httperors.HttpErr) {
	if err := product.Validate(); err != nil {
		return "", err
	}
	code, x := productRepo.geneCode()
	if x != nil {
		return "", x
	}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return "", err1
	}
	product.Productcode = code
	GormDB.Create(&product)
	IndexRepo.DbClose(GormDB)
	return "product created successifully", nil
}
func (productRepo productrepo) all() (t []model.Product, r httperors.HttpErr) {

	product := model.Product{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&product).Find(&t)
	IndexRepo.DbClose(GormDB)
	return t, nil

}
func (productRepo productrepo) geneCode() (string, httperors.HttpErr) {
	product := model.Product{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return "", err1
	}
	err := GormDB.Last(&product)
	if err.Error != nil {
		var c1 uint = 1
		code := "productCode" + strconv.FormatUint(uint64(c1), 10)
		return code, nil
	}
	c1 := product.ID + 1
	code := "productCode" + strconv.FormatUint(uint64(c1), 10)
	IndexRepo.DbClose(GormDB)
	return code, nil

}
func (productRepo productrepo) GetOne(id int) (*model.Product, httperors.HttpErr) {
	ok := productRepo.productExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("product with that code does not exists!")
	}
	product := model.Product{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}

	GormDB.Model(&product).Where("id = ?", id).First(&product)
	IndexRepo.DbClose(GormDB)
	return &product, nil
}
func (productRepo productrepo) GetOnebyCode(code string) (*model.Product, httperors.HttpErr) {
	ok := productRepo.productExistbycode(code)
	if !ok {
		return nil, httperors.NewNotFoundError("product with that code does not exists!")
	}
	product := model.Product{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}

	GormDB.Model(&product).Where("productcode = ?", code).First(&product)
	IndexRepo.DbClose(GormDB)
	return &product, nil
}
func (productRepo productrepo) productExistbycode(code string) bool {
	u := model.Product{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return false
	}
	GormDB.Where("productcode = ?", code).First(&u)
	if u.ID == 0 {
		return false
	}
	IndexRepo.DbClose(GormDB)
	return true

}
func (productRepo productrepo) productbycode(code string) *model.Product {
	u := model.Product{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil
	}
	GormDB.Where("productcode = ?", code).First(&u)
	if u.ID == 0 {
		return nil
	}
	IndexRepo.DbClose(GormDB)
	return &u

}
func (productRepo productrepo) GetAll(search string, page, pagesize int) ([]model.Product, httperors.HttpErr) {
	results := []model.Product{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	if search == "" {
		GormDB.Find(&results)
	}
	// db.Scopes(Paginate(r)).Find(&products)
	GormDB.Scopes(Paginate(page, pagesize)).Where("name LIKE ?", "%"+search+"%").Or("email LIKE ?", "%"+search+"%").Or("company LIKE ?", "%"+search+"%").Find(&results)

	IndexRepo.DbClose(GormDB)
	return results, nil
}

func (productRepo productrepo) Update(id int, product *model.Product) (*model.Product, httperors.HttpErr) {
	ok := productRepo.productExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("product with that id does not exists!")
	}

	uproduct := model.Product{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&product).Where("id = ?", id).First(&uproduct)
	if product.Name == "" {
		product.Name = uproduct.Name
	}
	if product.Description == "" {
		product.Description = uproduct.Description
	}
	if product.SellPrice == 0 {
		product.SellPrice = uproduct.SellPrice
	}
	if product.BuyPrice == 0 {
		product.BuyPrice = uproduct.BuyPrice
	}
	if product.Quantity == 0 {
		product.Quantity = uproduct.Quantity
	}
	if product.Picture == "" {
		product.Picture = uproduct.Picture
	}
	if product.Usercode == "" {
		product.Usercode = uproduct.Usercode
	}
	if product.Productcode == "" {
		product.Productcode = uproduct.Productcode
	}
	GormDB.Save(&product)

	IndexRepo.DbClose(GormDB)

	return product, nil
}
func (productRepo productrepo) Delete(id int) (string, httperors.HttpErr) {
	ok := productRepo.productExistByid(id)
	if !ok {
		return "", httperors.NewNotFoundError("product with that id does not exists!")
	}
	product := model.Product{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return "", err1
	}
	GormDB.Model(&product).Where("id = ?", id).First(&product)
	GormDB.Delete(product)
	IndexRepo.DbClose(GormDB)
	return "deleted successfully", nil
}
func (productRepo productrepo) productExist(email string) bool {
	product := model.Product{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return false
	}
	res := GormDB.First(&product, "email =?", email)
	if res.Error != nil {
		return false
	}
	IndexRepo.DbClose(GormDB)
	return true

}
func (productRepo productrepo) productExistByid(id int) bool {
	product := model.Product{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return false
	}
	res := GormDB.First(&product, "id =?", id)
	if res.Error != nil {
		return false
	}
	IndexRepo.DbClose(GormDB)
	return true

}
