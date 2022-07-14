package storage

// BlockLog ...
type BlockLog struct {
	Chain      string    `gorm:"type:TEXT"`
	BlockHash  string    `gorm:"type:TEXT"`
	ParentHash string    `gorm:"type:TEXT"`
	Height     int64     `gorm:"type:BIGINT"`
	BlockTime  int64     `gorm:"type:BIGINT"`
	Type       BlockType `gorm:"block_type"`
	CreateTime int64     `gorm:"type:BIGINT"`
}

// TxLog ...
type TxLog struct {
	Chain              string `gorm:"type:TEXT"`
	EventID            string
	TxType             TxType      `gorm:"type:tx_types"`
	TxHash             string      `gorm:"type:TEXT"`
	SenderAddr         string      `gorm:"type:TEXT"`
	Data               string      `gorm:"type:TEXT"`
	DestinationChainID string      `gorm:"type:TEXT"`
	ExpireHeight       int64       `gorm:"type:BIGINT"`
	Timestamp          int64       `gorm:"type:BIGINT"`
	BlockHash          string      `gorm:"type:TEXT"`
	Height             int64       `gorm:"type:BIGINT"`
	Status             TxLogStatus `gorm:"type:tx_log_statuses"`
	EventStatus        EventStatus
	ConfirmedNum       int64  `gorm:"type:BIGINT"`
	CreateTime         int64  `gorm:"type:BIGINT"`
	UpdateTime         int64  `gorm:"type:BIGINT"`
	SwapID             string `gorm:"primaryKey"`
	OriginChainID      string `gorm:"type:TEXT"`
	DepositNonce       uint64 `gorm:"type:BIGINT"`
	SwapStatus         uint8
	ResourceID         string `gorm:"type:TEXT"`
	ReceiverAddr       string `gorm:"type:TEXT"`
	WorkerChainAddr    string `gorm:"type:TEXT"`
	OutAmount          string `gorm:"type:TEXT"`
	InAmount           string `gorm:"type:TEXT"`
}

// Registration
type Registration struct {
	RelayerAddress string `gorm:"type:TEXT"`
	Status         EventStatus
	CreateTime     int64 `gorm:"type:BIGINT"`
	UpdateTime     int64 `gorm:"type:BIGINT"`
}

// Event ...
type Event struct {
	SwapID             string `gorm:"primaryKey"`
	ChainID            string
	DestinationChainID string
	OriginChainID      string
	SenderAddr         string
	ReceiverAddr       string
	InTokenAddr        string
	OutTokenAddr       string
	InAmount           string
	OutAmount          string
	Height             int64
	Status             EventStatus
	CreateTime         int64
	UpdateTime         int64
	DepositNonce       uint64
	ResourceID         string
	TxType             string
}

// TxSent ...
type TxSent struct {
	ID         int64    `json:"id"`
	Chain      string   `json:"chain" gorm:"type:TEXT"`
	SwapID     string   `json:"swap_id" gorm:"type:TEXT"`
	Type       TxType   `json:"type" gorm:"type:tx_types"`
	TxHash     string   `json:"tx_hash" gorm:"type:TEXT"`
	ErrMsg     string   `json:"err_msg" gorm:"type:TEXT"`
	Status     TxStatus `json:"status" gorm:"type:tx_statuses"`
	CreateTime int64    `json:"create_time" gorm:"type:BIGINT"`
	UpdateTime int64    `json:"update_time" gorm:"type:BIGINT"`
}

// GasPrice
type GasPrice struct {
	ChainName  string `gorm:"primaryKey"`
	Price      string `gorm:"type:TEXT"`
	UpdateTime int64  `json:"update_time" gorm:"type:BIGINT"`
}

type ResourceId struct {
	Name string `gorm:"primaryKey"`
	ID   string `gorm:"type:TEXT"`
}
