package main

// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	"testing"

// 	"github.com/cucumber/godog"
// 	"github.com/cucumber/godog/colors"
// )

// var opts = godog.Options{Output: colors.Colored(os.Stdout)}

// func init() {
// 	godog.BindFlags("godog.", flag.CommandLine, &opts)
// }

// func TestFeatures2(t *testing.T) {
// 	o := opts
// 	o.TestingT = t

// 	status := godog.TestSuite{
// 		Name:                 "godogs",
// 		Options:              &o,
// 		TestSuiteInitializer: InitializeTestSuite,
// 		ScenarioInitializer:  InitializeScenario,
// 	}.Run()

// 	if status == 2 {
// 		t.SkipNow()
// 	}

// 	if status != 0 {
// 		t.Fatalf("zero status code expected, %d received", status)
// 	}
// }

// func FeatureContext(s *godog.TestSuite) {
// 	cart := CreateCart()

// 	s.Step(`^an empty cart$`, func() {
// 		cart = CreateCart()
// 	})

// 	s.Step(`^I add the item "([^"]*)" with price (\d+) and quantity (\d+)$`, func(nameInp string, priceInp, quantityInp int) {
// 		item := Item{name: nameInp, price: priceInp, quantity: quantityInp}
// 		cart.AddItem(item)
// 	})

// 	s.Step(`^the cart should contain (\d+) items$`, func(expectedCount int) {
// 		if len(cart.items) != expectedCount {
// 			panic(fmt.Sprintf("Expected %d items in cart, but got %d", expectedCount, len(cart.items)))
// 		}
// 	})

// 	s.Step(`^the total price of the cart should be (\d+)$`, func(expectedTotal int) {
// 		totalPrice := 0
// 		for _, item := range cart.items {
// 			totalPrice += item.price * item.quantity
// 		}
// 		if totalPrice != expectedTotal {
// 			panic(fmt.Sprintf("Expected total price of %d, but got %d", expectedTotal, totalPrice))
// 		}
// 	})

// 	s.Step(`^the cart status should be "([^"]*)"$`, func(expectedStatus string) {
// 		if cart.status != expectedStatus {
// 			panic(fmt.Sprintf("Expected cart status %s, but got %s", expectedStatus, cart.status))
// 		}
// 	})
// }
