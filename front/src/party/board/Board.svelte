<script>
    import { createEventDispatcher } from 'svelte'
    import Grid from './Grid.svelte';
    import WarriorSelection from './WarriorSelection.svelte';

    export let active = false
    export let value = {}
    export let connectedPlayer = {}

    let selectedWarrior

    const dispatch = createEventDispatcher()

    $: cells = value && value.cells || []
    $: palisades = value && value.palisades || []

    const cellSelected = (details) => {
        if (!active) {
            return
        }
        dispatch('put-warrior', {
            ...details,
            strength: selectedWarrior,
        })
    }

    $: connectedPlayerWarriors = connectedPlayer && connectedPlayer.warriors

    const warriorSelected = ({ strength }) => {
        selectedWarrior = strength
    }

    const borderSelected = (palisades) => {
        dispatch('put-palisades', {
            palisades,
        })
    }
</script>

<article class="board">
    <Grid
        {active}
        {cells}
        {palisades}
        on:cell-selected={(e) => cellSelected(e.detail)}
        on:border-selected={(e) => borderSelected(e.detail)}
    ></Grid>
    {#if active}
    <WarriorSelection
        warriors={connectedPlayerWarriors}
        on:warrior-selected={(e) => warriorSelected(e.detail)}
    ></WarriorSelection>
    {/if}
</article>
