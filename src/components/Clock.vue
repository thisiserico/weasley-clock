<template>
  <svg :viewBox="viewBox">
    <circle
      cx=0
      cy=0
      :r="clockRadius"
      :style="clockStyle"
    />

    <Status
      v-for="(rotation, status) in rotatedStatuses"
      :key="status"
      :status="status"
      :rotation="rotation"
      :radius="clockRadius"
      :darkColor="darkColor"
      :lightColor="lightColor"
    />

    <Person
      v-for="(person, name) in people"
      :key="name"
      :name="person.name"
      :rotation="rotatedStatuses[person.status]"
      :radius="radius"
      :darkColor="darkColor"
      :lightColor="lightColor"
    />

    <circle
      cx=0
      cy=0
      :r="spinnerRadius"
      :style="spinnerStyle"
    />
  </svg>
</template>

<script>
import Person from './Person.vue'
import Status from './Status.vue'

export default {
  components: {
    Person,
    Status
  },
  props: {
    people: Object,
    statuses: Array,

    radius: Number,
    darkColor: String,
    lightColor: String
  },
  computed: {
    viewBox() {
      const viewboxRadius = this.clockRadius * 1.1
      const viewboxLength = viewboxRadius * 2

      return `-${viewboxRadius} -${viewboxRadius} ${viewboxLength} ${viewboxLength}`
    },

    clockRadius() {
      return this.radius
    },
    clockStyle() {
      const strokeWidth = this.clockRadius * .05
      return `stroke: ${this.darkColor}; stroke-width: ${strokeWidth}px; fill: ${this.lightColor}`
    },

    rotatedStatuses() {
      const elCount = this.statuses.length

      return this.statuses.reduce((statuses, status, index) => ({
        ...statuses,
        [status]: index * 360 / elCount,
      }), {})
    },

    spinnerRadius() {
      return this.clockRadius * .05
    },
    spinnerStyle() {
      const strokeWidth = this.clockRadius * .02
      return `stroke: ${this.darkColor}; stroke-width: ${strokeWidth}px; fill:${this.lightColor}`
    }
  },
}
</script>

<style scoped>
svg {
  width: 90%;
  max-width: 500px;
  display: block;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
</style>
