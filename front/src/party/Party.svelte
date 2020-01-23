<script>
    import JoinAGame from './JoinAGame.svelte';
    import Board from './board/Board.svelte';
    export let id = undefined;

    let partyWs
    let game
    // TODO: Manage real authentication
    let connected = false
    
    $: {
		partyWs = new WebSocket(`ws://${window.location.host}/parties/${id}`);
		partyWs.addEventListener('message', ({data}) => {
            game = JSON.parse(data)
		})
    }

    $: availableCharacters = game && game.available_characters
    $: players = game && game.players
    $: waitingForPlayers = game && game.state && game.state === 'WaitingForPlayers'

    const connectToTheGame = ({username, character}) => {
        partyWs.send(JSON.stringify({
            "command_type": "JoinGame",
            "payload": {
                "Nickname": username,
                "Character": character,
            },
        }))
        connected = true
    }

    const startTheGame = () => {
        partyWs.send(JSON.stringify({
            "command_type": "StartTheGame",
        }))
    }

    $: board = game && game.board
</script>

<h2>Party: {id}</h2>
{#if waitingForPlayers}
    {#if !connected}
    <JoinAGame
        availableCharacters={availableCharacters}
        on:connect={(e) => connectToTheGame(e.detail)}
    ></JoinAGame>
    {:else}
    <button on:click={startTheGame}>Start the game</button>
    {/if}
    <ul class="players">
        {#each players as player}
        <li>{player.nickname} playing as {player.character}.</li>
        {/each}
    </ul>
{:else}
<Board value={board}></Board>
{/if}
