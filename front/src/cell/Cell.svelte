<script>
    import { LAND, GOLD } from './cellTypes';
    import { WARRIORS } from '../editModes';

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
    $: cellValue = () => {
        if (value.warrior) {
            return value.warrior.playerDisplayName
        }
        if (value.type === GOLD) {
            return value.pile
        }
        return ''
    }
</script>

<button
    class={cellClasses}
    {disabled}
    on:click
>{cellValue()}</button>

<style>
    .cell {
        width: 100%;
        height: 100%;
        border: 1px solid black;
        margin: 0;
        padding: 0;
    }
</style>
