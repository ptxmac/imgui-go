package imgui

// #include "implot/wrapper/Context.h"
import "C"

type ImPlotContext struct {
	handle C.ImpContext
}

func ImPlotCreateContext() *ImPlotContext {
	return &ImPlotContext{
		handle: C.impCreateContext(),
	}
}
