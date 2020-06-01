<template>
  <g>
    <line
      :style="line"
      :x1="radius"
      :y1="radius"
      :x2="radius"
      :y2="p30radius"
      :transform="rotation"
    />
    <ellipse
      :style="ellipse"
      :cx="radius"
      :cy="shiftedp40radius"
      :rx="p8radius"
      :ry="p8radius"
      :transform="rotation"
    />
    <text
      :style="text"
      text-anchor="middle"
      alignment-baseline="middle"
      :x="radius"
      :y="shiftedp40radius"
      :transform="rotation">
      {{ name.charAt(0).toUpperCase() }}
    </text>
  </g>
</template>

<script>
export default {
  props: {
    degrees: Number,
    radius: Number,
    index: Number,
    name: String,
    dark: String,
    light: String
  },
  computed: {
    line() {
      return "stroke: #" + this.dark + "; stroke-width: " + this.radius * .01 + "px"
    },
    ellipse() {
      return "fill: #" + this.dark
    },
    text() {
      return "fill: #" + this.light + "; font-weight: bold"
    },
    p8radius() {
      return this.radius * .08
    },
    p30radius() {
      return this.radius * .3
    },
    shiftedp40radius() {
      return this.radius * .4 + this.index * this.p8radius * 1.4
    },
    rotation() {
      return "rotate(" + this.degrees + " " + this.radius + " " + this.radius + ")"
    }
  }
}
</script>
