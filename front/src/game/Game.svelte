<script>
    import { createEventDispatcher } from 'svelte';
    import Grid from '../Grid.svelte';
    import Players from '../Players.svelte';
    import ActionChoice from '../ActionChoice.svelte';
    import { WARRIORS, PALISADES } from '../editModes';
    import { PUT_WARRIOR, PUT_PALISADES } from '../actionTypes';

    export let hidden = false;
    export let players;
    export let palisadesCount;
    export let grid;
    export let currentPlayer;
    export let results = undefined;

    const dispatch = createEventDispatcher();

    let editMode = PALISADES;

    const palisadesEditMode = () => {
        selectedWarrior = null
        editMode = PALISADES
    }

    const warriorsEditMode = () => {
        editMode = WARRIORS
    }

    let selectedWarrior;
    const selectWarrior = (event) => {
        selectedWarrior = event.detail.warriorIndex
        warriorsEditMode()
    }

    $: currentPlayerWarriors = players[currentPlayer] && players[currentPlayer].warriors

    const cellClicked = (information) => {
        if (editMode === WARRIORS) {
            dispatch('play-turn', {
                type: PUT_WARRIOR,
                selectedWarrior,
                ...information
            })
        }
    }

    // TODO : Extract in a palisade module
    let selectedPalisades = []

    $: ongoingPalisadeSelection = selectedPalisades.length > 0

    const palisadeClicked = (information) => {
        if (selectedPalisades.length < 2) {
            selectedPalisades = [
                ...selectedPalisades,
                information
            ]
        }
    }

    const clearPalisadeSelection = () => {
        selectedPalisades = []
    }

    const validatePalisades = () => {
        if (editMode === PALISADES && selectedPalisades.length > 0) {
            dispatch('play-turn', {
                type: PUT_PALISADES,
                palisades: selectedPalisades
            })
            selectedPalisades = []
        }
    }

    $: winner = results && players[results.winner] && players[results.winner].race

    const newGame = () => {
        dispatch('new-game')
    }
</script>

<article class="game" class:hidden>
    {#if !results}
    <Players currentPlayer={currentPlayer} players={players}></Players>
    <p>Palisades: {palisadesCount}</p>
    <Grid
        {grid}
        mode={editMode}
        {selectedPalisades}
        on:cell-clicked={(event) => cellClicked(event.detail)}
        on:palisade-clicked={(event) => palisadeClicked(event.detail)}
    ></Grid>
    <ActionChoice
        selectedWarrior={selectedWarrior}
        currentPlayerWarriors={currentPlayerWarriors}
        hasPalisadesLeft={palisadesCount > 0}
        ongoingPalisadeSelection={ongoingPalisadeSelection}
        on:warrior-selected={selectWarrior}
        on:palisades-selected={palisadesEditMode}
        on:cancel-palisade-selection={clearPalisadeSelection}
        on:validate-palisade-selection={validatePalisades}
    ></ActionChoice>
    {:else}
    <span>{winner} wins !</span>
    <button on:click={newGame}>New game</button>
    {/if}
</article>

<style>
.game.hidden {
    display: none;
}
</style>
