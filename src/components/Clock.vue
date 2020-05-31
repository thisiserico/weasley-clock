<template>
  <div class="clock">
    <ul class="clock-statuses">
      <ClockStatus
        v-for="status in statuses"
        :key="status"
        :status="status"
      />
    </ul>

    <ul class="clock-hands">
      <ClockHand
        v-for="(person, name) in people"
        :key="name"
        :name="person.name"
        :status="person.status"
      />
    </ul>
  </div>
</template>

<script>
import ClockHand from './ClockHand.vue'
import ClockStatus from './ClockStatus.vue'

export default {
  components: {
    ClockHand,
    ClockStatus
  },
  props: {
    people: Object,
    statuses: Array
  }
}
</script>

<style scoped>
.clock-statuses {
  padding: 0;
  list-style: none;
}

.clock-hands {
  padding: 0;
  list-style: none;
}
</style>
