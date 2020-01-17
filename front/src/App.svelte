<script>
	import { onMount } from 'svelte'

	let ws;
	let username;
	let character;
	let game;

	$: canConnect = username && character && game
	// TODO: Real connection
	$: connected = game && game.players && game.players.find(player => player.nickname === username)
	$: availableCharacters = game && game.available_characters || []

	const connectToTheGame = () => {
		if (!canConnect) {
			return
		}
		ws.send(JSON.stringify({
			command_type: 'JoinGame',
			payload: {
				Nickname: username,
				Character: character,
			}
		}))
	}

	const startGame = () => {
		ws.send(JSON.stringify({
			command_type: 'StartTheGame',
		}))
	}

	onMount(() => {
		ws = new WebSocket('ws://' + window.location.host + '/ws');
		ws.addEventListener('message', (e) => {
			game = JSON.parse(e.data);
		})
	})
</script>

<main>
	<h1>Armadöra</h1>
	{#if !connected}
	<form on:submit|preventDefault={connectToTheGame}>
		<label>
			Username:
			<input type="text" bind:value={username}>
		</label>
		<label>
			Character:
			<select bind:value={character}>
				<option disabled>Select a character</option>
				{#each availableCharacters as availableCharacter}
				<option value={availableCharacter}>{availableCharacter}</option>
				{/each}
			</select>
		</label>
		<input type="submit" value="Connect to the game" disabled={!canConnect}>
	</form>
	{:else}
	Connected !!!!
	<button on:click={startGame}>Start that game</button>
	{/if}
	<pre>{JSON.stringify(game)}</pre>
</main>
