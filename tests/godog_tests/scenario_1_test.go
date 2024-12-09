package godog_tests

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"design-patterns/src/calc"

	"github.com/cucumber/godog"
)

func thereAreGodogs(ctx context.Context, available int) (context.Context, error) {
	return context.WithValue(ctx, CtxKey{}, available), nil
}

func iEat(ctx context.Context, num int) (context.Context, error) {
	available, ok := ctx.Value(CtxKey{}).(int)
	if !ok {
		return ctx, errors.New("there are no godogs available")
	}

	if available < num {
		return ctx, fmt.Errorf("you cannot eat %d godogs, there are %d available", num, available)
	}

	return context.WithValue(ctx, CtxKey{}, calc.Eat(available, num)), nil
}

func thereShouldBeRemaining(ctx context.Context, remaining int) error {
	available, ok := ctx.Value(CtxKey{}).(int)
	if !ok {
		return errors.New("there are no godogs available")
	}

	if available != remaining {
		return fmt.Errorf("expected %d godogs to be remaining, but there is %d", remaining, available)
	}

	return nil
}

func InitializeEatScenario(ctx *godog.ScenarioContext) {
	ctx.Given(`^there are (\d+) godogs$`, thereAreGodogs)
	ctx.When(`^I eat (\d+)$`, iEat)
	ctx.Then(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
}

func TestEatScenario(t *testing.T) {
	suite := godog.TestSuite{
		Name:                "eat godogs",
		ScenarioInitializer: InitializeEatScenario,
		Options:             &Options,
	}

	if suite.Run() != 0 {
		t.Fatal("failed to run the features test")
	}
}
