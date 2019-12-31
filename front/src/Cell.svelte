<script>
    import { LAND } from './cellTypes';
    import { PALISADES, WARRIORS } from './editModes';

    export let value;
    export let mode;

    const landCell = (cell) => cell && cell.type === LAND

    $: editMode = (() => {        
        if (mode === WARRIORS &&
            !landCell(value)) {

            return ''
        }

        return mode
    })()
    $: cellClasses = `cell ${editMode}`
    $: disabled = editMode !== WARRIORS || value.type != LAND || value.warrior
    $: cellValue = value.warrior ? value.warrior.playerDisplayName : value.type
</script>

<button
    class={cellClasses}
    class:palisade-left={value.palisades.left}
    class:palisade-top={value.palisades.top}
    class:palisade-right={value.palisades.right}
    class:palisade-bottom={value.palisades.bottom}
    {disabled}
    on:click
>{cellValue}</button>

<style>
    .cell {
        border: 1px solid black;
        --palisade-width: 2px;
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
