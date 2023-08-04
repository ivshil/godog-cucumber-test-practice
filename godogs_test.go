package main

// This example shows how to set up test suite runner with Go subtests and godog command line parameters.
// Sample commands:
// * run all scenarios from default directory (features): go test -test.run "^TestFeatures/"
// * run all scenarios and list subtest names: go test -test.v -test.run "^TestFeatures/"
// * run all scenarios from one feature file: go test -test.v -godog.paths features/nodogs.feature -test.run "^TestFeatures/"
// * run all scenarios from multiple feature files: go test -test.v -godog.paths features/nodogs.feature,features/godogs.feature -test.run "^TestFeatures/"
// * run single scenario as a subtest: go test -test.v -test.run "^TestFeatures/Eat_5_out_of_12$"
// * show usage help: go test -godog.help
// * show usage help if there were other test files in directory: go test -godog.help godogs_test.go
// * run scenarios with multiple formatters: go test -test.v -godog.format cucumber:cuc.json,pretty -test.run "^TestFeatures/"

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/cucumber/gherkin-go/v19"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/rdumont/assistdog"

)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}
var cart Cart

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestFeatures(t *testing.T) {
	o := opts
	o.TestingT = t

	status := godog.TestSuite{
		Name:                 "godogs",
		Options:              &o,
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
	}.Run()

	if status == 2 {
		t.SkipNow()
	}

	if status != 0 {
		t.Fatalf("zero status code expected, %d received", status)
	}
}

// func thereAreItems(items []Item) error {
// 	for _, i := range items {
// 		ItemsP.Add(i)
// 	}
// }

func thereAreGodogs(available int) error {
	Godogs = available
	return nil
}

func iEat(num int) error {
	if Godogs < num {
		return fmt.Errorf("you cannot eat %d godogs, there are %d available", num, Godogs)
	}
	Godogs -= num
	return nil
}

func thereShouldBeRemaining(remaining int) error {
	if Godogs != remaining {
		return fmt.Errorf("expected %d godogs to be remaining, but there is %d", remaining, Godogs)
	}
	return nil
}

func thereShouldBeNoneRemaining() error {
	return thereShouldBeRemaining(0)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() { Godogs = 0 })
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		Godogs = 0 // clean the state before every scenario
		return ctx, nil
	})

	ctx.Step(`^there are (\d+) godogs$`, thereAreGodogs)
	ctx.Step(`^I eat (\d+)$`, iEat)
	ctx.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
	ctx.Step(`^there should be none remaining$`, thereShouldBeNoneRemaining)

	cart := CreateCart()
	itemPool := make(map[string]Item)

	ctx.Step(`^an empty cart$`, anEmptyCart)
	ctx.Step(`^the following items are available in the pool:$`, theFollowingItemsAreAvailableInThePool)
	ctx.Step(`^I add the item "([^"]*)" with quantity (\d+)$`, iAddTheItemWithQuantity)
	ctx.Step(`^the cart should contain (\d+) items$`, theCartShouldContainItems)
	ctx.Step(`^the total price of the cart should be (\d+)$`, theTotalPriceOfTheCartShouldBe)

	anEmptyCartP(&cart)
	theFollowingItemsAreAvailableInThePool(itemPool)
}

func anEmptyCart() {
	cart = CreateCart()
}

func anEmptyCartP(cart *Cart) {
	*cart = CreateCart()
}


func theFollowingItemsAreAvailableInThePool(table assistdog.CreateSlice) error {
	for _, row := range table.Rows {
		nameP := row.Cells[0].Value
		priceP, _ := strconv.Atoi(row.Cells[1].Value)
		quantityP, _ := strconv.Atoi(row.Cells[2].Value)
		ItemsP[name] = Item{name: nameP, price: priceP, quantity: quantityP}
	}
	return nil
}

func iAddTheItemWithQuantity(name string, quantity int) error {
	item, ok := ItemsP[name]
	if !ok {
		return fmt.Errorf("item not found in pool: %s", name)
	}
	item.Quantity = quantity
	cart.AddItem(item)
	return nil
}

func theCartShouldContainItems(expectedCount int) error {
	if len(cart.items) != expectedCount {
		return fmt.Errorf("expected %d items in cart, but got %d", expectedCount, len(cart.items))
	}
	return nil
}

func theTotalPriceOfTheCartShouldBe(expectedTotal int) error {
	totalPrice := 0
	for _, item := range cart.items {
		totalPrice += item.price * item.quantity
	}
	if totalPrice != expectedTotal {
		return fmt.Errorf("expected total price of %d, but got %d", expectedTotal, totalPrice)
	}
	return nil
}
