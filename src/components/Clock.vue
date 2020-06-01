<template>
  <svg :viewBox="viewBox">
    <circle :style="circle" :cx="radius" :cy="radius" :r="circleRadius"></circle>

    <Status
      v-for="(assignation, index) in assignations"
      :key="assignation.status"
      :assignation="assignation"
      :index="index"
      :elements="assignations.length"
      :radius="radius"
      :dark="dark"
      :light="light"
    />

    <circle :style="spinner" :cx="radius" :cy="radius" :r="spinnerRadius"></circle>
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
    dark: String,
    light: String
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
    },
    diameter() {
      return this.radius * 2
    },

    viewBox() {
      return "0 0 " + this.diameter + " " + this.diameter
    },

    circle() {
      return "stroke: #" + this.dark + "; stroke-width: " + this.radius * .05 + "px; fill:#" + this.light
    },
    circleRadius() {
      return this.radius * .95
    },

    spinner() {
      return "stroke: #" + this.dark + "; stroke-width: " + this.radius * .02 + "px; fill:#" + this.light
    },
    spinnerRadius() {
      return this.radius * .05
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
