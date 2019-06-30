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
    <button @click="addTechInteraction">Add Tech</button>
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
        addTechInteraction() {
            console.log('Add tech!');
        }
    }
}
</script>

