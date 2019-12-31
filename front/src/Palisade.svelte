<script>
    import { PALISADES } from './editModes';

    export let present = undefined;
    export let disabled = false;
    export let selected = false;
    export let vertical;
    export let last;
    export let mode;

    $: lastVerticalPalisade = () => {
        return vertical && last
    }

    $: lastHorizontalPalisade = () => {
        return !vertical && last
    }

    $: isDisabled = mode !== PALISADES || disabled || present
</script>

<button
    disabled={isDisabled}
    on:click
    class="palisade"
    class:present={present}
    class:selected={selected}
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

.present {
    background-color: burlywood;
}

.selected {
    background-color: gray;
}
</style>
