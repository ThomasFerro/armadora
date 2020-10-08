<script>
    import { createEventDispatcher } from 'svelte'
    import Grid from './Grid.svelte'
    import PalisadeSelection from './PalisadeSelection.svelte'
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

    let warriorToPut
    const cellSelected = (details) => {
        warriorToPut = details
    }

    const putWarrior = (detail) => {
        if (!active) {
            return
        }
        dispatch('put-warrior', {
            ...detail.warrior,
            strength: selectedWarrior,
        })
        clearPalisades()
        clearWarriorToPut()
    }

    const clearWarriorToPut = () => {
        warriorToPut = undefined
    }

    $: connectedPlayerWarriors = connectedPlayer && connectedPlayer.warriors

    const warriorSelected = ({ strength }) => {
        selectedWarrior = strength
    }

    $: palisades = value && value.palisades || []
    $: palisadesLeft = value && value.palisades_left || 0
    let palisadeSelection = []

    const borderSelected = (newPalisade) => {
        if (palisadeSelection.length < 2) {
            palisadeSelection = [
                ...palisadeSelection,
                newPalisade,
            ]
        }
    }

    const clearPalisades = () => {
        palisadeSelection = []
    }
    
    const putPalisades = (palisades) => {
        dispatch('put-palisades', {
            palisades,
        })
        clearPalisades()
        clearWarriorToPut()
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
        selectedPalisades={palisadeSelection}
        on:cell-selected={(e) => cellSelected(e.detail)}
        on:border-selected={(e) => borderSelected(e.detail)}
    ></Grid>
    {#if active}
    <section class="player-actions">
        <PassTurn on:pass-turn={passTurn}></PassTurn>
        <WarriorSelection
            warriors={connectedPlayerWarriors}
            selectedWarrior={selectedWarrior}
            warriorToPut={warriorToPut}
            on:warrior-selected={(e) => warriorSelected(e.detail)}
            on:put-warrior={(e) => putWarrior(e.detail)}
            on:cancel-warrior-to-put={() => clearWarriorToPut()}
        ></WarriorSelection>
        <PalisadeSelection
            {palisadesLeft}
            {palisadeSelection}
            on:put-palisades={(e) => putPalisades(e.detail)}
            on:clear-palisades-selection={() => clearPalisades()}
        ></PalisadeSelection>
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
    flex-wrap: wrap;
}

.grid {
    grid-area: grid;
}

.current-player {
    display: flex;
    justify-content: center;
}
</style>
