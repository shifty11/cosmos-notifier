package common

import (
	"github.com/shifty11/cosmos-notifier/log"
)

const DefaultMaxErrorCntUntilReport = 96

type ErrorReporter interface {
	ReportErrorIfNecessary(id int, errorText string)
	ResetErrorCount(id int)
}

type errorReporter struct {
	errorCnt               map[int]int
	maxErrorCntUntilReport int
}

func NewErrorReporter(maxErrorCntUntilReport int) ErrorReporter {
	return &errorReporter{
		errorCnt:               make(map[int]int),
		maxErrorCntUntilReport: maxErrorCntUntilReport,
	}
}

func (er *errorReporter) ReportErrorIfNecessary(id int, errorText string) {
	er.errorCnt[id]++
	if er.errorCnt[id]%er.maxErrorCntUntilReport == 0 { // report every `maxErrorCntUntilReport` times
		log.Sugar.Error(errorText)
	}
}

func (er *errorReporter) ResetErrorCount(id int) {
	er.errorCnt[id] = 0
}
