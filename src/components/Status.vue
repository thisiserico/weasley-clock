<template>
  <text
    x=0
    :y="statusPosition"
    :transform="statusRotation"
    text-anchor="middle"
    :style="statusStyle"
  >
    {{ status }}
  </text>
</template>

<script>

export default {
  props: {
    status: String,
    rotation: Number,

    radius: Number,
    darkColor: String,
    lightColor: String
  },
  computed: {
    statusPosition() {
      return this.radius * -.8
    },
    statusStyle() {
      const fontSize = this.radius * .1
      return `fill: ${this.darkColor}; font-size: ${fontSize}px; font-weight: bold`
    },
    statusRotation() {
      return `rotate(${this.rotation})`
    }
  }
}
</script>
