import { defineStore } from 'pinia'
import type { GetValidatorDashboardRewardsResponse } from '~/types/api/validator_dashboard'
import type { DashboardKey } from '~/types/dashboard'
import type { TableQueryParams } from '~/types/datatable'
import { DAHSHBOARDS_NEXT_EPOCH_ID } from '~/types/dashboard'

const validatorDashboardRewardsStore = defineStore(
  'validator_dashboard_rewards_store',
  () => {
    const data = ref<GetValidatorDashboardRewardsResponse>()
    const query = ref<TableQueryParams>()

    return {
      data,
      query,
    }
  },
)

export function useValidatorDashboardRewardsStore() {
  const { fetch } = useCustomFetch()
  const {
    data, query: storedQuery,
  } = storeToRefs(
    validatorDashboardRewardsStore(),
  )
  const isLoading = ref(false)

  const { slotViz } = useValidatorSlotVizStore()

  const rewards = computed(() => data.value)
  const query = computed(() => storedQuery.value)

  async function getRewards(
    dashboardKey: DashboardKey,
    query?: TableQueryParams,
  ) {
    if (!dashboardKey) {
      data.value = undefined
      isLoading.value = false
      storedQuery.value = undefined
      return undefined
    }
    isLoading.value = true
    storedQuery.value = query
    const res = await fetch<GetValidatorDashboardRewardsResponse>(
      'DASHBOARD_VALIDATOR_REWARDS',
      undefined,
      { dashboardKey },
      query,
    )

    isLoading.value = false
    if (JSON.stringify(storedQuery.value) !== JSON.stringify(query)) {
      return // in case some query params change while loading
    }

    // If we are on the first page we get the next Epoch slot viz data and create a future entry
    if (!query?.cursor && slotViz.value && res.data?.length) {
      const searchEpoch = res.data[0].epoch
      const nextEpoch = slotViz.value?.findLast(e => e.epoch > searchEpoch)

      if (nextEpoch) {
        res.data = [
          {
            duty: {
              attestation: 0,
              proposal: 0,
              slashing: 0,
              sync: 0,
            },
            epoch: nextEpoch.epoch,
            group_id: DAHSHBOARDS_NEXT_EPOCH_ID,
            reward: {
              cl: '0',
              el: '0',
            },
          },
          ...res.data,
        ]
      }
    }

    data.value = res
    return res
  }

  return {
    getRewards,
    isLoading,
    query,
    rewards,
  }
}
