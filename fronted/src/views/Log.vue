<template>
  <div class="about">
    <div class="loganalyze">
      <div class="card shadow mb-4">
        <div class="card-body">
          <div class="table-responsive">
          <table class="table table-hover table-sm table-bordered">
            <thead class="thead-light">
              <tr>
                <th>Ip</th>
                <th>Country</th>
                <th>City</th>
                <th>Anonymous</th>
              </tr>
            </thead>
            <tbody>
              <tr :class="{ 'table-warning' : dataInfo.proxy.is_proxy >= 1 }" v-for="(dataInfo, id) in ipData" :key="id">
                <td> {{ dataInfo.ip }} </td>
                <td>{{ dataInfo.country_name }}</td>
                <td>{{ dataInfo.city }}</td>
                <td v-if="dataInfo.proxy.is_proxy >= 1" >True</td>
                <td v-else>-</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
  </div>
</template>
<script>

import {
  IP_ANALYZE_QUERY,
  IP_ANALYZE_SUBSCRIPTION
} from '@/logic/graphql/graphql'

export default {
  name: 'LogAnalyze',
  apollo: {
    ipData: {
      query: IP_ANALYZE_QUERY,
      subscribeToMore: {
        document: IP_ANALYZE_SUBSCRIPTION,
        updateQuery: (previousData, { subscriptionData }) => {
          return {
            ipData: [...previousData.ipData, subscriptionData.data.ipData]
          }
        }
      }
    }
  },
  created () {
    this.$store.commit('newTittleBar', 'Real time activity log')
  }
}
</script>
