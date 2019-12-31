<script>
    import { LAND } from './cellTypes';

    export let value;
    export let editMode;

    $: cellClasses = `cell ${editMode}`
    // TODO : || cell already taken
    $: disabled = !editMode || value.type != LAND
</script>

<button
    class={cellClasses}
    class:palisade-left={value.palisades.left}
    class:palisade-top={value.palisades.top}
    class:palisade-right={value.palisades.right}
    class:palisade-bottom={value.palisades.bottom}
    {disabled}
    on:click
>{value.type}</button>

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
