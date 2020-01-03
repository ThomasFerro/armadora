<script>
	import { onMount } from 'svelte';
	import Game from './game/Game.svelte';
	import { createGame } from './game/gameFactory';
	import { playTurn } from './game/gameEngine';

	let game;

	const newGame = () => {
		game = createGame();
	}

	onMount(newGame);
</script>

<main>
	<h1>Armadöra</h1>
	{#if game}
		<Game
			on:play-turn={(event) => game = playTurn(game, event.detail)}
			{...game}
		></Game>
	{:else}
		<!-- TODO: New game form -->
		<button on:click={newGame}>Start a game</button>
	{/if}

	<!-- TODO: Licence -->
</main>

<style>
</style>