<script>
    import { createEventDispatcher } from 'svelte';
    import Cell from './Cell.svelte';
    import Palisade from './Palisade.svelte';
    import { WARRIORS } from './editModes';

    const dispatch = createEventDispatcher();

    export let value;
    export let mode;

    $: lastHorizontalPalisade = (index) => {
        return index === value[0].length - 1
    }

    $: lastVerticalPalisade = (index) => {
        return index + 1 === value[0].length - 1
    }

    const cellClicked = (x, y) => {
        dispatch('cell-clicked', { x, y })
    }
</script>

<!-- TODO: Extract palisade -->

<section class="grid">
    {#each value as line, lineIndex}
        {#each line as cell, cellIndex}
            <Cell
                value={cell}
                {mode}
                on:click={() => cellClicked(lineIndex, cellIndex)}
            ></Cell>
            {#if cellIndex < line.length - 1}
            <Palisade
                vertical
                last={lastVerticalPalisade(cellIndex)}
            ></Palisade>
            {/if}
        {/each}
        {#each line as horizontalPalisade, palisadeIndex}
            {#if palisadeIndex < line.length}
            <Palisade
                last={lastHorizontalPalisade(palisadeIndex)}
            ></Palisade>
            {/if}
        {/each}
    {/each}
</section>

<style>
    .grid {
        /* TODO: Cell-width */
        --grid-width: 100px;
        --palisade-width: 20px;

        display: grid;
        grid-template-columns: var(--grid-width) repeat(7, var(--palisade-width) var(--grid-width));
        grid-template-rows: var(--grid-width) repeat(4, var(--palisade-width) var(--grid-width));
    }
</style>
