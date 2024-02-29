package handlers

import "net/http"

// All handler function names must include the HTTP method and the path they handle
// Public handlers may only be authenticated by an API key
// Public handlers must never call internal handlers

func (h HandlerService) PublicGetHealthz(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetHealthzLoadbalancer(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicPostOauthToken(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetUserDashboards(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicPostAccountDashboards(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h HandlerService) PublicGetAccountDashboard(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicDeleteAccountDashboard(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

func (h HandlerService) PublicPostAccountDashboardGroups(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h HandlerService) PublicDeleteAccountDashboardGroups(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

func (h HandlerService) PublicPostAccountDashboardAccounts(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h HandlerService) PublicGetAccountDashboardAccounts(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicDeleteAccountDashboardAccounts(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

func (h HandlerService) PublicPutAccountDashboardAccount(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicPostAccountDashboardPublicIds(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h HandlerService) PublicPutAccountDashboardPublicId(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicDeleteAccountDashboardPublicId(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

func (h HandlerService) PublicGetAccountDashboardTransactions(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicPutAccountDashboardTransactionsSettings(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicPostValidatorDashboards(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h HandlerService) PublicGetValidatorDashboard(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicDeleteValidatorDashboard(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

func (h HandlerService) PublicPostValidatorDashboardGroups(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h HandlerService) PublicDeleteValidatorDashboardGroups(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

func (h HandlerService) PublicPostValidatorDashboardValidators(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h HandlerService) PublicGetValidatorDashboardValidators(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicDeleteValidatorDashboardValidators(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicPostValidatorDashboardPublicIds(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h HandlerService) PublicPutValidatorDashboardPublicId(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicDeleteValidatorDashboardPublicId(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

func (h HandlerService) PublicGetValidatorDashboardSlotViz(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetValidatorDashboardSummary(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetValidatorDashboardGroupSummary(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetValidatorDashboardSummaryChart(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetValidatorDashboardRewards(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetValidatorDashboardGroupRewards(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetValidatorDashboardRewardsChart(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetValidatorDashboardDuties(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetValidatorDashboardBlocks(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetValidatorDashboardHeatmap(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetValidatorDashboardGroupHeatmap(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetValidatorDashboardExecutionLayerDeposits(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetValidatorDashboardConsensusLayerDeposits(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetValidatorDashboardWithdrawals(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkValidators(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkValidator(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkValidatorDuties(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkAddressValidators(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkWithdrawalCredentialValidators(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkValidatorStatuses(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkValidatorLeaderboard(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkValidatorQueue(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkEpochs(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkEpoch(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkBlocks(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkBlock(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkSlots(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkSlot(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkValidatorBlocks(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkAddressPriorityFeeBlocks(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkAddressProposerRewardBlocks(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkForkedBlocks(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkForkedBlock(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkForkedSlot(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkBlockSizes(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkValidatorAttestations(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkEpochAttestations(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkSlotAttestations(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkBlockAttestations(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkAggregatedAttestations(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkEthStore(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkValidatorRewardHistory(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkValidatorBalanceHistory(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkValidatorPerformanceHistory(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkSlashings(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkValidatorSlashings(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkDeposits(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkValidatorDeposits(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkTransactionDeposits(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkWithdrawals(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkSlotWithdrawals(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}
func (h HandlerService) PublicGetNetworkBlockWithdrawals(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkValidatorWithdrawals(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkWithdrawalCredentialWithdrawals(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkVoluntaryExits(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkEpochVoluntaryExits(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkSlotVoluntaryExits(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkBlockVoluntaryExits(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkAddressBalanceHistory(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkAddressTokenSupplyHistory(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkAddressEventLogs(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkTransactions(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkTransaction(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkAddressTransactions(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkSlotTransactions(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkBlockTransactions(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkBlockBlobs(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkBlsChanges(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkEpochBlsChanges(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkSlotBlsChanges(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkBlockBlsChanges(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkValidatorBlsChanges(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkAddressEns(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkEns(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkBatches(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkLayer2ToLayer1Transactions(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkLayer1ToLayer2Transactions(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicPostNetworkBroadcasts(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h HandlerService) PublicGetEthPriceHistory(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkGasNow(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkAverageGasLimitHistory(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkGasUsedHistory(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetRocketPoolNodes(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetRocketPoolMinipools(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetNetworkSyncCommittee(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetMultisigSafe(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetMultisigSafeTransactions(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h HandlerService) PublicGetMultisigTransactionConfirmations(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}