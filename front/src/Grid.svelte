<script>
    import Cell from './Cell.svelte';

    export let grid;
</script>

<section class="grid">
    {#each grid as line}
        {#each line as cell}
            <Cell cell={cell}></Cell>
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
