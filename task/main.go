package task

import "github.com/axieinfinity/bridge-v2/models"

const (
	ACK_WITHDREW_TASK           = "acknowledgeWithdrew"
	DEPOSIT_TASK                = "deposit"
	WITHDRAWAL_TASK             = "withdrawal"
	VOTE_BRIDGE_OPERATORS_TASK  = "voteBridgeOperatorsBySignatures"
	RELAY_BRIDGE_OPERATORS_TASK = "relayBridgeOperators"
	BRIDGE_SYNC_REWARD_TASK     = "bridgeSyncReward"

	STATUS_PENDING    = "pending"
	STATUS_FAILED     = "failed"
	STATUS_PROCESSING = "processing"
	STATUS_DONE       = "done"

	GATEWAY_CONTRACT              = "gateway"
	GOVERNANCE_CONTRACT           = "governance"
	TRUSTED_ORGANIZATION_CONTRACT = "trustedorganization"
	ETH_GOVERNANCE_CONTRACT       = "ethgovernance"
	ETH_GATEWAY_CONTRACT          = "ethgateway"
	BRIDGEADMIN_CONTRACT          = "bridgeadmin"
	BRIDGE_REWARD_CONTRACT        = "bridgereward"
	RONIN_VALIDATOR_SET_CONTRACT  = "roninvalidatorset"
)

const (
	VoteStatusPending = iota
	VoteStatusApproved
	VoteStatusExecuted
)

type Tasker interface {
	collectTask(t *models.Task)
	send()
}
