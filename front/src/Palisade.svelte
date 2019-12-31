<script>
    export let vertical;
    export let last;

    $: lastVerticalPalisade = () => {
        return vertical && last
    }

    $: lastHorizontalPalisade = () => {
        return !vertical && last
    }
</script>

<button
    class="palisade"
    class:vertical-palisade={vertical}
    class:horizontal-palisade={!vertical}
    class:last-vertical-palisade={lastVerticalPalisade()}
    class:last-horizontal-palisade={lastHorizontalPalisade()}
></button>

<style>
.palisade {
    margin: 0;
    padding: 0;
}

.horizontal-palisade {
    width: 100%;
    height: var(--palisade-width);
}

.horizontal-palisade:not(.last-horizontal-palisade) {
    grid-column: span 2;
}
</style>
