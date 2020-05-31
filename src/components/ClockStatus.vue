<template>
  <li class="clock-status">
    {{ status }}
  </li>
</template>

<script>
export default {
  props: {
    status: String
  }
}
</script>

<style scoped>
.clock-status {
  display: inline-block;
  margin: 1rem 0;
  width: 100%;
}
</style>
