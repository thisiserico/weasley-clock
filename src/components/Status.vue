<template>
  <g>
    <g v-if="Object.keys(assignation.people).length > 0">
      <Person
        v-for="(person, name, index) in assignation.people"
        :key="name"
        :degrees="degrees"
        :radius="radius"
        :index="index"
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
    p20radius() {
      return this.radius * .2
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
