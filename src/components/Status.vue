<template>
  <g>
    <g v-if="Object.keys(assignation.people).length > 0">
      <line
        :style="line"
        :x1="radius"
        :y1="radius"
        :x2="radius"
        :y2="p30radius"
        :transform="rotation"
      />
      <Person
        v-for="(person, name) in assignation.people"
        :key="name"
        :degrees="degrees"
        :radius="radius"
        :name="person.name"
        :dark="dark"
        :light="light"
      />
    </g>

    <text
      text-anchor="middle"
      :x="radius"
      :y="p20radius"
      :transform="rotation">
      {{ assignation.status }}
    </text>
  </g>
</template>

<script>
import Person from './Person.vue'

export default {
  components: {
    Person
  },
  props: {
    assignation: Object,
    index: Number,
    elements: Number,
    radius: Number,
    dark: String,
    light: String
  },
  computed: {
    line() {
      return "stroke: #" + this.dark + "; stroke-width: " + this.radius * .02 + "px"
    },
    p20radius() {
      return this.radius * .2
    },
    p30radius() {
      return this.radius * .3
    },
    degrees() {
      return this.index * 360 / this.elements
    },
    rotation() {
      return "rotate(" + this.degrees + " " + this.radius + " " + this.radius + ")"
    }
  }
}
</script>
