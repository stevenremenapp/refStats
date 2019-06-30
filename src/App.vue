<template>
  <div id="app">
    <Header />
    <Interactions
      :interactions="interactions"
      @add:techInteraction="addTechInteraction"
    />
  </div>
</template>

<script>
import Header from './components/Header'
import Interactions from './components/Interactions'

export default {
  name: 'app',
  components: {
    Header,
    Interactions
  },
  data() {
    return {
      interactions: []
    }
  },
  mounted() {
    this.getInteractions()
  },
  methods: {
    async getInteractions() {
      try {
        const response = await fetch(`${process.env.VUE_APP_ROOT_API}/interactions`)
        const data = await response.json()
        this.interactions = data
      } catch (error) {
        console.error(error)
      }
    },

    async addTechInteraction(techInteraction) {
      try {
        const response = await fetch(`${process.env.VUE_APP_ROOT_API}/interactions`, {
          method: 'POST',
          body: JSON.stringify(techInteraction),
          headers: { "Content-type": "application/json; charset=UTF-8" }
        })
        const data = await response.json()
        this.interactions = [...this.interactions, data]
      } catch (error) {
        console.error(error)
      }
    }
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
