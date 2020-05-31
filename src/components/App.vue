<template>
  <div id="app">
    <Clock
      v-bind:people="people"
      v-bind:statuses="statuses"
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
    return { people, statuses }
  },
  created() {
    statusesAPI.fetchEverything()
      .then(json => {
        this.people = json.people
        this.statuses = json.statuses
      })
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
