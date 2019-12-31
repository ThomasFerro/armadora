<script>
    import { createEventDispatcher } from 'svelte';
    import Cell from './Cell.svelte';
    import { WARRIORS } from './editModes';

    const dispatch = createEventDispatcher();

    export let value;
    export let mode;

    const cellClicked = (x, y) => {
        dispatch('cell-clicked', { x, y })
    }
</script>

<section class="grid">
    {#each value as line, lineIndex}
        {#each line as cell, cellIndex}
            <Cell
                value={cell}
                {mode}
                on:click={() => cellClicked(lineIndex, cellIndex)}
            ></Cell>
        {/each}
    {/each}
</section>

<style>
    .grid {
        --grid-width: 50px;

        display: grid;
        grid-template-columns: repeat(8, var(--grid-width));
        grid-template-rows: repeat(5, var(--grid-width));
    }
</style>
