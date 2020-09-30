<template>
  <div class="home">
    <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <div class="container-fluid">
      <div class="row">
        <!-- Total analyzed -->
        <div class="col-xl-3 col-md-6 mb-4">
          <b-card
          border-variant="primary"
          header="Total analyzed"
          header-bg-variant="primary"
          header-text-variant="white"
          align="center">
          <b-card-text>{{totalAnalyzed}}</b-card-text>
          </b-card>
        </div>

        <!-- Proxy percentage -->
        <div class="col-xl-3 col-md-6 mb-4">
          <b-card
          border-variant="warning"
          header="Proxy percentage"
          header-bg-variant="warning"
          header-text-variant="white"
          align="center">
          <b-card-text>{{percentageProxy}} %</b-card-text>
          </b-card>
        </div>

        <!-- Errors -->
        <div class="col-xl-3 col-md-6 mb-4">
          <b-card
          border-variant="danger"
          header="Errors"
          header-bg-variant="danger"
          header-text-variant="white"
          align="center">
          <b-card-text>0</b-card-text>
          </b-card>
        </div>

        <!-- Status backend -->
        <div class="col-xl-3 col-md-6 mb-4">
          <b-card
          border-variant="info"
          header="Status backend"
          header-bg-variant="info"
          header-text-variant="white"
          align="center">
          <b-card-text>{{statusBackend}}</b-card-text>
          </b-card>
        </div>

      </div>

      <div class="row">

        <!-- Area Chart -->
            <div class="col-xl-8 col-lg-7">
              <div class="card shadow mb-4">
                <!-- Card Header - Dropdown -->
                <div class="card-header py-3 d-flex flex-row align-items-center justify-content-between">
                  <h6 class="m-0 font-weight-bold text-primary">IP report analyzed</h6>
                </div>
                <!-- Card Body -->
                <div class="card-body">

                  <!-- <vue-datamaps></vue-datamaps> -->
                  <div>
                    <vue-datamaps
                      :geographyConfig="geographyConfig"
                      :bubblesConfig="bubblesConfig"
                      :fills="fills"
                      @custom:popup-bubble="popupTemplate"
                      bubbles
                    >
                      <div slot="hoverBubbleInfo" class="hoverinfo" style="text-align:center;">
                          <b>{{ popupData.country }}</b> : {{ popupData.quantity }}<br>
                      </div>
                    </vue-datamaps>
                  </div>
                </div>
              </div>
            </div>

        <!-- Pie Chart -->
            <div class="col-xl-4 col-lg-5">
              <div class="card shadow mb-4">
                <!-- Card Header - Dropdown -->
                <div class="card-header py-3 d-flex flex-row align-items-center justify-content-between">
                  <h6 class="m-0 font-weight-bold text-primary">Statistics proxy</h6>
                </div>
                <!-- Card Body -->
                <div class="card-body">
                  <!-- <b-table  small responsive borderless striped hover :fields="fields" :items="topCountrys" ></b-table> -->
                    <vc-donut
                      background="white" foreground="grey"
                      :size="250" unit="px" :thickness="50"
                      has-legend legend-placement="top"
                      :sections="sections" :total="100"
                      :start-angle="0" :auto-adjust-text-size="true">
                    </vc-donut>
                </div>
              </div>
            </div>

          </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
import { mapState } from 'vuex'

export default {
  name: 'Home',
  data () {
    return {
      fields: [
        { key: 'name', sortable: false },
        { key: 'quantity', sortable: true }
      ],

      geographyConfig: {
        popupOnHover: true,
        highlightOnHover: true
      },
      fills: {
        defaultFill: '#ABDDA4',
        BLUE: 'blue',
        GREEN: 'green'
      },
      popupData: {
        quantity: '',
        date: '',
        country: ''
      }
    }
  },
  components: {
    // HelloWorld
  },
  methods: {
    popupTemplate ({ datum }) {
      this.popupData = {
        quantity: datum.quantity,
        date: datum.date,
        country: datum.country
      }
    }
  },
  computed: {
    ...mapState([
      'totalAnalyzed', 'percentageProxy', 'statusBackend', 'topCountrys', 'bubblesConfig', 'sections'
    ])
  },
  created () {
    this.$store.commit('GetDataDashboard')
  }
}
</script>
