<template>
  <g>
    <g v-if="Object.keys(status.people).length > 0">
      <Person
        v-for="(person, name, index) in status.people"
        :key="name"
        :index="index"
        :name="person.name"
        :degrees="degrees"
        :radius="radius"
        :darkColor="darkColor"
        :lightColor="lightColor"
      />
    </g>

    <text
      x=0
      :y="statusPosition"
      :transform="statusRotation"
      text-anchor="middle"
      :style="statusStyle"
    >
      {{ status.status }}
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
    status: Object,
    index: Number,
    elCount: Number,

    radius: Number,
    darkColor: String,
    lightColor: String
  },
  computed: {
    degrees() {
      return this.index * 360 / this.elCount
    },

    statusPosition() {
      return this.radius * -.8
    },
    statusStyle() {
      const fontSize = this.radius * .1
      return `fill: ${this.darkColor}; font-size: ${fontSize}px; font-weight: bold`
    },
    statusRotation() {
      return `rotate(${this.degrees})`
    }
  }
}
</script>
