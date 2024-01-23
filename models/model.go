package models

import "github.com/ucanbaklava/webassembly/carrier"

type Model struct {
	ID           string        `json:"id"`
	UserID       string        `json:"user_id"`
	ShipmentID   string        `json:"shipment_id"`
	OrderID      string        `json:"order_id"`
	OrderDate    string        `json:"order_date"`
	LabelDate    string        `json:"label_date"`
	TrackingID   string        `json:"tracking_id"`
	CarrierID    string        `json:"carrier_id"`
	Note         string        `json:"note"`
	SalesChannel SalesChannel  `json:"sales_channel"`
	Content      string        `json:"content"`
	Sender       Sender        `json:"sender"`
	Recipient    Recipient     `json:"recipient"`
	Settings     LabelSettings `json:"settings"`
	Items        []Item        `json:"items"`
	Payment      Payment       `json:"payment"`
	Label        Label         `json:"label"`
}

type Label struct {
	ID  string `json:"id"`
	PDF string `json:"pdf"`
	PNG string `json:"png"`
	ZPL string `json:"zpl"`
}

type LabelSettings struct {
	Size string       `json:"size"`
	Show ShowSettings `json:"show"`
}

type ShowSettings struct {
	Items         bool `json:"items"`
	ItemImage     bool `json:"item_image"`
	ItemPrice     bool `json:"item_price"`
	SenderPhone   bool `json:"sender_phone"`
	SenderAddress bool `json:"sender_address"`
}

type Payment struct {
	Currency carrier.Currency `json:"currency"`
	Amount   float64          `json:"amount"`
	Type     carrier.PType    `json:"type"`
}

type SalesChannel struct {
	Name        string `json:"name"`
	OrderNumber string `json:"order_number"`
	ID          string `json:"id"`
	URL         string `json:"url"`
	Logo        string `json:"logo"`
}

type Item struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	SKU      string  `json:"sku"`
	Image    string  `json:"image"`
}

type Sender struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Zip     string `json:"zip"`
}

type Recipient struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Zip     string `json:"zip"`
}
