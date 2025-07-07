package sorter

import "github.com/pasDamola/product-sorter/internal/product"

// SortingStrategy defines the contract for any comparison logic.
// This is the core of the Strategy Pattern. Any new sorter
// must implement this interface.
type SortingStrategy interface {
	Less(p1, p2 *product.Product) bool
}

// --- Concrete Strategy Implementations ---

// PriceSorter sorts products by price in ascending order.
type PriceSorter struct{}

func (s *PriceSorter) Less(p1, p2 *product.Product) bool {
	return p1.Price < p2.Price
}

// SalesRatioSorter sorts products by their sales/view ratio in descending order.
type SalesRatioSorter struct{}

func (s *SalesRatioSorter) Less(p1, p2 *product.Product) bool {
	// A higher ratio is better, so we use > for descending sort.
	return p1.SalesPerViewRatio() > p2.SalesPerViewRatio()
}

// A NewtProductSorter can be added here. For example that sorts products by creation date, newest first.
type NewProductSorter struct{}

func (s *NewProductSorter) Less(p1, p2 *product.Product) bool {
	// Use After for descending chronological order (newest first)
	return p1.Created.After(p2.Created)
}
