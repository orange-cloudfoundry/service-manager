package types

type Plans struct {
	Plans []*Plan `json:"service_plans"`
}

type Plan struct {
	CatalogID string `json:"catalog_id"`
}
