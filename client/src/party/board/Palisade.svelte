<script>
    export let active = false;
    export let present = false;
    export let selected = false;
    export let area = "";

    $: disabled = !active || present || selected
</script>

<button
    disabled={disabled}
    on:click
    class={`palisade ${area}`}
    style={`grid-area: ${area}`}
    class:present={present}
    class:selected={selected}
></button>

<style>
.palisade {
    margin: 0;
    padding: 0;
    height: 100%;
    width: 100%;
}
.present {
    background-color: var(--palisade-present, saddlebrown);
}
.selected {
    background-color: gray;
}
</style>