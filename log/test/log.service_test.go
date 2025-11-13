package log_test

import (
	"golang-api/core"
	"golang-api/log"
	"golang-api/query"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testLog(t *testing.T, expected *log.Log, actual *log.Log) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Type, actual.Type)
	for i := range expected.Tags {
		assert.Equal(t, expected.Tags[i], actual.Tags[i])
	}
	assert.Equal(t, expected.Message, actual.Message)
}

func NewLogModuleTest() *log.LogModule {
	module := &log.LogModule{
		Module: core.NewModule("LogTestModule"),
	}

	module.AddProvider(NewLogModelMock(module))
	module.AddProvider(log.NewLogService(module))

	return module
}

func TestSetupLogModule(t *testing.T) {
	module := NewLogModuleTest()
	logService := module.Get("LogService").(*log.LogService)
	logModel := module.Get("LogModel").(*LogModelMock)

	assert.NotNil(t, module, "Log module should be created")
	assert.NotNil(t, logService, "LogService should be created")
	assert.NotNil(t, logModel, "LogModel should be created")
}

func TestLogService_FindAll(t *testing.T) {
	module := NewLogModuleTest()
	logService := module.Get("LogService").(*log.LogService)
	logModel := module.Get("LogModel").(*LogModelMock)

	q := query.QueryFilter{}
	nbLogs := 3
	expectedLogs := CreateManyLogs(nbLogs)

	logModel.MockMethod("QueryFindAll", func(params ...any) any { return expectedLogs })

	newLogs, _ := logService.FindAll(q)

	called := logModel.IsMethodCalled("QueryFindAll")
	if !assert.Equal(t, true, called, "QueryFindAll method should be called") {
		return
	}

	if !assert.Len(t, newLogs, nbLogs, "Number of logs should be equal to expected") {
		return
	}
	for i := 0; i < nbLogs; i++ {
		testLog(t, expectedLogs[i], newLogs[i])
	}
}

func TestLogService_FindByID(t *testing.T) {
	module := NewLogModuleTest()
	logService := module.Get("LogService").(*log.LogService)
	logModel := module.Get("LogModel").(*LogModelMock)

	expectedLog := CreateLog()

	logModel.MockMethod("FindByID", func(params ...any) any {
		return expectedLog
	})

	newLog, _ := logService.FindByID(expectedLog.ID)

	called := logModel.IsMethodCalled("FindByID")
	if !assert.Equal(t, true, called, "FindByID method should be called") {
		return
	}
	params := logModel.IsParamsEqual("FindByID", expectedLog.ID)
	if !assert.Equal(t, true, params, "FindByID parameter should be the log ID") {
		return
	}

	testLog(t, expectedLog, newLog)
}

func TestLogService_FindOneBy(t *testing.T) {
	module := NewLogModuleTest()
	logService := module.Get("LogService").(*log.LogService)
	logModel := module.Get("LogModel").(*LogModelMock)

	expectedLog := CreateLog()

	logModel.MockMethod("FindOneBy", func(params ...any) any {
		return expectedLog
	})

	newLog, _ := logService.FindOneBy("type", expectedLog.Type)

	called := logModel.IsMethodCalled("FindOneBy")
	if !assert.Equal(t, true, called, "FindOneBy method should be called") {
		return
	}
	params := logModel.IsParamsEqual("FindOneBy", "type", expectedLog.Type)
	if !assert.Equal(t, true, params, "FindOneBy parameters should be the field and value") {
		return
	}

	testLog(t, expectedLog, newLog)
}

func TestLogService_Create(t *testing.T) {
	module := NewLogModuleTest()
	logService := module.Get("LogService").(*log.LogService)
	logModel := module.Get("LogModel").(*LogModelMock)

	expectedLog := CreateLog()
	var createdLog *log.Log

	logModel.MockMethod("Create", func(params ...any) any {
		createdLog = params[0].(*log.Log)
		return nil
	})

	err := logService.Create(expectedLog)
	if !assert.Nil(t, err, "Create should not return an error") {
		return
	}

	called := logModel.IsMethodCalled("Create")
	if !assert.Equal(t, true, called, "Create method should be called") {
		return
	}
	params := logModel.GetMethodParams("Create")
	if !assert.Len(t, params, 1, "Create should have one parameter") {
		return
	}
	paramLog, ok := params[0].(*log.Log)
	if !assert.Equal(t, true, ok, "Create parameter should be a log") {
		return
	}

	testLog(t, expectedLog, paramLog)
	testLog(t, expectedLog, createdLog)
}

func TestLogService_Printf(t *testing.T) {
	module := NewLogModuleTest()
	logService := module.Get("LogService").(*log.LogService)
	logModel := module.Get("LogModel").(*LogModelMock)

	var createdLog *log.Log

	logModel.MockMethod("Create", func(params ...any) any {
		createdLog = params[0].(*log.Log)
		return nil
	})

	tags := []string{"service", "info"}
	format := "This is a %s message with number %d"
	arg1 := "log"
	arg2 := 42

	logService.Printf(tags, format, arg1, arg2)

	called := logModel.IsMethodCalled("Create")
	if !assert.Equal(t, true, called, "Create method should be called") {
		return
	}

	expectedMessage := "This is a log message with number 42"
	if !assert.Equal(t, log.INFO, createdLog.Type, "Log type should be INFO") {
		return
	}
	if !assert.Equal(t, expectedMessage, createdLog.Message, "Log message should be formatted") {
		return
	}
	for i := range tags {
		if !assert.Equal(t, tags[i], createdLog.Tags[i], "Log tags should match") {
			return
		}
	}
}

func TestLogService_Errorf(t *testing.T) {
	module := NewLogModuleTest()
	logService := module.Get("LogService").(*log.LogService)
	logModel := module.Get("LogModel").(*LogModelMock)

	var createdLog *log.Log

	logModel.MockMethod("Create", func(params ...any) any {
		createdLog = params[0].(*log.Log)
		return nil
	})

	tags := []string{"service", "error"}
	format := "This is an %s message with code %d"
	arg1 := "error"
	arg2 := 500

	logService.Errorf(tags, format, arg1, arg2)

	called := logModel.IsMethodCalled("Create")
	if !assert.Equal(t, true, called, "Create method should be called") {
		return
	}

	expectedMessage := "This is an error message with code 500"
	if !assert.Equal(t, log.ERROR, createdLog.Type, "Log type should be ERROR") {
		return
	}
	if !assert.Equal(t, expectedMessage, createdLog.Message, "Log message should be formatted") {
		return
	}
	for i := range tags {
		if !assert.Equal(t, tags[i], createdLog.Tags[i], "Log tags should match") {
			return
		}
	}
}
