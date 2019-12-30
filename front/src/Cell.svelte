<script>
    export let value;
</script>

<!-- TODO: Change tag ? -->
<div
    class="cell"
    class:palisade-left={value.palisades.left}
    class:palisade-top={value.palisades.top}
    class:palisade-right={value.palisades.right}
    class:palisade-bottom={value.palisades.bottom}
>{value.type}</div>

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
