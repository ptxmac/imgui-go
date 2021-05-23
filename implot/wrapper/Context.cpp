
#include "implot.h"

#include "Context.h"
#include "../../wrapper/WrapperConverter.h"

ImpContext impCreateContext()
{
  ImPlotContext *context = ImPlot::CreateContext();
  return reinterpret_cast<ImpContext>(context);
}
