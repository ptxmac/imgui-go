package imgui

// #include "implot/wrapper/Demo.h"
import "C"

func ImPlotShowDemoWindow(open *bool) {
	openArg, openFin := wrapBool(open)
	defer openFin()

	C.impShowDemoWindow(openArg)
}
