<script>
	import { onMount } from 'svelte';
	import Game from './Game.svelte';
	import { createGame } from './gameFactory';

	let game;

	const newGame = () => {
		game = createGame();
	}

	onMount(newGame);
</script>

<!-- TODO:
- Gérer un tour de jeu:
	- Ajouter une ou deux palisade(s)
	- Ajouter un combattant
- Gérer la fin de partie
 -->

<main>
	<h1>Armadöra</h1>
	{#if game}
		<Game {...game}></Game>
	{:else}
		<!-- TODO: New game form -->
		<button on:click={newGame}>Start a game</button>
	{/if}
</main>

<style>
</style>