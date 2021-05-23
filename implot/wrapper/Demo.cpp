
#include "implot.h"

#include "Demo.h"
#include "../../wrapper/WrapperConverter.h"


void impShowDemoWindow(IggBool *open) {
  BoolWrapper openArg(open);

  ImPlot::ShowDemoWindow(openArg);
}
