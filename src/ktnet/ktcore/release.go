package ktcore

import (
	"ktnet/ktcore/ktlog"
)

//Release ...
func Release() {
	ktlog.CloseLog()
}
