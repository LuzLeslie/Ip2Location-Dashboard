import axios from 'axios'

const ENDPOINT_PATH = 'http://localhost:4000/graphql'

export default {
  GetDataDashboard () {
    const newQuery = `{
        totalAnalyzed
        proxyStatistic{
            proxy
            clean
        }
        topCountrys{
            name
            quantity
            code
        }
    }
    `
    const query = { query: newQuery }
    return axios.post(ENDPOINT_PATH, query)
  },
  RunQuery (queryConfig) {
    const query = { query: queryConfig }
    return axios.post(ENDPOINT_PATH, query)
  }
}
