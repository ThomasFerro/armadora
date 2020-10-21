<script>
    import { createEventDispatcher } from 'svelte'
    import { i18n } from '../../i18n'
    export let palisadesLeft = 0
    export let palisadeSelection = []
    
    $: ongoingPalisadeSelection = palisadeSelection.length > 0

    const dispatch = createEventDispatcher()

    const validatePalisades = () => {
        dispatch('put-palisades', palisadeSelection)
        clearPalisades()
    }

    const clearPalisades = () => {
        dispatch('clear-palisades-selection')
    }

    const randomPosition = (max) => {
        return Math.floor(Math.random() * Math.floor(max));
    }
</script>
<svg
    class="palisades-left"
    xmlns="http://www.w3.org/2000/svg"
    version="1.1"
    width="100" height="100"
>
    <g stroke="#222831" stroke-width="1">
    {#each Array(palisadesLeft) as _}
        <rect x={randomPosition(90)} y={randomPosition(10)} class="palisade" width="10" height="40"/>
    {/each}
    </g>
</svg>
{#if ongoingPalisadeSelection}
<button on:click={validatePalisades}>{$i18n('palisadeSelection.validate')}</button>
<button on:click={clearPalisades}>{$i18n('palisadeSelection.clear')}</button>
{/if}

<style>
    .palisades-left {
        height: 100%
    }

    .palisade {
        fill: var(--palisade-color, saddlebrown);
    }
</style>
