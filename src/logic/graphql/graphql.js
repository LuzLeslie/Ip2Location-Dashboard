import gql from 'graphql-tag'

export const IP_ANALYZE_QUERY = gql`
  query IpInfoQuery {
    ipData{
      ip
      country_name
      city
      proxy{
        is_proxy
      }
  }
  }
`

export const IP_ANALYZE_SUBSCRIPTION = gql`
  subscription IpInfoSubscription {
    ipData {
      ip
      country_name
      city
      proxy{
        is_proxy
      }
    }
  }
`
