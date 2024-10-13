package godog_tests

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"bdd-poc/src/calc"
	"github.com/cucumber/godog"
)

type godogsCtxKey struct{}

func thereAreGodogs(ctx context.Context, available int) (context.Context, error) {
	return context.WithValue(ctx, godogsCtxKey{}, available), nil
}

func iEat(ctx context.Context, num int) (context.Context, error) {
	available, ok := ctx.Value(godogsCtxKey{}).(int)
	if !ok {
		return ctx, errors.New("there are no godogs available")
	}

	if available < num {
		return ctx, fmt.Errorf("you cannot eat %d godogs, there are %d available", num, available)
	}

	return context.WithValue(ctx, godogsCtxKey{}, calc.Eat(available, num)), nil
}

func thereShouldBeRemaining(ctx context.Context, remaining int) error {
	available, ok := ctx.Value(godogsCtxKey{}).(int)
	if !ok {
		return errors.New("there are no godogs available")
	}

	if available != remaining {
		return fmt.Errorf("expected %d godogs to be remaining, but there is %d", remaining, available)
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Given(`^there are (\d+) godogs$`, thereAreGodogs)
	ctx.When(`^I eat (\d+)$`, iEat)
	ctx.Then(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
}

func TestFirstScenario(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("failed to run the features test")
	}
}
