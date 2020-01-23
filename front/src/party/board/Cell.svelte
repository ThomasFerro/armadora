<script>
    export let value = {}
    export let active = false

    $: cellValue = () => {
        if (value.type === 'warrior') {
            return value.character
        }
        if (value.gold) {
            return value.gold
        }
        return ''
    }

    $: disabled = !active || value.type !== 'land'
</script>

<button class="cell" on:click {disabled}>
    {cellValue()}
</button>

<style>
    .cell {
        width: 100%;
        height: 100%;
        border: 1px solid black;
        margin: 0;
        padding: 0;
    }
</style>
