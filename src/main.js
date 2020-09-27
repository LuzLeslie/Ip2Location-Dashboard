import Vue from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import store from './store'

import { BootstrapVue } from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import VueApollo from 'vue-apollo'
import { InMemoryCache } from 'apollo-cache-inmemory'
import { ApolloClient } from 'apollo-client'
import { split } from 'apollo-link'
import { HttpLink } from 'apollo-link-http'
import { WebSocketLink } from 'apollo-link-ws'
import { getMainDefinition } from 'apollo-utilities'

import { library } from '@fortawesome/fontawesome-svg-core'
import { faTachometerAlt, faTable, faSearchLocation, faInfoCircle } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import VueDatamaps from 'vue-datamaps'

import Donut from 'vue-css-donut-chart'
import 'vue-css-donut-chart/dist/vcdonut.css'

library.add(faTachometerAlt, faTable, faSearchLocation, faInfoCircle)
Vue.component('font-awesome-icon', FontAwesomeIcon)

Vue.use(VueDatamaps)

Vue.config.productionTip = false

Vue.use(BootstrapVue)

Vue.use(Donut)

// Graphql Config
const httpLink = new HttpLink({
  uri: 'http://localhost:4000/graphql'
})
const wsLink = new WebSocketLink({
  uri: 'ws://localhost:4000/subscriptions',
  options: {
    reconnect: true
  }
})

const link = split(
  ({ query }) => {
    const { kind, operation } = getMainDefinition(query)
    return kind === 'OperationDefinition' && operation === 'subscription'
  },
  wsLink,
  httpLink
)

const apolloClient = new ApolloClient({
  link,
  cache: new InMemoryCache(),
  connectToDevTools: true
})

const apolloProvider = new VueApollo({
  defaultClient: apolloClient
})

Vue.use(VueApollo)
// Graphql Config

new Vue({
  router,
  store,
  apolloProvider,
  render: h => h(App)
}).$mount('#app')
