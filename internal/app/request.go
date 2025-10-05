package app

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"beseller/internal/helpers"
)

type Request struct {
	url     string
	timeout time.Duration
}

type QueryRequest struct {
	Query string `json:"query"`
}

func NewRequest(appURL, apiURL string) *Request {
	return &Request{helpers.JoinURL(appURL, apiURL), 10 * time.Second}
}

func (r *Request) do() (qr *QueryResponse, err error) {
	jsonData, err := r.buildQuery()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", r.url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &qr); err != nil {
		return nil, err
	}

	return qr, nil
}

func (r *Request) buildQuery() (jsonData []byte, err error) {
	query := `
		query FilterData {
			filterCategory {
				id
				name
				parentCategory {
					additionalInfo {
						categoryId
					}
				}
			}
			filterProduct(filter: { statusId: 1 }) {
				id
				name
				price
				oldPrice
				category {
					id
					name
				}
				images {
					id
					image
				}
			}
		}`

	reqBody := QueryRequest{Query: query}
	jsonData, err = json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
