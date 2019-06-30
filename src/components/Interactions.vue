<template>
    <div>
    <h1>Interactions</h1>
    <ul>
        <li v-for="interaction in interactions" :key="interaction.id">
            <span>ID: {{ interaction.id }}</span><br>
            <span>Type: {{ interaction.type }}</span><br>
            <span>Time: {{ interaction.time }}</span>
        </li>
    </ul>
    <button @click="submitTechInteraction">Add Tech</button>
    <button>Add Reference</button>
    </div>
</template>

<script>
export default {
    name: 'interactions',
    props: {
        interactions: Array
    },
    methods: {
        submitTechInteraction() {
            let time = new Date().toLocaleString('en-US', { hour: 'numeric', minute: 'numeric', hour12: true})
            let type = "technology"
            let techInteraction = {
                type: type,
                time: time
            }
            console.log(techInteraction)

            this.$emit('add:techInteraction', techInteraction)
        }
    }
}
</script>

