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
            <Cell value={cell} {active} on:click={() => cellClicked(x, y)} area={`cell--${x}-${y}`}></Cell>
            {#if x < columns.length - 1}
                <Palisade 
                    active={active}
                    present={palisadePresent({ x1: x, y1: y, x2: x + 1, y2: y })}
                    selected={palisadeSelected({ x1: x, y1: y, x2: x + 1, y2: y })}
                    on:click={() => borderSelected({ x1: x, y1: y, x2: x + 1, y2: y })}
                    area={`palisade--${x}-${y}-${x+1}-${y}`}
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
                    area={`palisade--${palisadeIndex}-${y}-${palisadeIndex}-${y + 1}`}
                ></Palisade>
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
    /* FIXME: Not convinced with this way of writing the grid but cannot find another working way... */
    .grid {
        margin: 0.5em;
        --cell-width: 4fr;
        --palisade-width: 1fr;
        --grid-template:
            "cell--0-0          palisade--0-0-1-0   cell--1-0           palisade--1-0-2-0   cell--2-0           palisade--2-0-3-0   cell--3-0           palisade--3-0-4-0   cell--4-0           palisade--4-0-5-0   cell--5-0           palisade--5-0-6-0   cell--6-0           palisade--6-0-7-0   cell--7-0" var(--cell-width)
            "palisade--0-0-0-1  .                   palisade--1-0-1-1   .                   palisade--2-0-2-1   .                   palisade--3-0-3-1   .                   palisade--4-0-4-1   .                   palisade--5-0-5-1   .                   palisade--6-0-6-1   .                   palisade--7-0-7-1" var(--palisade-width)
            "cell--0-1          palisade--0-1-1-1   cell--1-1           palisade--1-1-2-1   cell--2-1           palisade--2-1-3-1   cell--3-1           palisade--3-1-4-1   cell--4-1           palisade--4-1-5-1   cell--5-1           palisade--5-1-6-1   cell--6-1           palisade--6-1-7-1   cell--7-1" var(--cell-width)
            "palisade--0-1-0-2  .                   palisade--1-1-1-2   .                   palisade--2-1-2-2   .                   palisade--3-1-3-2   .                   palisade--4-1-4-2   .                   palisade--5-1-5-2   .                   palisade--6-1-6-2   .                   palisade--7-1-7-2" var(--palisade-width)
            "cell--0-2          palisade--0-2-1-2   cell--1-2           palisade--1-2-2-2   cell--2-2           palisade--2-2-3-2   cell--3-2           palisade--3-2-4-2   cell--4-2           palisade--4-2-5-2   cell--5-2           palisade--5-2-6-2   cell--6-2           palisade--6-2-7-2   cell--7-2" var(--cell-width)
            "palisade--0-2-0-3  .                   palisade--1-2-1-3   .                   palisade--2-2-2-3   .                   palisade--3-2-3-3   .                   palisade--4-2-4-3   .                   palisade--5-2-5-3   .                   palisade--6-2-6-3   .                   palisade--7-2-7-3" var(--palisade-width)
            "cell--0-3          palisade--0-3-1-3   cell--1-3           palisade--1-3-2-3   cell--2-3           palisade--2-3-3-3   cell--3-3           palisade--3-3-4-3   cell--4-3           palisade--4-3-5-3   cell--5-3           palisade--5-3-6-3   cell--6-3           palisade--6-3-7-3   cell--7-3" var(--cell-width)
            "palisade--0-3-0-4  .                   palisade--1-3-1-4   .                   palisade--2-3-2-4   .                   palisade--3-3-3-4   .                   palisade--4-3-4-4   .                   palisade--5-3-5-4   .                   palisade--6-3-6-4   .                   palisade--7-3-7-4" var(--palisade-width)
            "cell--0-4          palisade--0-4-1-4   cell--1-4           palisade--1-4-2-4   cell--2-4           palisade--2-4-3-4   cell--3-4           palisade--3-4-4-4   cell--4-4           palisade--4-4-5-4   cell--5-4           palisade--5-4-6-4   cell--6-4           palisade--6-4-7-4   cell--7-4" var(--cell-width) /
            var(--cell-width) var(--palisade-width) var(--cell-width) var(--palisade-width) var(--cell-width) var(--palisade-width) var(--cell-width) var(--palisade-width) var(--cell-width) var(--palisade-width) var(--cell-width) var(--palisade-width) var(--cell-width) var(--palisade-width) var(--cell-width);
        display: grid;
        grid-template: var(--grid-template);  
    }

    @media (orientation:portrait) {
        .grid {
            --grid-template:
            "cell--0-4          palisade--0-3-0-4   cell--0-3           palisade--0-2-0-3   cell--0-2           palisade--0-1-0-2   cell--0-1           palisade--0-0-0-1   cell--0-0" var(--cell-width)
            "palisade--0-4-1-4  .                   palisade--0-3-1-3   .                   palisade--0-2-1-2   .                   palisade--0-1-1-1   .                   palisade--0-0-1-0" var(--palisade-width)
            "cell--1-4          palisade--1-3-1-4   cell--1-3           palisade--1-2-1-3   cell--1-2           palisade--1-1-1-2   cell--1-1           palisade--1-0-1-1   cell--1-0" var(--cell-width)
            "palisade--1-4-2-4  .                   palisade--1-3-2-3   .                   palisade--1-2-2-2   .                   palisade--1-1-2-1   .                   palisade--1-0-2-0" var(--palisade-width)
            "cell--2-4          palisade--2-3-2-4   cell--2-3           palisade--2-2-2-3   cell--2-2           palisade--2-1-2-2   cell--2-1           palisade--2-0-2-1   cell--2-0" var(--cell-width)
            "palisade--2-4-3-4  .                   palisade--2-3-3-3   .                   palisade--2-2-3-2   .                   palisade--2-1-3-1   .                   palisade--2-0-3-0" var(--palisade-width)
            "cell--3-4          palisade--3-3-3-4   cell--3-3           palisade--3-2-3-3   cell--3-2           palisade--3-1-3-2   cell--3-1           palisade--3-0-3-1   cell--3-0" var(--cell-width)
            "palisade--3-4-4-4  .                   palisade--3-3-4-3   .                   palisade--3-2-4-2   .                   palisade--3-1-4-1   .                   palisade--3-0-4-0" var(--palisade-width)
            "cell--4-4          palisade--4-3-4-4   cell--4-3           palisade--4-2-4-3   cell--4-2           palisade--4-1-4-2   cell--4-1           palisade--4-0-4-1   cell--4-0" var(--cell-width)
            "palisade--4-4-5-4  .                   palisade--4-3-5-3   .                   palisade--4-2-5-2   .                   palisade--4-1-5-1   .                   palisade--4-0-5-0" var(--palisade-width)
            "cell--5-4          palisade--5-3-5-4   cell--5-3           palisade--5-2-5-3   cell--5-2           palisade--5-1-5-2   cell--5-1           palisade--5-0-5-1   cell--5-0" var(--cell-width)
            "palisade--5-4-6-4  .                   palisade--5-3-6-3   .                   palisade--5-2-6-2   .                   palisade--5-1-6-1   .                   palisade--5-0-6-0" var(--palisade-width)
            "cell--6-4          palisade--6-3-6-4   cell--6-3           palisade--6-2-6-3   cell--6-2           palisade--6-1-6-2   cell--6-1           palisade--6-0-6-1   cell--6-0" var(--cell-width)
            "palisade--6-4-7-4  .                   palisade--6-3-7-3   .                   palisade--6-2-7-2   .                   palisade--6-1-7-1   .                   palisade--6-0-7-0" var(--palisade-width)
            "cell--7-4          palisade--7-3-7-4   cell--7-3           palisade--7-2-7-3   cell--7-2           palisade--7-1-7-2   cell--7-1           palisade--7-0-7-1   cell--7-0" var(--cell-width) /
            var(--cell-width) var(--palisade-width) var(--cell-width) var(--palisade-width) var(--cell-width) var(--palisade-width) var(--cell-width) var(--palisade-width) var(--cell-width);
        }
    }
</style>