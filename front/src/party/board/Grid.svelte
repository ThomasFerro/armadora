<script>
    import { createEventDispatcher } from 'svelte'
    import Cell from './Cell.svelte'
    import Palisade from './Palisade.svelte'
    export let cells = []
    export let palisades = []
    export let active = false

    const palisadeInSelection = (palisadeSelection, { x1, y1, x2, y2 }) => {
        return palisadeSelection.find(palisade => 
            (palisade.x1 === x1 && palisade.y1 === y1 && palisade.x2 === x2 && palisade.y2 === y2) ||
            (palisade.x1 === x2 && palisade.y1 === y2 && palisade.x2 === x1 && palisade.y2 === y1)
        )
    }

    $: palisadePresent = (palisadeToCheck) => {
        return palisadeInSelection(palisades, palisadeToCheck)
    }

    const dispatch = createEventDispatcher()
    
    const cellClicked = (x, y) => {
        dispatch('cell-selected', {
            x,
            y,
        })
    }

    let selectedPalisades = []

    $: palisadeSelectionOngoing = selectedPalisades.length > 0

    $: palisadeSelected = (palisadeToCheck) => {
        return palisadeInSelection(selectedPalisades, palisadeToCheck)
    }

    const borderSelected = (newPalisade) => {
        if (selectedPalisades.length < 2) {
            selectedPalisades = [
                ...selectedPalisades,
                newPalisade,
            ]
        }
    }

    const validatePalisades = () => {
        dispatch('border-selected', selectedPalisades)
        clearPalisades()
    }

    const clearPalisades = () => {
        selectedPalisades = []
    }
</script>

<article class="grid">
    {#each cells as columns, y}
        {#each columns as cell, x}
            <Cell value={cell} {active} on:click={() => cellClicked(x, y)}></Cell>
            {#if x < columns.length - 1}
                <Palisade 
                    active={active}
                    present={palisadePresent({ x1: x, y1: y, x2: x + 1, y2: y })}
                    selected={palisadeSelected({ x1: x, y1: y, x2: x + 1, y2: y })}
                    on:click={() => borderSelected({ x1: x, y1: y, x2: x + 1, y2: y })}
                    vertical
                ></Palisade>
            {/if}
        {/each}
        {#if y < cells.length - 1}
            {#each columns as horizontalPalisade, palisadeIndex}
                {#if palisadeIndex < columns.length}
                <Palisade
                    active={active}
                    present={palisadePresent({ x1: palisadeIndex, y1: y, x2: palisadeIndex, y2: y + 1 })}
                    selected={palisadeSelected({ x1: palisadeIndex, y1: y, x2: palisadeIndex, y2: y + 1 })}
                    on:click={() => borderSelected({ x1: palisadeIndex, y1: y, x2: palisadeIndex, y2: y + 1 })}
                ></Palisade>
                {/if}
                <!-- FIXME: Solution withtout blank div -->
                {#if palisadeIndex < columns.length - 1}
                    <div class="blank"></div>
                {/if}
            {/each}
        {/if}
    {/each}
</article>
{#if palisadeSelectionOngoing}
<button on:click={validatePalisades}>Validate palisades</button>
<button on:click={clearPalisades}>Clear palisades</button>
{/if}

<style>
    .grid {
        /* TODO: responsive */
        --cell-width: 100px;
        --palisade-width: 20px;
        display: grid;
        grid-template-columns: var(--cell-width) repeat(7, var(--palisade-width) var(--cell-width));
        grid-template-rows: var(--cell-width) repeat(4, var(--palisade-width) var(--cell-width));
    }
</style>