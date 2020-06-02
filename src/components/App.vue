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

    return { people, statuses, radius, darkColor, lightColor }
  },
  created() {
    const fetch = () => statusesAPI.fetchEverything()
      .then(json => {
        this.people = json.people
        this.statuses = json.statuses
      })

    fetch()
    setInterval(fetch, 60000)
  }
}
</script>

<style>
body {
  background: #b89f9f;
  font-family: Avenir, Helvetica, Arial, sans-serif;
}
</style>
