<script>
	import Grid from './Grid.svelte';

	export let game;
</script>

<main>
	<h1>Armadöra</h1>
	<article>
		<h2>Summary</h2>
		<ul>
			<li>
				Players:
				<ol>
					{#each game.players as player}
						<li>{player.race}: {player.warriors}</li>
					{/each}
				</ol>
			</li>
			<li>Palisages: {game.palisadesCount}</li>
			<li>
				Grid:
				<Grid grid={game.grid}></Grid>
			</li>
		</ul>
	</article>
</main>

<style>
</style>