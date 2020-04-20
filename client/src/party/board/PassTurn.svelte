<script>
    import { createEventDispatcher } from 'svelte'
    let validationProcess = false

    const dipatch = createEventDispatcher()

    const passTurn = () => validationProcess = true
    const cancel = () => validationProcess = false
    const validate = () => dipatch('pass-turn')
</script>

{#if !validationProcess}
<button on:click={passTurn}>Pass your turn</button>
{:else}
<button on:click={validate}>Validate (You will not be able to play again for the rest of the game)</button>
<button on:click={cancel}>Cancel</button>
{/if}