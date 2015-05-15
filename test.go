package main
import (
	"fmt"
	"dtrace"
)

func main() {
	var error int
	var hdl dtrace.DTraceHandle
	
	hdl = dtrace.Open(dtrace.DTraceVersion, 0, &error)
	if (hdl == nil) {
		fmt.Printf("Unable to initialise DTrace: %s\n",
			dtrace.ErrMsg(nil, error))
		return
	}
}
