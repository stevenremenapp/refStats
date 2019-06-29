<template>
  <div id="app">
    <Header />
    <h1>Interactions</h1>
    <ul>
      <li v-for="interaction in interactions" v-bind:key="interaction.id">
        <span>ID: {{ interaction.id }}</span><br>
        <span>Type: {{ interaction.type }}</span><br>
        <span>Time: {{ interaction.time }}</span>
      </li>
    </ul>
  </div>
</template>

<script>
import Header from './components/Header'

export default {
  name: 'app',
  components: {
    Header
  },
  data() {
    return {
      interactions: []
    }
  },
  created() {
    fetch('http://localhost:5000/interactions')
    .then(response => response.json())
    .then(json => {
      this.interactions = json
    })
  }
}
</script>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
</style>
