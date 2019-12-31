<script>
    import { createEventDispatcher } from 'svelte';
    import Grid from './Grid.svelte';
    import Players from './Players.svelte';
    import ActionChoice from './ActionChoice.svelte';
    import { WARRIORS, PALISADES } from './editModes';
    import { PUT_WARRIOR, PUT_PALISADES } from './actionTypes';

    export let players;
    export let palisadesCount;
    export let grid;
    export let currentPlayer;

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

    // TODO: Manage two palisades
    const palisadeClicked = (information) => {
        if (editMode === PALISADES) {
            dispatch('play-turn', {
                type: PUT_PALISADES,
                palisades: [
                    {
                        ...information
                    }
                ]
            })
        }
    }
</script>

<article class="game">
    <Players currentPlayer={currentPlayer} players={players}></Players>
    <p>Palisades: {palisadesCount}</p>
    <Grid
        {grid}
        mode={editMode}
        on:cell-clicked={(event) => cellClicked(event.detail)}
        on:palisade-clicked={(event) => palisadeClicked(event.detail)}
    ></Grid>
    <ActionChoice
        selectedWarrior={selectedWarrior}
        currentPlayerWarriors={currentPlayerWarriors}
        on:warrior-selected={selectWarrior}
        on:palisades-selected={palisadesEditMode}
    ></ActionChoice>
</article>