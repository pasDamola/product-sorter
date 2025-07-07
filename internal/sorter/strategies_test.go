package sorter

import (
	"testing"

	"github.com/pasDamola/product-sorter/internal/product"
)

// TestPriceSorter verifies ascending sort by price.
func TestPriceSorter(t *testing.T) {
	products := []product.Product{
		{ID: 1, Name: "Mid-Price", Price: 50.00},
		{ID: 2, Name: "Low-Price", Price: 10.00},
		{ID: 3, Name: "High-Price", Price: 100.00},
	}

	sortedProducts := SortProducts(products, &PriceSorter{})

	if sortedProducts[0].Name != "Low-Price" {
		t.Errorf("Expected first product to be 'Low-Price', but got '%s'", sortedProducts[0].Name)
	}
	if sortedProducts[1].Name != "Mid-Price" {
		t.Errorf("Expected second product to be 'Mid-Price', but got '%s'", sortedProducts[1].Name)
	}
	if sortedProducts[2].Name != "High-Price" {
		t.Errorf("Expected last product to be 'High-Price', but got '%s'", sortedProducts[2].Name)
	}
}

// TestSalesRatioSorter verifies descending sort by sales/view ratio.
func TestSalesRatioSorter(t *testing.T) {
	products := []product.Product{
		{ID: 1, Name: "Low-Ratio", SalesCount: 10, ViewsCount: 100},  // Ratio: 0.1
		{ID: 2, Name: "High-Ratio", SalesCount: 50, ViewsCount: 100}, // Ratio: 0.5
		{ID: 3, Name: "Mid-Ratio", SalesCount: 20, ViewsCount: 100},  // Ratio: 0.2
	}

	sortedProducts := SortProducts(products, &SalesRatioSorter{})

	if sortedProducts[0].Name != "High-Ratio" {
		t.Errorf("Expected first product to be 'High-Ratio', but got '%s'", sortedProducts[0].Name)
	}
	if sortedProducts[2].Name != "Low-Ratio" {
		t.Errorf("Expected last product to be 'Low-Ratio', but got '%s'", sortedProducts[2].Name)
	}
}

// TestImmutability ensures the original slice is not modified.
func TestImmutability(t *testing.T) {
	originalProducts := []product.Product{
		{ID: 1, Name: "A"},
		{ID: 2, Name: "B"},
	}
	originalFirstItemName := originalProducts[0].Name

	// Create a copy to compare against later
	originalCopy := make([]product.Product, len(originalProducts))
	copy(originalCopy, originalProducts)

	// Perform a sort
	SortProducts(originalProducts, &PriceSorter{})

	// Check if the original slice has been modified
	if originalProducts[0].Name != originalFirstItemName {
		t.Errorf("Immutability test failed: The original slice was modified.")
	}
}
