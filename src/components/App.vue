<template>
  <div id="app">
    <Clock
      v-bind:people="people"
      v-bind:statuses="statuses"

      :radius="radius"
      :darkColor="darkColor"
      :lightColor="lightColor"
    />
  </div>
</template>

<script>
import * as fakeAPI from '@/services/fake-statuses'
import * as statusesAPI from '@/services/statuses'
import Clock from './Clock.vue'

export default {
  components: {
    Clock
  },
  data() {
    const people = {}
    const statuses = []

    const radius = 250
    const darkColor = "#4c4c4c"
    const lightColor = "#945353"

    const useFakeData = false
    const refreshIntervalMs = 20000

    return { people, statuses, radius, darkColor, lightColor, useFakeData, refreshIntervalMs }
  },
  mounted() {
    const fetch = () => {
      const api = this.useFakeData ? fakeAPI : statusesAPI

      api.fetchEverything()
        .then(json => {
          this.people = json.people
          this.statuses = json.statuses
        })
      }

    fetch()

    let looper = setInterval(fetch, this.refreshIntervalMs)

    window.addEventListener('blur', () => {
      clearInterval(looper)
    })

    window.addEventListener('focus', () => {
      fetch()
      looper = setInterval(fetch, this.refreshIntervalMs)
    })
  }
}
</script>

<style>
body {
  background: #b89f9f;
  font-family: Avenir, Helvetica, Arial, sans-serif;
}
</style>
