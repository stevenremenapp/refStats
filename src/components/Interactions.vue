<template>
    <div>
    <h1>Interactions</h1>
    <ul>
        <li v-for="interaction in interactions" :key="interaction.id">
            <span>ID: {{ interaction.id }}</span><br>
            <span>Type: {{ interaction.type }}</span><br>
            <span>Time: {{ interaction.time }}</span>
            <button @click="$emit('delete:interaction', interaction.id)">Delete</button>
        </li>
    </ul>
    <button @click="submitInteraction" name="technology">Add Tech</button>
    <button @click="submitInteraction" name="reference">Add Reference</button>
    </div>
</template>

<script>
export default {
    name: 'interactions',
    props: {
        interactions: Array
    },
    methods: {
        submitInteraction() {
            let time = new Date().toLocaleString('en-US', { hour: 'numeric', minute: 'numeric', hour12: true})
            let type = event.target.name === "technology" ? "technology" : "reference"
            let interaction = {
                type: type,
                time: time
            }

            this.$emit('add:interaction', interaction)
        }
    }
}
</script>

