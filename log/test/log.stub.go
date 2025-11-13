package log_test

import (
	"golang-api/log"

	"github.com/jaswdr/faker/v2"
)

var fake = faker.New()

func CreateLog() *log.Log {
	return &log.Log{
		ID:      fake.UUID().V4(),
		Type:    log.INFO,
		Tags:    []string{"tag1", "tag2"},
		Message: "This is a log message",
	}
}

func CreateManyLogs(n int) []*log.Log {
	logs := make([]*log.Log, n)
	for i := 0; i < n; i++ {
		logs[i] = CreateLog()
	}
	return logs
}
