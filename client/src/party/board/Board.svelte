<script>
    import { createEventDispatcher } from 'svelte'
    import Grid from './Grid.svelte'
    import PassTurn from './PassTurn.svelte'
    import WarriorSelection from './WarriorSelection.svelte'

    export let active = false
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
        <WarriorSelection
            warriors={connectedPlayerWarriors}
            on:warrior-selected={(e) => warriorSelected(e.detail)}
        ></WarriorSelection>
        <PassTurn on:pass-turn={passTurn}></PassTurn>
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
}

.grid {
    grid-area: grid;
}

.current-player {
    display: flex;
    justify-content: center;
}
</style>
