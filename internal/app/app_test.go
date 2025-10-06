package app

import (
	"testing"

	"beseller/internal/config"
	"beseller/internal/helpers"
	"github.com/stretchr/testify/assert"
)

func TestHandleCategories(t *testing.T) {
	tests := []struct {
		name     string
		input    []GraphCategory
		expected []Category
	}{
		{
			name: "ParentID is present",
			input: []GraphCategory{
				{
					ID:             10,
					Name:           "Category",
					ParentCategory: &ParentCategory{&AdditionalInfo{CategoryID: 1}},
				},
			},
			expected: []Category{
				{ID: 10, Name: "Category", ParentID: 1},
			},
		},
		{
			name: "Multiple categories",
			input: []GraphCategory{
				{
					ID:             1,
					Name:           "Main",
					ParentCategory: &ParentCategory{&AdditionalInfo{CategoryID: 1}},
				},
				{
					ID:             5,
					Name:           "Child1",
					ParentCategory: &ParentCategory{&AdditionalInfo{CategoryID: 1}},
				},
				{
					ID:             10,
					Name:           "Child2",
					ParentCategory: &ParentCategory{&AdditionalInfo{CategoryID: 1}},
				},
				{
					ID:             100,
					Name:           "ChildNext",
					ParentCategory: &ParentCategory{&AdditionalInfo{CategoryID: 10}},
				},
			},
			expected: []Category{
				{ID: 1, Name: "Main", ParentID: 0},
				{ID: 5, Name: "Child1", ParentID: 1},
				{ID: 10, Name: "Child2", ParentID: 1},
				{ID: 100, Name: "ChildNext", ParentID: 10},
			},
		},
		{
			name: "ParentID is zero",
			input: []GraphCategory{
				{
					ID:             10,
					Name:           "Автомагнитолы",
					ParentCategory: &ParentCategory{&AdditionalInfo{CategoryID: 10}},
				},
			},
			expected: []Category{
				{ID: 10, Name: "Автомагнитолы", ParentID: 0},
			},
		},

		{
			name: "Parent category is nil",
			input: []GraphCategory{
				{
					ID:             100,
					Name:           "Велосипеды",
					ParentCategory: nil,
				},
			},
			expected: []Category{
				{ID: 100, Name: "Велосипеды", ParentID: 0},
			},
		},
		{
			name: "Additional info is nil",
			input: []GraphCategory{
				{
					ID:             6,
					Name:           "Самокаты",
					ParentCategory: &ParentCategory{nil},
				},
			},
			expected: []Category{
				{ID: 6, Name: "Самокаты", ParentID: 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := NewApp(&config.Config{})

			app.handleCategories(tt.input)

			assert.Equal(t, tt.expected, app.categories, "categories should match")
		})
	}
}

func TestHandleProducts(t *testing.T) {
	oldPrice := 150.0
	tests := []struct {
		name     string
		input    []GraphProduct
		expected []Product
	}{
		{
			name: "Product with category",
			input: []GraphProduct{
				{
					ID:       1,
					Name:     "Liebherr IKB 2320",
					Price:    250,
					OldPrice: &oldPrice,
					ProductCategory: &ProductCategory{
						ID:   10,
						Name: "Category A",
					},
					Images: []Image{
						{ID: 1, Image: "fake.jpg"},
					},
				},
			},
			expected: []Product{
				{
					ID:           1,
					Name:         "Liebherr IKB 2320",
					Price:        250,
					OldPrice:     &oldPrice,
					CategoryID:   10,
					CategoryName: "Category A",
					Picture:      "http://app.com/pics/items/fake.jpg",
				},
			},
		},
		{
			name: "Product with Images",
			input: []GraphProduct{
				{
					ID:       5,
					Name:     "IKB 500",
					Price:    200,
					OldPrice: &oldPrice,
					Images: []Image{
						{ID: 1, Image: "fake1.jpg"},
						{ID: 2, Image: "fake2.jpg"},
					},
				},
			},
			expected: []Product{
				{
					ID:       5,
					Name:     "IKB 500",
					Price:    200,
					OldPrice: &oldPrice,
					Picture:  "http://app.com/pics/items/fake1.jpg",
				},
			},
		},
		{
			name: "Product without category, images",
			input: []GraphProduct{
				{
					ID:       1,
					Name:     "Electric MR",
					Price:    100,
					OldPrice: &oldPrice,
				},
			},
			expected: []Product{
				{
					ID:       1,
					Name:     "Electric MR",
					Price:    100,
					OldPrice: &oldPrice,
				},
			},
		},
	}

	conf := helpers.NewTestConfig()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := NewApp(conf)

			app.handleProducts(tt.input)

			assert.Equal(t, tt.expected, app.products, "products should match")
		})
	}
}
