<template>
  <div>
    <Map/>
  </div>
</template>

<script>
import Map from './components/Map.vue'

export default {
  name: 'App',
  components: {
    Map
  }
}
</script>

<style>
body {
  margin: 0;
  padding: 0;
  font-family: 'Roboto', sans-serif;
  background-color: black;
  color: white;
}
</style>