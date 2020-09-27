import Vue from 'vue'
import Vuex from 'vuex'
import dashboard from '@/logic/dashboard.js'

Vue.use(Vuex)

function convertCountryCode (codeIp2location) {
  switch (codeIp2location) {
    case 'US':
      return 'USA'
    case 'JP':
      return 'JPN'
    case 'MX':
      return 'MEX'
    case 'BD':
      return 'BGD'
  }
}

export default new Vuex.Store({
  state: {
    totalAnalyzed: 0,
    percentageProxy: 0,
    statusBackend: 'down',
    topCountrys: null,
    tittleBar: '',
    bubblesConfig: {
      popupTemplate: true,
      data: []
    },
    sections: []
  },
  mutations: {
    async GetDataDashboard (state) {
      state.tittleBar = 'Dashboard'
      try {
        const result = await dashboard.GetDataDashboard()
        state.totalAnalyzed = result.data.data.totalAnalyzed
        state.percentageProxy = result.data.data.proxyStatistic.proxy
        state.statusBackend = 'Live'
        state.topCountrys = result.data.data.topCountrys

        state.sections = [
          { label: 'Ip clean', value: result.data.data.proxyStatistic.clean, color: '#36A2EB' },
          { label: 'Proxy', value: result.data.data.proxyStatistic.proxy, color: '#ffc107' }
        ]

        for (let i = 0; i < state.topCountrys.length; i++) {
          const dataTemp = {
            radius: state.topCountrys[i].quantity,
            quantity: state.topCountrys[i].quantity,
            country: state.topCountrys[i].name,
            centered: convertCountryCode(state.topCountrys[i].code),
            fillKey: 'BLUE'
          }
          state.bubblesConfig.data.push(dataTemp)
        }
      } catch (error) {
        state.statusBackend = 'Down'
      }
    },
    newTittleBar (state, newTittle) {
      state.tittleBar = newTittle
    }

  },
  actions: {
  },
  modules: {
  }
})
