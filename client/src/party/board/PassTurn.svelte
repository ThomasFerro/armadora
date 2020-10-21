<!-- TODO: Extract the validation process -->
<script>
    import { createEventDispatcher } from 'svelte'
    import { i18n } from '../../i18n'

    let validationProcess = false

    const dipatch = createEventDispatcher()

    const passTurn = () => validationProcess = true
    const cancel = () => validationProcess = false
    const validate = () => dipatch('pass-turn')
</script>

{#if !validationProcess}
<button class="player-action" on:click={passTurn}>{$i18n('passTurn.pass')}</button>
{:else}
<button class="player-action" on:click={validate}>{$i18n('passTurn.validate')}</button>
<button class="player-action" on:click={cancel}>{$i18n('passTurn.cancel')}</button>
{/if}
