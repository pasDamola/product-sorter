# Product Catalog Sorting Solution

This project is a Go-based solution to the "Product Catalog Sorting" code challenge. It demonstrates a clean, extensible, and robust implementation designed according to modern software engineering principles.

## Core Requirements

The primary goal is to sort a slice of products by different criteria (e.g., price, sales-per-view ratio). The key constraint is that the system must be **extensible**, allowing new sorting strategies to be added in the future without modifying existing code.

## Design Approach

To meet the extensibility requirement, this solution is built upon two core concepts:

1.  **The Strategy Design Pattern**: The logic for each sorting method is encapsulated in its own "strategy" struct. This decouples the main sorting function from the specific comparison logic.
2.  **Immutability**: The primary sorting function `sorter.SortProducts` does **not** modify the original slice. Instead, it returns a new, sorted slice. This promotes safety and predictability, preventing accidental data mutation of the original data source.

This design adheres to SOLID principles, particularly:
*   **Open/Closed Principle**: The system is open to new sorting strategies but closed for modification of the core sorting logic.
*   **Single Responsibility Principle**: Each strategy has only one job: to compare two products based on a single criterion.

## Project Structure

```
product-sorter/
├── .gitignore
├── README.md
├── go.mod
├── main.go
└── internal/
    ├── product/
    │   └── product.go        # Defines the Product model
    └── sorter/
        ├── sorter.go         # Contains the main SortProducts function
        ├── strategies.go     # Defines the SortingStrategy interface and concrete implementations
        └── strategies_test.go # Unit tests for the strategies
```
The core logic is placed within the `internal` directory to demonstrate good package encapsulation, making it private to this module.

## How to Run

### Prerequisites
- Go 1.18 or later

### Running the Application
1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/product-sorter.git
   cd product-sorter
   ```
2. Run the main application to see a demonstration of the different sorting strategies:
   ```sh
   go run main.go
   ```

### Running Tests
Unit tests have been written to verify the correctness of each sorting strategy.
```sh
go test ./... -v
```
The `-v` flag provides verbose output.

<!-- ## Future Improvements

While this solution effectively solves the challenge, a production-grade system could be enhanced with:

1.  **Sorter Factory**: A factory function could be introduced to create sorting strategies from a string input (e.g., from an API query parameter `?sort_by=price`), further decoupling the client code.
2.  **Configurable Sort Direction**: Strategies could be enhanced to accept a direction (e.g., `ASC`, `DESC`) during instantiation for greater flexibility.
3.  **Database-Side Sorting**: For very large datasets, the sorting logic could be moved to the database query (`ORDER BY`) for optimal performance, with the Go code responsible for building the correct query. -->

---