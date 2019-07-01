<template>
    <div>
        <h2>Interactions</h2>
        <!-- <ul>
            <li v-for="interaction in interactions" :key="interaction.id">
                <span>ID: {{ interaction.id }}</span><br>
                <span>Type: {{ interaction.type }}</span><br>
                <span>Time: {{ interaction.time }}</span>
                <button @click="$emit('delete:interaction', interaction.id)">Delete</button>
            </li>
        </ul> -->
        <div class="interactionDisplay">
            <div class="technologyInteractionDisplay">
                <h3>Technology (Total: {{ technologyInteractions.length }})</h3>
                <button @click="submitInteraction" name="technology" class="addInteractionButton">Add Tech</button>
                <div
                v-for="interaction in technologyInteractions"
                :key="interaction.id"
                class="interactionInfoCard">
                    <div class="interactionInfo">
                        <p>ID: {{ interaction.id }}</p>
                        <p class="type">Type: {{ interaction.type }}</p>
                        <p>Time: {{ interaction.time }}</p>
                    </div>
                    <button @click="$emit('delete:interaction', interaction.id)">Delete</button>
                </div>
            </div>
            <div class="referenceInteractionDisplay">
                <h3>Reference (Total: {{ referenceInteractions.length }})</h3>
                <button @click="submitInteraction" name="reference" class="addInteractionButton">Add Reference</button>
                <div
                v-for="interaction in referenceInteractions"
                :key="interaction.id"
                class="interactionInfoCard">
                    <div class="interactionInfo">
                        <p>ID: {{ interaction.id }}</p>
                        <p class="type">Type: {{ interaction.type }}</p>
                        <p>Time: {{ interaction.time }}</p>
                    </div>
                    <button @click="$emit('delete:interaction', interaction.id)">Delete</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'interactions',
    props: {
        interactions: Array
    },
    computed: {
        technologyInteractions() {
            return this.interactions.filter(interaction =>  interaction.type === "technology")
        },
        referenceInteractions() {
            return this.interactions.filter(interaction => interaction.type === "reference")
        }
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

<style lang="scss" scoped>
h2 {
    text-align: center;
}

.interactionDisplay {
    display: flex;
    justify-content: center;
    .technologyInteractionDisplay, .referenceInteractionDisplay {
        margin: 0 30px;
    }
    h3 {
        text-align: center;
        margin-bottom: 5px;
    }
    .addInteractionButton {
        width: 100%;
        border: none;
        border-radius: 5px;
        padding: 5px;
        background: #16940a;
        color: #fff;
        &:hover {
            cursor: pointer;
        }
    }
}

.interactionInfoCard {
    display: flex;
    padding: 10px;
    margin-bottom: 10px;
    border: 1px solid rgba(0, 0, 0, 0.3);
    border-radius: 5px;
    &:first-of-type {
        margin-top: 10px;
    }
    .interactionInfo {
        min-width: 150px;
        p {
            margin: 0;
        }
        .type {
            text-transform: capitalize;
        }
    }
    button {
        align-self: center;
        border: none;
        border-radius: 5px;
        background: #d60909;
        color: #fff;
        padding: 5px;
        &:hover {
            cursor: pointer;
        }
    }
}
</style>
