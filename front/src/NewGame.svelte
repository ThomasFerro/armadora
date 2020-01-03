<script>
    import { createEventDispatcher } from 'svelte';
	import { createGame } from './game/gameFactory';

    const dispatch = createEventDispatcher();

    const players = [
        {
            race: 'Orc',
            selected: true
        },
        {
            race: 'Goblin',
            selected: true
        },
        {
            race: 'Elf',
            selected: true
        },
        {
            race: 'Mage',
            selected: true
        },
    ]

    $: selectedPlayers = players.filter(player => player.selected).map(player => player.race)

    const newGame = () => {
        dispatch('new-game', createGame(selectedPlayers));
    }
</script>

Players: { selectedPlayers }
{#each players as player}
<label>
    Play as {player.race} ?
    <input type=checkbox bind:checked={player.selected}>
</label>
{/each}
<button on:click={newGame}>Start a game</button>
