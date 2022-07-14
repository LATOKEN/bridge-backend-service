package storage

// BlockType ...
type BlockType string

const (
	BlockTypeCurrent BlockType = "CURRENT"
	BlockTypeParent  BlockType = "PARENT"
)

// TxStatus ...
type TxStatus string

const (
	TxSentStatusInit     TxStatus = "INIT"
	TxSentStatusNotFound TxStatus = "NOT_FOUND"
	TxSentStatusPending  TxStatus = "PENDING"
	TxSentStatusFailed   TxStatus = "FAILED"
	TxSentStatusSuccess  TxStatus = "SUCCESS"
	TxSentStatusLost     TxStatus = "LOST"
)

// !!! TODO !!!

//
type TxType string

const (
	TxTypeDeposit TxType = "DEPOSIT"
	TxTypeClaim   TxType = "VOTE"
	TxTypePassed  TxType = "PASSED"
	TxTypeSpend   TxType = "SPEND"
	TxTypeExpired TxType = "EXPIRED"
	TxTypeUpdate  TxType = "UPDATE"
)

type EventStatus string

const (
	// PASSSED
	EventStatusPassedInit          EventStatus = "PASSSED_INIT"
	EventStatusPassedInitConfrimed EventStatus = "PASSSED_INIT_CONFIRMED"
	EventStatusPassedSent          EventStatus = "PASSSED_SENT"
	EventStatusPassedConfirmed     EventStatus = "PASSSED_CONFIRMED"
	EventStatusPassedFailed        EventStatus = "PASSSED_FAILED"
	EventStatusPassedSentFailed    EventStatus = "PASSSED_SENT_FAILED"

	EventStatusSpendConfirmed EventStatus = "SPEND_CONFIRMED"

	EventStatusDepositConfirmed EventStatus = "DEPOSIT_CONFIRMED"

	EventStatusClaimConfirmed EventStatus = "CLAIM_CONFIRM"

	EventStatusExpiredConfirmed EventStatus = "EXPIRED_CONFIRMED"

	EventStatusUpdateConfirmed EventStatus = "UPDATE_CONFIRMED"
	EventStatusUpdateFailed    EventStatus = "UPDATE_FAILED"
)

// TxLogStatus ...
type TxLogStatus string

const (
	TxStatusInit      TxLogStatus = "INIT"
	TxStatusConfirmed TxLogStatus = "CONFIRMED"
)
