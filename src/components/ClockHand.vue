<template>
  <li class="clock-hand">
    {{ name }} -> {{ location}}
  </li>
</template>

<script>
export default {
  props: {
    name: String,
    location: String
  }
}
</script>

<style scoped>
.clock-hand {
  display: inline-block;
  margin: 1rem 0;
  width: 100%;
}
</style>
