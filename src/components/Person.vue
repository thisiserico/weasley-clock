<template>
  <g>
    <ellipse
      :style="ellipse"
      :cx="radius"
      :cy="p40p80radius"
      :rx="p8radius"
      :ry="p8radius"
      :transform="rotation"
    />
    <text
      :style="text"
      text-anchor="middle"
      alignment-baseline="middle"
      :x="radius"
      :y="p40p80radius"
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
    name: String,
    dark: String,
    light: String
  },
  computed: {
    ellipse() {
      return "fill: #" + this.dark
    },
    text() {
      return "fill: #" + this.light + "; font-weight: bold"
    },
    p8radius() {
      return this.radius * .08
    },
    p40p80radius() {
      const rand = Math.random() * 40 / 100
      return this.radius * (0.4 + rand)
    },
    rotation() {
      return "rotate(" + this.degrees + " " + this.radius + " " + this.radius + ")"
    }
  }
}
</script>
