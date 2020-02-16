package xresilient

import (
	"context"
	"errors"
	"fmt"
	"gitlab.myteksi.net/gophers/go/commons/util/resilience/retry"

	"gitlab.myteksi.net/gophers/go/commons/util/resilience/circuitbreaker"
)

const circuit = "errorCircuit"

func errorFunc() error {
	return errors.New("err")
}

func TestCircuitBreaker() {
	cbSetting := &circuitbreaker.Setting{
		TagStr:              circuit,
		Timeout:             100,
		ErrPercentThreshold: 50,
	}

	circuitbreaker.Initialize(cbSetting)

	var i int
	for !circuitbreaker.IsCircuitOpen(circuit) {
		i++
		_ = circuitbreaker.Do(context.Background(), circuit, errorFunc)
	}

	fmt.Printf("Circuit opened after %v iterations\n", i)

	fmt.Println(circuitbreaker.Do(context.Background(), circuit, errorFunc))
}

func TestRetry() {
	fmt.Println(retry.Do(context.Background(), doSomething))
}

var iterations = 0

func doSomething() error {
	iterations++
	fmt.Println("Iteration", iterations)
	if iterations < 3 {
		return errors.New("error")
	}
	return nil
}
