package mycasbin

import (
	"github.com/jason-wj/logger/logbase"
	"sync/atomic"
)

// Logger is the implementation for a Logger using golang log.
type Logger struct {
	enable int32
}

// EnableLog controls whether print the message.
func (l *Logger) EnableLog(enable bool) {
	i := 0
	if enable {
		i = 1
	}
	atomic.StoreInt32(&(l.enable), int32(i))
}

// IsEnabled returns if logger is enabled.
func (l *Logger) IsEnabled() bool {
	return atomic.LoadInt32(&(l.enable)) != 0
}

// LogModel log info related to model.
func (l *Logger) LogModel(model [][]string) {
	var str string
	for i := range model {
		for j := range model[i] {
			str += " " + model[i][j]
		}
		str += "\n"
	}
	logbase.DefaultLogger.Log(logbase.InfoLevel, str)
}

// LogEnforce log info related to enforce.
func (l *Logger) LogEnforce(matcher string, request []interface{}, result bool, explains [][]string) {
	logbase.DefaultLogger.Fields(map[string]interface{}{
		"matcher":  matcher,
		"request":  request,
		"result":   result,
		"explains": explains,
	}).Log(logbase.InfoLevel, nil)
}

// LogRole log info related to role.
func (l *Logger) LogRole(roles []string) {
	logbase.DefaultLogger.Fields(map[string]interface{}{
		"roles": roles,
	})
}

// LogPolicy log info related to policy.
func (l *Logger) LogPolicy(policy map[string][][]string) {
	data := make(map[string]interface{}, len(policy))
	for k := range policy {
		data[k] = policy[k]
	}
	logbase.DefaultLogger.Fields(data).Log(logbase.InfoLevel, nil)
}

//func (l *Logger) Print(v ...interface{}) {
//	if l.IsEnabled() {
//		logger.DefaultLogger.Log(logger.InfoLevel, v...)
//	}
//}
//
//func (l *Logger) Printf(format string, v ...interface{}) {
//	if l.IsEnabled() {
//		logger.DefaultLogger.Logf(logger.InfoLevel, format, v...)
//	}
//}
