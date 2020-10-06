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
    class={`palisade ${area} player-action`}
    style={`grid-area: ${area}`}
    class:palisade--present={present}
    class:selected={selected}
></button>

<style>
.palisade {
    margin: 0;
    padding: 0;
    height: 100%;
    width: 100%;
    border-radius: 5px;
}
.palisade--present {
    background-color: var(--palisade-color, saddlebrown);
}
.selected {
    background-color: gray;
}
</style>