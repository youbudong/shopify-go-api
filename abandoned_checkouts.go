package shopify

import (
	"fmt"
	"time"
)

const (
	AbandonedCheckoutsBasePath     = "checkouts"
	AbandonedCheckoutsResourceName = "checkouts"
)

// AbandonedCheckoutsService is an interface for interfacing with the abandoned checkouts endpoints of
// the Shopify API.
// See: https://shopify.dev/api/admin-rest/2021-10/resources/abandoned-checkouts
type AbandonedCheckoutsService interface {
	List(interface{}) ([]AbandonedCheckouts, error)
}

// AbandonedCheckoutsServiceOp handles communication with the draft order related methods of the
// Shopify API.
type AbandonedCheckoutsServiceOp struct {
	client *Client
}

// AbandonedCheckouts represents a shopify draft order
type AbandonedCheckouts struct {
	ID                       int64          `json:"id,omitempty"`
	AbandonedCheckoutUrl     string         `json:"abandoned_checkout_url,omitempty"`
	BillingAddress           *Address       `json:"billing_address,omitempty"`
	BuyerAcceptsMarketing    bool           `json:"buyer_accepts_marketing,omitempty"`
	BuyerAcceptsSmsMarketing bool           `json:"buyer_accepts_sms_marketing,omitempty"`
	CartToken                string         `json:"cart_token,omitempty"`
	Customer                 *Customer      `json:"customer,omitempty"`
	CustomerLocale           string         `json:"customer_locale,omitempty"`
	DeviceID                 int64          `json:"device_id,omitempty"`
	DiscountCodes            []DiscountCode `json:"discount_codes,omitempty"`
	Email                    string         `json:"email,omitempty"`
	Gateway                  string         `json:"gateway,omitempty"`
	LandingSite              string         `json:"landing_site,omitempty"`
	LineItems                []LineItem     `json:"line_items,omitempty"`
	LocationId               int64          `json:"location_id,omitempty"`
	Note                     string         `json:"note,omitempty"`
	Phone                    string         `json:"phone,omitempty"`
	ReferringSite            string         `json:"referring_site,omitempty"`
	ShippingAddress          *Address       `json:"shipping_address,omitempty"`
	SourceName               string         `json:"source_name,omitempty"`
	SubtotalPrice            string         `json:"subtotal_price,omitempty"`
	TaxLines                 []TaxLine      `json:"tax_lines,omitempty"`
	TaxesIncluded            bool           `json:"taxes_included,omitempty"`
	Token                    string         `json:"token,omitempty"`
	TotalDiscounts           string         `json:"total_discounts,omitempty"`
	TotalLineItemsPrice      string         `json:"total_line_items_price,omitempty"`
	TotalPrice               string         `json:"total_price,omitempty"`
	TotalTax                 string         `json:"total_tax,omitempty"`
	TotalWeight              int            `json:"total_weight,omitempty"`
	UserId                   int64          `json:"user_id,omitempty"`
	Currency                 string         `json:"currency,omitempty"`
	ClosedAt                 *time.Time     `json:"closed_at,omitempty"`
	CompletedAt              *time.Time     `json:"completed_at,omitempty"`
	CreatedAt                *time.Time     `json:"created_at,omitempty"`
	UpdatedAt                *time.Time     `json:"updated_at,omitempty"`
}

type AbandonedCheckoutsesResource struct {
	AbandonedCheckoutses []AbandonedCheckouts `json:"checkouts"`
}

// AbandonedCheckoutsListOptions represents the possible options that can be used
// to further query the list abandoned checkouts endpoint
type AbandonedCheckoutsListOptions struct {
	Limit        int        `url:"limit,omitempty"`
	SinceID      int64      `url:"since_id,omitempty"`
	UpdatedAtMin *time.Time `url:"updated_at_min,omitempty"`
	UpdatedAtMax *time.Time `url:"updated_at_max,omitempty"`
	CreatedAtMin *time.Time `url:"created_at_min,omitempty"`
	CreatedAtMax *time.Time `url:"created_at_max,omitempty"`
	Status       string     `url:"status,omitempty"`
}

// List abandoned checkouts
func (s *AbandonedCheckoutsServiceOp) List(options interface{}) ([]AbandonedCheckouts, error) {
	path := fmt.Sprintf("%s.json", AbandonedCheckoutsBasePath)
	resource := new(AbandonedCheckoutsesResource)
	fmt.Println(path)
	err := s.client.Get(path, resource, options)
	return resource.AbandonedCheckoutses, err
}
