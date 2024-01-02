package pkg

import (
	"fmt"
)

func Print(baseCurrency string, value float32, currencies map[string]float32, verbose bool) {
	if verbose {
		for k, v := range currencies {
			fmt.Printf("%f %s converted to %s is equal to %f\n", value, baseCurrency, k, v)
		}
	} else {
		fmt.Printf("%s: %f\n", baseCurrency, value)
		for k, v := range currencies {
			fmt.Printf("%s: %f\n", k, v)
		}
	}
}
