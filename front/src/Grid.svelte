<script>
    import { createEventDispatcher } from 'svelte';
    import Cell from './Cell.svelte';
    import { WARRIORS } from './editMode';
    import { LAND } from './cellTypes';

    const dispatch = createEventDispatcher();

    export let value;
    export let mode;

    const landCell = (cell) => cell && cell.type === LAND

    $: editMode = (x, y) => {        
        const cell = value[x][y]
        if (mode === WARRIORS &&
            landCell(cell)) {

            return mode
        }

        return ''
    }

    const cellClicked = (x, y) => {
        dispatch('cell-clicked', { x, y })
    }
</script>

<section class="grid">
    {#each value as line, lineIndex}
        {#each line as cell, cellIndex}
            <Cell
                value={cell}
                editMode={editMode(lineIndex, cellIndex)}
                on:click={() => cellClicked(lineIndex, cellIndex)}
            ></Cell>
        {/each}
    {/each}
</section>

<style>
    .grid {
        --grid-width: 50px;
        --palisade-width: 2px;

        display: grid;
        grid-template-columns: repeat(8, var(--grid-width));
        grid-template-rows: repeat(5, var(--grid-width));
    }
</style>
