<script>
    export let value = {}

    $: cells = value && value.cells || []
</script>

<article class="grid">
    {#each cells as columns}
        {#each columns as cell}
        <section class="cell">
            {cell.type}
        </section>
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
    
    .cell {
        width: 100%;
        height: 100%;
        border: 1px solid black;
        margin: 0;
        padding: 0;
    }
</style>
