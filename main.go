package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pasDamola/product-sorter/internal/product"
	"github.com/pasDamola/product-sorter/internal/sorter"
)

// getProducts initializes and returns the sample product data.
func getProducts() []product.Product {
	// A helper function for robust date parsing
	parseDate := func(dateStr string) time.Time {
		t, err := time.Parse("2006-01-04", dateStr)
		if err != nil {
			log.Fatalf("Error parsing date '%s': %v", dateStr, err)
		}
		return t
	}

	return []product.Product{
		{ID: 1, Name: "Alabaster Table", Price: 12.99, Created: parseDate("2019-01-04"), SalesCount: 32, ViewsCount: 730},
		{ID: 2, Name: "Zebra Table", Price: 44.49, Created: parseDate("2012-01-04"), SalesCount: 301, ViewsCount: 3279},
		{ID: 3, Name: "Coffee Table", Price: 10.00, Created: parseDate("2014-05-28"), SalesCount: 1048, ViewsCount: 20123},
	}
}

func main() {
	originalProducts := getProducts()

	fmt.Println("--- Original Unsorted Products ---")
	printProducts(originalProducts)

	// --- Use Case 1: Sort by Price (Ascending) ---
	fmt.Println("\n--- Sorting by Price (Ascending) ---")
	sortedByPrice := sorter.SortProducts(originalProducts, &sorter.PriceSorter{})
	printProducts(sortedByPrice)

	// --- Use Case 2: Sort by Sales per View Ratio (Descending) ---
	fmt.Println("\n--- Sorting by Sales/View Ratio (Descending) ---")
	sortedByRatio := sorter.SortProducts(originalProducts, &sorter.SalesRatioSorter{})
	printProducts(sortedByRatio)

	// --- DEMONSTRATING EXTENSIBILITY ---
	// A new team can add this without touching any of the files above,
	// just by adding a new strategy implementation.
	// Like this below

	// sortedByNewest := sorter.SortProducts(originalProducts, &sorter.NewestProductSorter{})

	fmt.Println("\n--- Verifying Original Slice Unchanged ---")
	fmt.Println("The first product in the original slice is still:")
	fmt.Println(originalProducts[0])
}

// printProducts is a helper to neatly display a slice of products.
func printProducts(products []product.Product) {
	for _, p := range products {
		fmt.Println(p)
	}
}
