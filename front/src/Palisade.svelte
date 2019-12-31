<script>
    import { PALISADES } from './editModes';

    export let value;
    export let vertical;
    export let last;
    export let mode;

    $: lastVerticalPalisade = () => {
        return vertical && last
    }

    $: lastHorizontalPalisade = () => {
        return !vertical && last
    }

    $: disabled = mode !== PALISADES || (value && value.disabled)

    $: active = value && value.active
</script>

<button
    {disabled}
    on:click
    class="palisade"
    class:active={active}
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
</style>
