<script>
    import Grid from './Grid.svelte';
    import Players from './Players.svelte';
    import ActionChoice from './ActionChoice.svelte';
    import { WARRIORS, PALISADES } from './editMode';

    export let players;
    export let palisadesCount;
    export let grid;
    export let currentPlayer;

    let editMode;

    const palisadesEditMode = async () => {
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
    $: currentPlayerWarriors = players[currentPlayer].warriors

    const cellClicked = ({ x, y }) => {
        if (editMode === WARRIORS) {
            // TODO
            console.log(`Put warrior ${selectedWarrior} (${currentPlayerWarriors[selectedWarrior]}) of player ${currentPlayer} in the cell ${x},${y}`);

            console.log('Next turn')
        }
    }
</script>

<article class="game">
    <h2>Summary</h2>
    <ul>
        <li>
            <Players currentPlayer={currentPlayer} players={players}></Players>
        </li>
        <li>Palisades: {palisadesCount}</li>
        <li>
            Grid:
            <Grid
                value={grid}
                mode={editMode}
                on:cell-clicked={(event) => cellClicked(event.detail)}
            ></Grid>
        </li>
    </ul>
    <ActionChoice
        selectedWarrior={selectedWarrior}
        on:warrior-selected={selectWarrior}
        currentPlayerWarriors={currentPlayerWarriors}
        on:palisades-edit-mode={palisadesEditMode}
    ></ActionChoice>
</article>