<template>
  <div id="app">
    <Clock v-bind:locations="locations"/>
  </div>
</template>

<script>
import * as locationsAPI from '@/services/locations'
import Clock from './Clock.vue'

export default {
  components: {
    Clock
  },
  data() {
    const locations = {}
    return { locations }
  },
  created() {
    locationsAPI.fetchAll()
      .then(locations => {
        this.locations = locations
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
