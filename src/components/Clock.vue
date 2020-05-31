<template>
  <ul class="clock-hands">
    <ClockHand
      v-for="(location, name) in locations"
      :key="name"
      :name="name"
      :location="location"
    />
  </ul>
</template>

<script>
import ClockHand from './ClockHand.vue'

export default {
  components: {
    ClockHand
  },
  props: {
    locations: Object
  }
}
</script>

<style scoped>
.clock-hands {
  padding: 0;
  list-style: none;
}
</style>
