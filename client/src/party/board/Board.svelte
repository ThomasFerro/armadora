<script>
    import { createEventDispatcher } from 'svelte'
    import Grid from './Grid.svelte'
    import PassTurn from './PassTurn.svelte'
    import WarriorSelection from './WarriorSelection.svelte'

    // FIXME remove the const
    // export let active = false
    const active = true
    export let value = {}
    export let connectedPlayer = {}
    export let currentPlayer = null

    $: currentPlayerDisplayedInformation = currentPlayer ?
        `${currentPlayer.nickname}'s (${currentPlayer.character}) turns`
        : "No current player..."

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

    const passTurn = () => {
        dispatch('pass-turn')
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
    <section class="player-actions">
        <PassTurn on:pass-turn={passTurn}></PassTurn>
        <WarriorSelection
            warriors={connectedPlayerWarriors}
            selectedWarrior={selectedWarrior}
            on:warrior-selected={(e) => warriorSelected(e.detail)}
        ></WarriorSelection>
    </section>
    {:else}
    <p class="current-player">{currentPlayerDisplayedInformation}</p>
    {/if}
</article>

<style>
.board {
    width: 100%;
    flex: 1;
    display: grid;
    /* TODO: Not in px ? */
    grid-template:
        "grid" 1fr
        "player-actions" 50px;
}

.player-actions {
    grid-area: player-actions;
    display: flex;
    align-items: center;
    justify-content: center;
}

.grid {
    grid-area: grid;
}

.current-player {
    display: flex;
    justify-content: center;
}
</style>
