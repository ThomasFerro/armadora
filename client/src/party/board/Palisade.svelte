<script>
    export let vertical = false;
    export let active = false;
    export let present = false;
    export let selected = false;

    $: disabled = !active || present || selected
</script>

<button
    disabled={disabled}
    on:click
    class="palisade"
    class:vertical-palisade={vertical}
    class:horizontal-palisade={!vertical}
    class:present={present}
    class:selected={selected}
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