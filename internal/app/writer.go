package app

import (
	"encoding/xml"
	"os"
	"time"
)

type Writer struct {
	fileName string
	file     *os.File
}

type Catalog struct {
	XMLName xml.Name `xml:"catalog"`
	Date    string   `xml:"date,attr"`
	Shop    Shop     `xml:"shop"`
}

type Shop struct {
	XMLName    xml.Name   `xml:"shop"`
	Categories Categories `xml:"categories"`
	Offers     Offers     `xml:"offers"`
}

type Categories struct {
	Categories []Category `xml:"category"`
}

type Category struct {
	ID       int    `xml:"id,attr"`
	ParentID int    `xml:"parentId,attr,omitempty"`
	Name     string `xml:",chardata"`
}

type Offers struct {
	Products []Product `xml:"offers"`
}

type Product struct {
	XMLName      xml.Name `xml:"offer"`
	ID           int      `xml:"id,attr"`
	Name         string   `xml:"name"`
	Price        float64  `xml:"price"`
	OldPrice     *float64 `xml:"oldprice,omitempty"`
	CategoryID   int      `xml:"categoryId"`
	CategoryName string   `xml:"-"`
	Picture      string   `xml:"picture,omitempty"`
	Barcode      string   `xml:"barcode,omitempty"`
}

func NewWriter(fileName string) (w *Writer, err error) {
	w = &Writer{fileName: fileName}
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	w.file = file
	return
}

func (w *Writer) write(categories []Category, products []Product) (err error) {
	encoder := xml.NewEncoder(w.file)
	encoder.Indent("", "  ")
	_, err = w.file.WriteString(xml.Header)
	if err != nil {
		return err
	}

	yml := xml.StartElement{
		Name: xml.Name{Local: "catalog"},
		Attr: []xml.Attr{
			{Name: xml.Name{Local: "date"}, Value: time.Now().Format("2006-01-02 15:04")},
		},
	}
	if err = encoder.EncodeToken(yml); err != nil {
		return err
	}

	if err = encoder.EncodeToken(xml.StartElement{Name: xml.Name{Local: "shop"}}); err != nil {
		return err
	}

	if err = encoder.EncodeToken(xml.StartElement{Name: xml.Name{Local: "categories"}}); err != nil {
		return err
	}
	for _, role := range categories {
		if err = encoder.Encode(role); err != nil {
			return err
		}
	}
	if err = encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "categories"}}); err != nil {
		return err
	}

	if err = encoder.EncodeToken(xml.StartElement{Name: xml.Name{Local: "offers"}}); err != nil {
		return err
	}

	for _, p := range products {
		if err = encoder.EncodeToken(xml.Comment(" " + p.CategoryName + " ")); err != nil {
			return err
		}

		if err = encoder.Encode(p); err != nil {
			return err
		}
	}

	if err = encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "offers"}}); err != nil {
		return err
	}

	if err = encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "shop"}}); err != nil {
		return err
	}
	if err = encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "catalog"}}); err != nil {
		return err
	}

	if err = encoder.Flush(); err != nil {
		return err
	}

	return nil
}

func (w *Writer) close() error {
	if w.file != nil {
		if err := w.file.Close(); err != nil {
			return err
		}
	}
	return nil
}
