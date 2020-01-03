<script>
    import { createEventDispatcher } from 'svelte';
    import Cell from './cell/Cell.svelte';
    import Palisade from './Palisade.svelte';
    import { WARRIORS } from './editModes';

    const dispatch = createEventDispatcher();

    export let grid;
    export let mode;
    export let selectedPalisades = []

    $: lastHorizontalPalisade = (index) => {
        return index === grid[0].length - 1
    }

    $: lastVerticalPalisade = (index) => {
        return index + 1 === grid[0].length - 1
    }

    $: isPalisadePresent = (x, y, vertical) => {
        if (vertical) {
            return grid[x][y].palisades.right
        }
        return grid[x][y].palisades.bottom
    }

    $: isPalisadeSelected = (x, y, vertical) => {
        return selectedPalisades.find(palisade => {
            return palisade.x === x && palisade.y === y && palisade.vertical === vertical
        })
    }

    const cellClicked = (x, y) => {
        dispatch('cell-clicked', { x, y })
    }

    const palisadeClicked = (x, y, vertical) => {
        dispatch('palisade-clicked', { x, y, vertical })
    }
</script>

<section class="grid">
    <!-- TODO: Change line by row and invert everywhere -->
    {#each grid as line, lineIndex}
        {#each line as cell, cellIndex}
            <Cell
                value={cell}
                {mode}
                on:click={() => cellClicked(lineIndex, cellIndex)}
            ></Cell>
            {#if cellIndex < line.length - 1}
            <Palisade
                present={isPalisadePresent(lineIndex, cellIndex, true)}
                selected={isPalisadeSelected(lineIndex, cellIndex, true)}
                vertical={true}
                {mode}
                last={lastVerticalPalisade(cellIndex)}
                on:click={() => palisadeClicked(lineIndex, cellIndex, true)}
            ></Palisade>
            {/if}
        {/each}
        {#if lineIndex < grid.length - 1}
            {#each line as horizontalPalisade, cellIndex}
                {#if cellIndex < line.length}
                <Palisade
                    present={isPalisadePresent(lineIndex, cellIndex)}
                selected={isPalisadeSelected(lineIndex, cellIndex)}
                    vertical={false}
                    {mode}
                    last={lastHorizontalPalisade(cellIndex)}
                    on:click={() => palisadeClicked(lineIndex, cellIndex)}
                ></Palisade>
                {/if}

                {#if cellIndex < line.length - 1}
                <div class="blank"></div>
                {/if}
            {/each}
        {/if}
    {/each}
</section>

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
