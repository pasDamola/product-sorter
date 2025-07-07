package main

import (
	"fmt"
	"time"
)

type Product struct {
	ID         int
	Name       string
	Price      float64
	Created    time.Time
	SalesCount int
	ViewsCount int
}

func (p Product) String() string {
	return fmt.Sprintf("ID: %d, Name: %-15s, Price: %5.2f, Sales/View Ratio: %.4f, Created: %s",
		p.ID, p.Name, p.Price, p.salesPerViewRatio(), p.Created.Format("2006-01-02"))
}

func (p *Product) salesPerViewRatio() float64 {
	if p.ViewsCount == 0 {
		return 0.0
	}
	return float64(p.SalesCount) / float64(p.ViewsCount)
}
