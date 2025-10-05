package app

import (
	"beseller/internal/config"
	"beseller/internal/helpers"
)

type App struct {
	name string
	file string
	conf *config.Config

	req    *Request
	writer *Writer

	categories []Category
	products   []Product
}

func NewApp(conf *config.Config) *App {
	a := &App{name: conf.Name, file: conf.File, conf: conf}
	a.req = NewRequest(conf.AppURL, conf.APIURL)
	return a
}

func (a *App) DoRequest() (err error) {
	qr, err := a.req.do()
	if err != nil {
		return err
	}
	a.handleCategories(qr.Data.FilterCategory)
	a.handleProducts(qr.Data.FilterProduct)
	return nil
}

func (a *App) handleCategories(categories []GraphCategory) {
	a.categories = make([]Category, 0, len(categories))
	for _, c := range categories {
		category := Category{ID: c.ID, Name: c.Name}
		if pc := c.ParentCategory; pc != nil {
			if ai := pc.AdditionalInfo; ai != nil && c.ID != ai.CategoryID {
				category.ParentID = c.ParentCategory.AdditionalInfo.CategoryID
			}
		}
		a.categories = append(a.categories, category)
	}
}

func (a *App) handleProducts(products []GraphProduct) {
	a.products = make([]Product, 0, len(products))
	for _, p := range products {
		product := Product{ID: p.ID, Name: p.Name, Price: p.Price, OldPrice: p.OldPrice}
		if pc := p.ProductCategory; pc != nil {
			product.CategoryID = p.ProductCategory.ID
			product.CategoryName = p.ProductCategory.Name
		}
		if len(p.Images) > 0 {
			product.Picture = helpers.JoinURL(a.conf.AppURL, helpers.JoinURL(a.conf.ImageURL, p.Images[0].Image))
		}
		a.products = append(a.products, product)
	}
}

func (a *App) Write() (err error) {
	a.writer, err = NewWriter(a.file)
	if err != nil {
		return err
	}

	err = a.writer.write(a.categories, a.products)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) Close() error {
	return a.writer.close()
}
