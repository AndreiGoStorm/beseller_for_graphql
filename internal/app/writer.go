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
	XMLName xml.Name `xml:"yml_catalog"`
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
	XMLName  xml.Name `xml:"category"`
	ID       int      `xml:"id,attr"`
	ParentID int      `xml:"parentId,attr,omitempty"`
	Name     string   `xml:",chardata"`
}

type Offers struct {
	Products []Product `xml:"offer"`
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
	enc := xml.NewEncoder(w.file)
	enc.Indent("", "  ")
	_, err = w.file.WriteString(xml.Header)
	if err != nil {
		return err
	}

	yml := xml.StartElement{
		Name: xml.Name{Local: "yml_catalog"},
		Attr: []xml.Attr{
			{Name: xml.Name{Local: "date"}, Value: time.Now().Format("2006-01-02 15:04")},
		},
	}
	if err = enc.EncodeToken(yml); err != nil {
		return err
	}

	if err = enc.EncodeToken(xml.StartElement{Name: xml.Name{Local: "shop"}}); err != nil {
		return err
	}

	if err = enc.EncodeToken(xml.StartElement{Name: xml.Name{Local: "categories"}}); err != nil {
		return err
	}
	for _, c := range categories {
		if err = enc.Encode(c); err != nil {
			return err
		}
	}
	if err = enc.EncodeToken(xml.EndElement{Name: xml.Name{Local: "categories"}}); err != nil {
		return err
	}

	if err = enc.EncodeToken(xml.StartElement{Name: xml.Name{Local: "offers"}}); err != nil {
		return err
	}

	for _, p := range products {
		if err = enc.EncodeToken(xml.Comment(" " + p.CategoryName + " ")); err != nil {
			return err
		}

		if err = enc.Encode(p); err != nil {
			return err
		}
	}

	if err = enc.EncodeToken(xml.EndElement{Name: xml.Name{Local: "offers"}}); err != nil {
		return err
	}

	if err = enc.EncodeToken(xml.EndElement{Name: xml.Name{Local: "shop"}}); err != nil {
		return err
	}
	if err = enc.EncodeToken(xml.EndElement{Name: xml.Name{Local: "yml_catalog"}}); err != nil {
		return err
	}

	return enc.Flush()
}

func (w *Writer) close() error {
	if w.file != nil {
		if err := w.file.Close(); err != nil {
			return err
		}
	}
	return nil
}
