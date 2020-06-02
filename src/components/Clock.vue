<template>
  <svg :viewBox="viewBox">
    <circle
      cx=0
      cy=0
      :r="clockRadius"
      :style="clockStyle"
    />

    <Status
      v-for="(status, index) in statusAssignations"
      :key="status.status"
      :index="index"
      :elCount="statusAssignations.length"
      :status="status"
      :radius="clockRadius"
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
import Status from './Status.vue'

export default {
  components: {
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

    statusAssignations() {
      const statuses = Object.entries(this.people).reduce((statuses, [name, person]) => {
        return {
          ...statuses,
          [person.status]: {
            ...statuses[person.status],
            [name]: person
          }
        }
      }, {})

      return this.statuses.map(status => ({
        'status': status,
        'people': statuses[status] || {}
      }))
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
