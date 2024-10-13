package godog_tests

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"bdd-poc/src/calc"

	"github.com/cucumber/godog"
)

func theStartNumber(ctx context.Context, available int) (context.Context, error) {
	return context.WithValue(ctx, CtxKey{}, available), nil
}

func followANumberAs(ctx context.Context, num int) (context.Context, error) {
	available, ok := ctx.Value(CtxKey{}).(int)
	if !ok {
		return ctx, errors.New("there are no godogs available")
	}

	if available < num {
		return ctx, fmt.Errorf("you cannot eat %d godogs, there are %d available", num, available)
	}

	return context.WithValue(ctx, CtxKey{}, calc.Generate(available, num)), nil
}

func theResultShouldBe(ctx context.Context, remaining int) error {
	available, ok := ctx.Value(CtxKey{}).(int)
	if !ok {
		return errors.New("there are no godogs available")
	}

	if available != remaining {
		return fmt.Errorf("expected %d godogs to be remaining, but there is %d", remaining, available)
	}

	return nil
}

func InitializeGenScenario(ctx *godog.ScenarioContext) {
	ctx.Given(`^the start number (\d+)$`, theStartNumber)
	ctx.When(`^follow a number as (\d+)$`, followANumberAs)
	ctx.Then(`^the result should be (\d+)$`, theResultShouldBe)
}

func TestGenScenario(t *testing.T) {
	suite := godog.TestSuite{
		Name:                "generate",
		ScenarioInitializer: InitializeGenScenario,
		Options:             &Options,
	}

	if suite.Run() != 0 {
		t.Fatal("failed to run the features test")
	}
}
