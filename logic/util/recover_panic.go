package util

import "log"

type panicHandler func(err interface{})

type optionPanicHandler func(*optionPanicHandlerParams)

type optionPanicHandlerParams struct {
	panicHandler panicHandler
}

func WithPanicHandler(panicHandler panicHandler) optionPanicHandler {
	return func(optParams *optionPanicHandlerParams) {
		optParams.panicHandler = panicHandler
	}
}

func WithRecover(f func(), opts ...optionPanicHandler) {
	defer func() {
		if err := recover(); err != nil {
			optParams := &optionPanicHandlerParams{}
			for _, opt := range opts {
				opt(optParams)
			}
			if optParams.panicHandler == nil {
				optParams.panicHandler = defaultPanicHandler
			}
			optParams.panicHandler(err)
		}
	}()

	f()
}

func defaultPanicHandler(err interface{}) {
	log.Printf("panic: %v", err)
}
