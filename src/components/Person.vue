<template>
  <g :style="handStyle">
    <line
      x1=0
      y1=0
      x2=0
      :y2="needleLength"
      :style="needleStyle"
    />
    <circle
      cx=0
      :cy="personPosition"
      :r="personRadius"
      :style="personStyle"
    />
    <text
      text-anchor="middle"
      alignment-baseline="middle"
      x=0
      :y="personPosition"
      :style="nameStyle"
    >
      {{ name.charAt(0).toUpperCase() }}
    </text>
  </g>
</template>

<script>
export default {
  props: {
    index: Number,
    name: String,
    degrees: Number,

    radius: Number,
    darkColor: String,
    lightColor: String
  },
  computed: {
    handStyle() {
      return `transform: rotate(${this.degrees}deg)`
    },

    needleLength() {
      return this.radius * -.7
    },
    needleStyle() {
      const strokeWidth = this.radius * .02
      return `stroke: ${this.darkColor}; stroke-width: ${strokeWidth}px`
    },

    personPosition() {
      const shiftedPosition = this.index * this.personRadius * 1.4
      return this.radius * -.6 + shiftedPosition
    },
    personRadius() {
      return this.radius * .08
    },
    personStyle() {
      return `fill: ${this.darkColor}`
    },

    nameStyle() {
      const fontSize = this.radius * .06
      return `fill: ${this.lightColor}; font-size: ${fontSize}px; font-weight: bold`
    }
  }
}
</script>
