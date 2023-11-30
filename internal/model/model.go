package model

import "github.com/web-of-things-open-source/tm-catalog-cli/internal"

// ThingModel is a model for unmarshalling a Thing Model to be
// imported. It contains only the fields required to be accepted into
// the catalog.
type ThingModel struct {
	Manufacturer SchemaManufacturer `json:"schema:manufacturer" validate:"required"`
	Mpn          string             `json:"schema:mpn" validate:"required"`
	Author       SchemaAuthor       `json:"schema:author" validate:"required"`
	Version      Version            `json:"version"`
}

func (tm *ThingModel) IsOfficial() bool {
	compareMan := internal.ToTrimmedLower(tm.Manufacturer.Name)
	compareAuthor := internal.ToTrimmedLower(tm.Author.Name)
	return compareMan == compareAuthor
}

type SchemaAuthor struct {
	Name string `json:"schema:name" validate:"required"`
}
type SchemaManufacturer struct {
	Name string `json:"schema:name" validate:"required"`
}

type Version struct {
	Model string `json:"model"`
}

type ExtendedFields struct {
	Links `json:"links"`
	// TODO: why is ID field not in ThingModel? It is also needed for importing to be moved to "original"
	ID          string `json:"id,omitempty"`
	Description string `json:"description"`
}

type CatalogThingModel struct {
	ThingModel
	ExtendedFields
}
