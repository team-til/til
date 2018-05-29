<template>
  <div id="app">
    <Home msg="Team TIL - Today I Learned"/>
  </div>
</template>

<script>
import Home from './components/Home.vue'

export default {
  name: 'app',
  components: {
    Home
  }
}
</script>

<style lang="scss">
.clearfix {
  overflow: auto;
}

.clearfix::after {
  content: "";
  clear: both;
  display: table;
}
</style>
