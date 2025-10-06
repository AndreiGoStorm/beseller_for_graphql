package app

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var fileName = "export.yml"

func TestWriteYml(t *testing.T) {
	wr, err := NewWriter(fileName)
	require.NoError(t, err)

	categories := getCategories()
	products := getProducts()

	t.Run("write xml file", func(t *testing.T) {
		err = wr.write(categories, products)
		require.NoError(t, wr.close())

		data, err := os.ReadFile(fileName)
		require.NoError(t, err)

		var yml Catalog
		err = xml.Unmarshal(data, &yml)
		require.NoError(t, err)

		assert.Len(t, yml.Shop.Categories.Categories, 3)
		for i, c := range yml.Shop.Categories.Categories {
			assert.Equal(t, categories[i].ID, c.ID)
			assert.Equal(t, categories[i].ParentID, c.ParentID)
			assert.Equal(t, categories[i].Name, c.Name)
		}

		assert.Len(t, yml.Shop.Offers.Products, 3)
		for i, p := range yml.Shop.Offers.Products {
			assert.Equal(t, products[i].ID, p.ID)
			assert.Equal(t, products[i].Name, p.Name)
			assert.Equal(t, products[i].Price, p.Price)
			assert.Equal(t, products[i].OldPrice, p.OldPrice)
			assert.Equal(t, products[i].CategoryID, p.CategoryID)
			assert.Equal(t, products[i].Picture, p.Picture)
			assert.Equal(t, products[i].Barcode, p.Barcode)
		}

		require.NoError(t, os.Remove(fileName))
	})
}

func getCategories() []Category {
	return []Category{
		{
			ID:       1,
			ParentID: 0,
			Name:     "Main",
		},
		{
			ID:       5,
			ParentID: 1,
			Name:     "Category A",
		},
		{
			ID:       10,
			ParentID: 1,
			Name:     "Category B",
		},
	}
}

func getProducts() []Product {
	oldPrice := 200.0
	return []Product{
		{
			ID:           1,
			Name:         "Аккумулятор Exide",
			Price:        250,
			OldPrice:     &oldPrice,
			CategoryID:   1,
			CategoryName: "Category A",
			Picture:      "http://app.com/pics/items/fake.jpg",
			Barcode:      "2000053473227",
		},
		{
			ID:      3,
			Name:    "IKB 500",
			Price:   200,
			Picture: "http://app.com/pics/items/fake1.png",
			Barcode: "2000053415331",
		},
		{
			ID:           5,
			Name:         "Electric MR",
			Price:        100,
			OldPrice:     &oldPrice,
			CategoryID:   10,
			CategoryName: "Category B",
			Barcode:      "2000053413448",
		},
	}
}
