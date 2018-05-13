package unmarshal

import (
	"github.com/nawazish-github/airfare-poller/models"
)

//Unmarshal ...
type Unmarshal interface {
	Unmarshal(data []byte) (*models.AirfareResponse, error)
}
