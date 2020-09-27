<template>
  <div class="analyze">
    <b-col lg="4" class="pb-2">
        <b-button @click="runRequest" squared size="sm">Run query</b-button>
      </b-col>
    <b-row>
    <b-col class="col-xxl-4 col-xl-6 mb-4">
      <b-form-textarea
      id="textarea"
      v-model="queryRequest"
      placeholder="Enter query..."
      rows="3"
      max-rows="12"
    ></b-form-textarea>
    </b-col>

    <b-col xl="6" mb="4">
      <b-form-textarea style = "border-style: solid; border-color: green;"
      v-model="responseRequest"
      plaintext
      placeholder="Response appears here..."
      rows="3"
      max-rows="12"
    ></b-form-textarea>
    </b-col>

    </b-row>

  </div>
</template>
<script>

import dashboard from '@/logic/dashboard.js'

export default {
  name: 'Analyze',
  data () {
    return {
      queryRequest: `
  mutation {
    analyze(ip:"1.1.1.1"){
      ip
      city
      country_code
      country_name
      proxy{
        is_proxy
      }
    }
  }
      `,
      responseRequest: null
    }
  },
  methods: {
    async runRequest () {
      const resp = await dashboard.RunQuery(this.queryRequest)
      this.responseRequest = resp.data.data
    }
  },
  created () {
    this.$store.commit('newTittleBar', 'Analyze ip')
  }
}
</script>
