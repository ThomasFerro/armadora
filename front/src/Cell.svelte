<script>
    export let cell;
</script>

<!-- TODO: Change tag -->
<div
    class="cell"
    class:palisade-left={cell.palisades.left}
    class:palisade-top={cell.palisades.top}
    class:palisade-right={cell.palisades.right}
    class:palisade-bottom={cell.palisades.bottom}
>{cell.type}</div>

<style>
    .cell {
        border: 1px solid black;
    }

    .palisade-left {
        border-left: var(--palisade-width) solid burlywood;
    }

    .palisade-top {
        border-top: var(--palisade-width) solid burlywood;
    }

    .palisade-right {
        border-right: var(--palisade-width) solid burlywood;
    }

    .palisade-bottom {
        border-bottom: var(--palisade-width) solid burlywood;
    }
</style>
