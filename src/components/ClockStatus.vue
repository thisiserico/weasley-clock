<template>
  <li class="clock-status">
    {{ assignation.status }}:

    <ClockHand
      v-for="(person, name) in assignation.people"
      :key="name"
      :name="person.name"
    />
  </li>
</template>

<script>
import ClockHand from './ClockHand.vue'

export default {
  components: {
    ClockHand
  },
  props: {
    assignation: Object
  }
}
</script>

<style scoped>
.clock-status {
  display: inline-block;
  margin: 1rem 0;
  width: 100%;
}
</style>
