<template>
  <ul class="clock-statuses">
    <ClockStatus
      v-for="assignation in assignations"
      :key="assignation.status"
      :assignation="assignation"
    />
  </ul>
</template>

<script>
import ClockStatus from './ClockStatus.vue'

export default {
  components: {
    ClockStatus
  },
  props: {
    people: Object,
    statuses: Array
  },
  computed: {
    assignations() {
      const usedStatuses = {}
      for (const name in this.people) {
        const person = this.people[name]
        if (usedStatuses[person.status] === undefined) {
          usedStatuses[person.status] = {}
        }

        usedStatuses[person.status][name] = person
      }

      return this.statuses.map(status => ({
        'status': status,
        'people': usedStatuses[status] || {}
      }))
    }
  }
}
</script>

<style scoped>
.clock-statuses {
  padding: 0;
  list-style: none;
}
</style>
