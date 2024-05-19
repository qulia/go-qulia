package mock

import (
	"time"

	"github.com/qulia/go-qulia/v2/lib/common"
	"github.com/qulia/go-qulia/v2/mock/mock_time"
)

func GetMockTimeProviderDefault() common.TimeProvider {
	return mock_time.NewMockTimeProvider(time.Now(), time.Millisecond*10)
}

func GetMockTimeProviderLateScheduling() common.TimeProvider {
	return mock_time.NewMockTimeProvider(time.Now(), time.Second)
}
