// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package http

import (
	"time"
)

const (
	Api_keyScopes = "api_key.Scopes"
)

// Defines values for LinkRel.
const (
	Content LinkRel = "content"
	Self    LinkRel = "self"
)

// Defines values for GetInventoryParamsSort.
const (
	Author       GetInventoryParamsSort = "author"
	Manufacturer GetInventoryParamsSort = "manufacturer"
	Mpn          GetInventoryParamsSort = "mpn"
)

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Detail   *string `json:"detail,omitempty"`
	Instance *string `json:"instance,omitempty"`
	Status   int     `json:"status"`
	Title    string  `json:"title"`
	Type     *string `json:"type,omitempty"`
}

// InventoryEntry defines model for InventoryEntry.
type InventoryEntry struct {
	Links              *[]Link                 `json:"links,omitempty"`
	SchemaAuthor       SchemaAuthor            `json:"schema:author"`
	SchemaManufacturer SchemaManufacturer      `json:"schema:manufacturer"`
	SchemaMpn          string                  `json:"schema:mpn"`
	Versions           []InventoryEntryVersion `json:"versions"`
}

// InventoryEntryResponse defines model for InventoryEntryResponse.
type InventoryEntryResponse struct {
	Data InventoryEntry `json:"data"`
}

// InventoryEntryVersion defines model for InventoryEntryVersion.
type InventoryEntryVersion struct {
	Description string       `json:"description"`
	Links       *[]Link      `json:"links,omitempty"`
	Original    string       `json:"original"`
	Timestamp   *string      `json:"timestamp,omitempty"`
	TmId        string       `json:"tmId"`
	Version     ModelVersion `json:"version"`
}

// InventoryEntryVersionsResponse defines model for InventoryEntryVersionsResponse.
type InventoryEntryVersionsResponse struct {
	Data []InventoryEntryVersion `json:"data"`
}

// InventoryResponse defines model for InventoryResponse.
type InventoryResponse struct {
	Data map[string]InventoryEntry `json:"data"`
	Meta *Meta                     `json:"meta,omitempty"`
}

// Link defines model for Link.
type Link struct {
	Href string  `json:"href"`
	Rel  LinkRel `json:"rel"`
}

// LinkRel defines model for Link.Rel.
type LinkRel string

// Meta defines model for Meta.
type Meta struct {
	Created time.Time `json:"created"`
}

// ModelVersion defines model for ModelVersion.
type ModelVersion struct {
	Model string `json:"model"`
}

// SchemaAuthor defines model for SchemaAuthor.
type SchemaAuthor struct {
	SchemaName string `json:"schema:name"`
}

// SchemaManufacturer defines model for SchemaManufacturer.
type SchemaManufacturer struct {
	SchemaName string `json:"schema:name"`
}

// GetInventoryParams defines parameters for GetInventory.
type GetInventoryParams struct {
	// FilterAuthor Filters the inventory by one or more authors having exact match. The filter works additive to other filters.
	FilterAuthor *string `form:"filter[author],omitempty" json:"filter[author],omitempty"`

	// FilterManufacturer Filters the inventory by one or more manufacturers having exact match. The filter works additive to other filters.
	FilterManufacturer *string `form:"filter[manufacturer],omitempty" json:"filter[manufacturer],omitempty"`

	// FilterMpn Filters the inventory by one ore more mpn (manufacturer part number) having exact match. The filter works additive to other filters.
	FilterMpn *string `form:"filter[mpn],omitempty" json:"filter[mpn],omitempty"`

	// FilterOriginal Filters the inventory by one or more original ID having exact match. The filter works additive to other filters.
	FilterOriginal *string `form:"filter[original],omitempty" json:"filter[original],omitempty"`

	// SearchContent Filters the inventory by content search with the help of an awesome query language
	SearchContent *string `form:"search[content],omitempty" json:"search[content],omitempty"`

	// Sort Sorts the inventory by one or more fields. The sort is applied in the order of the fields.  The sorting is done ascending per field by default. If a field needs to be sorted descending, prefix it with a HYPHEN-MINUS "-")
	Sort *GetInventoryParamsSort `form:"sort,omitempty" json:"sort,omitempty"`
}

// GetInventoryParamsSort defines parameters for GetInventory.
type GetInventoryParamsSort string
