<script>
    import { createEventDispatcher } from 'svelte'
    import Cell from './Cell.svelte'
    export let cells = []
    export let active = false

    const dispatch = createEventDispatcher()
    
    const cellClicked = (x, y) => {
        dispatch('cell-selected', {
            x,
            y,
        })
    }
</script>

<article class="grid">
    {#each cells as columns, y}
        {#each columns as cell, x}
        <Cell value={cell} {active} on:click={() => cellClicked(x, y)}></Cell>
        {/each}
    {/each}
</article>

<style>
    .grid {
        /* TODO: responsive */
        --cell-width: 100px;
        --palisade-width: 20px;
        display: grid;
        /* grid-template-columns: var(--cell-width) repeat(7, var(--palisade-width) var(--cell-width));
        grid-template-rows: var(--cell-width) repeat(4, var(--palisade-width) var(--cell-width)); */
        grid-template-columns: repeat(8, var(--cell-width));
        grid-template-rows: repeat(5, var(--cell-width));
    }
</style>