package controllers

import (
	"strconv"

	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	httperors "github.com/myrachanto/custom-http-error"
	"github.com/myrachanto/microservice/product/src/model"
	service "github.com/myrachanto/microservice/product/src/services"
)

//productController ..
var (
	ProductController productcontrollerInterface = &productController{}
)

type productController struct {
	service service.ProductServiceInterface
}
type productcontrollerInterface interface {
	Create(c echo.Context) error
	GetOne(c echo.Context) error
	GetAll(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetOnebyCode(c echo.Context) error
}

func NewproductController(ser service.ProductServiceInterface) productcontrollerInterface {
	return &productController{
		ser,
	}
}

/////////controllers/////////////////
func (controller productController) Create(c echo.Context) error {
	product := &model.Product{}
	product.Name = c.FormValue("name")
	product.Description = c.FormValue("description")

	pic, err2 := c.FormFile("picture")
	//    fmt.Println(pic.Filename)
	if err2 != nil {
		httperror := httperors.NewBadRequestError("Invalid picture")
		return c.JSON(httperror.Code(), err2)
	}
	src, err := pic.Open()
	if err != nil {
		httperror := httperors.NewBadRequestError("the picture is corrupted")
		return c.JSON(httperror.Code(), err)
	}
	defer src.Close()
	// filePath := "./public/imgs/products/"
	filePath := "./public/imgs/products/" + pic.Filename
	filePath1 := "/imgs/products/" + pic.Filename
	// Destination
	dst, err4 := os.Create(filePath)
	if err4 != nil {
		httperror := httperors.NewBadRequestError("the Directory mess")
		return c.JSON(httperror.Code(), err4)
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		if err2 != nil {
			httperror := httperors.NewBadRequestError("error filling")
			return c.JSON(httperror.Code(), httperror)
		}
	}

	product.Picture = filePath1
	s, err1 := service.ProductService.Create(product)
	if err1 != nil {
		return c.JSON(err1.Code(), err1)
	}
	if _, err = io.Copy(dst, src); err != nil {
		if err2 != nil {
			httperror := httperors.NewBadRequestError("error filling")
			return c.JSON(httperror.Code(), httperror)
		}
	}
	return c.JSON(http.StatusCreated, s)
}
func (controller productController) GetAll(c echo.Context) error {
	search := string(c.QueryParam("q"))
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid page number")
		return c.JSON(httperror.Code(), httperror)
	}
	pagesize, err := strconv.Atoi(c.QueryParam("pagesize"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid pagesize")
		return c.JSON(httperror.Code(), httperror)
	}

	results, err3 := service.ProductService.GetAll(search, page, pagesize)
	if err3 != nil {
		return c.JSON(err3.Code(), err3)
	}
	return c.JSON(http.StatusOK, results)
}
func (controller productController) GetOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code(), httperror)
	}
	result, problem := service.ProductService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code(), problem)
	}
	return c.JSON(http.StatusOK, result)
}
func (controller productController) GetOnebyCode(c echo.Context) error {
	code := c.Param("code")
	result, problem := service.ProductService.GetOnebyCode(code)
	if problem != nil {
		return c.JSON(problem.Code(), problem)
	}
	return c.JSON(http.StatusOK, result)
}

func (controller productController) Update(c echo.Context) error {
	product := &model.Product{}
	product.Name = c.FormValue("name")
	product.Description = c.FormValue("description")

	pic, err2 := c.FormFile("picture")
	//    fmt.Println(pic.Filename)
	if err2 != nil {
		httperror := httperors.NewBadRequestError("Invalid picture")
		return c.JSON(httperror.Code(), err2)
	}
	src, err := pic.Open()
	if err != nil {
		httperror := httperors.NewBadRequestError("the picture is corrupted")
		return c.JSON(httperror.Code(), err)
	}
	defer src.Close()
	filePath := "./public/imgs/products/" + pic.Filename
	filePath1 := "/imgs/products/" + pic.Filename
	// Destination
	dst, err4 := os.Create(filePath)
	if err4 != nil {
		httperror := httperors.NewBadRequestError("the Directory mess")
		return c.JSON(httperror.Code(), err4)
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		if err2 != nil {
			httperror := httperors.NewBadRequestError("error filling")
			return c.JSON(httperror.Code(), httperror)
		}
	}

	product.Picture = filePath1
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code(), httperror)
	}
	updatedproduct, problem := service.ProductService.Update(id, product)
	if problem != nil {
		return c.JSON(problem.Code(), problem)
	}
	if _, err = io.Copy(dst, src); err != nil {
		if err2 != nil {
			httperror := httperors.NewBadRequestError("error filling")
			return c.JSON(httperror.Code(), httperror)
		}
	}

	return c.JSON(http.StatusOK, updatedproduct)
}

func (controller productController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code(), httperror)
	}
	success, failure := service.ProductService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code(), failure)
	}
	return c.JSON(http.StatusOK, success)

}
