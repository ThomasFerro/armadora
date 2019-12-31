<script>
	import { onMount } from 'svelte';
	import Game from './Game.svelte';
	import { createGame } from './gameFactory';
	import { playTurn } from './gameEngine';

	let game;

	const newGame = () => {
		game = createGame();
	}

	onMount(newGame);
</script>

<!-- TODO:
- Gérer un tour de jeu:
	- Empêcher l'ajout d'une palisade si ça ferme une zone à moins de 4 cellules
- Gérer la fin de partie
 -->

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