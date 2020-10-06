<!-- TODO: Extract the validation process -->
<script>
    import { createEventDispatcher } from 'svelte'
    let validationProcess = false

    const dipatch = createEventDispatcher()

    const passTurn = () => validationProcess = true
    const cancel = () => validationProcess = false
    const validate = () => dipatch('pass-turn')
</script>

{#if !validationProcess}
<button class="player-action" on:click={passTurn}>Pass your turn</button>
{:else}
<button class="player-action" on:click={validate}>Validate (You will not be able to play again for the rest of the game)</button>
<button class="player-action" on:click={cancel}>Cancel</button>
{/if}
