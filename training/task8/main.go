package main

import (
	"context"
	"errors"
	"fmt"
)

type ErrorHandlerFunc func(ctx context.Context, err error) error

type Doer struct {
	errHandler ErrorHandlerFunc
}

func (doer Doer) DoSth(ctx context.Context) error {
	err := errors.New("something wrong here")

	if doer.errHandler == nil {
		return defautErrorHandler(ctx, err)
	} else {
		return doer.errHandler(ctx, err)
	}
}

func defautErrorHandler(ctx context.Context, e error) error {
	fmt.Println("Error is handling by default handler: ", e)
	return e
}

func main() {
	ctx := context.Background();
	d1 := Doer{}

	d1.DoSth(ctx)

	d2 := Doer{
		errHandler: func(ctx context.Context, err error) error{
			fmt.Println("This is specific handler of object:", err)
			return err
		},
	}
	d2.DoSth(ctx)
}