<template>
  <g>
    <line
      v-if="Object.keys(assignation.people).length > 0"
      :style="line"
      :x1="radius"
      :y1="radius"
      :x2="radius"
      :y2="p30radius"
      :transform="rotation"
    />
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
export default {
  props: {
    assignation: Object,
    index: Number,
    elements: Number,
    radius: Number,
    dark: String
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
    rotation() {
      const degrees = this.index * 360 / this.elements

      return "rotate(" + degrees + " " + this.radius + " " + this.radius + ")"
    }
  }
}
</script>
