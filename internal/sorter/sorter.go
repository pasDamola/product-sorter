// product-sorter/internal/sorter/sorter.go
package sorter

import (
	"sort"

	"github.com/pasDamola/product-sorter/internal/product"
)

// SortProducts applies any sorting strategy of our specification to a slice of products and returns a
// NEW sorted slice, leaving the original data untouched.
func SortProducts(products []product.Product, strategy SortingStrategy) []product.Product {
	// 1. Create a defensive copy to protect the original data.
	productsCopy := make([]product.Product, len(products))
	copy(productsCopy, products)

	// 2. Create an instance of our private sorter helper type.
	s := &productSorter{
		products: productsCopy,
		strategy: strategy,
	}

	// 3. Sort the copy using Go's standard sort library.
	sort.Sort(s)

	// 4. Return the newly sorted copy.
	return s.products
}

// productSorter is an unexported helper type that implements sort.Interface.
// It's used internally by SortProducts to perform the sort.
type productSorter struct {
	products []product.Product
	strategy SortingStrategy
}

// The Len, Swap and Less methods are part of sort.Interface and we implement the interface on the productSorter

func (s *productSorter) Len() int {
	return len(s.products)
}

func (s *productSorter) Swap(i, j int) {
	s.products[i], s.products[j] = s.products[j], s.products[i]
}

func (s *productSorter) Less(i, j int) bool {
	return s.strategy.Less(&s.products[i], &s.products[j])
}
