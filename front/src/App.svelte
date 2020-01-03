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

	let nextPlayerMask = false
	let turnInformation = undefined
	const turnPlayed = (information) => {
		const currentPlayer = game.currentPlayer
		game = playTurn(game, information)
		if (currentPlayer != game.currentPlayer) {
			nextPlayerMask = true
		}
	}

	$: nextPlayer = game && game.players[game.currentPlayer]

	const nextPlayerIsReady = () => {
		nextPlayerMask = false
	}
</script>

<main>
	<h1>Armadöra</h1>
	{#if game}
		{#if nextPlayerMask}
		<button on:click={nextPlayerIsReady}>Next player: {nextPlayer.race}</button>
		{/if}
		<Game
			{...game}
			hidden={nextPlayerMask}
			on:play-turn={(event) => turnPlayed(event.detail)}
			on:new-game={newGame}
		></Game>
	{:else}
		<!-- TODO: New game form -->
		<button on:click={newGame}>Start a game</button>
	{/if}

	<!-- TODO: Licence -->
</main>

<style>
</style>