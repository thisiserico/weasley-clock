<template>
  <g>
    <path :id="status" fill="none" :d="pathPoints" />

    <text :style="statusStyle">
      <textPath :xlink:href="statusID" text-anchor="middle" startOffset="50%">
        {{ status }}
      </textPath>
    </text>
  </g>
</template>

<script>
import { arcPath } from '../lib/path'

export default {
  props: {
    status: String,
    rotation: Number,
    sliceDegrees: Number,

    radius: Number,
    darkColor: String,
    lightColor: String
  },
  computed: {
    pathPoints() {
      const radius = this.radius * .9
      const startAngle = -1 * this.sliceDegrees / 2 + this.rotation
      const endAngle = this.sliceDegrees / 2 + this.rotation

      return arcPath(0, 0, radius, startAngle, endAngle)
    },
    statusID() {
      return `#${this.status}`
    },

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
