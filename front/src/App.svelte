<script>
	import NewGame from './NewGame.svelte';
	import Game from './game/Game.svelte';
	import { playTurn } from './game/gameEngine';

	let game;

	const newGame = () => {
		game = null
	}

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
		<NewGame
			on:new-game={({ detail }) => game = detail}
		></NewGame>
	{/if}

	<!-- TODO: Design -->
	<!-- TODO: Tuto / rules -->
	<!-- TODO: Licence -->
</main>

<style>
</style>