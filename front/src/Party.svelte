<script>
    import JoinAGame from './JoinAGame.svelte';
    export let id = undefined;

    let partyWs
    let game
    
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
    }
</script>

<h2>Party: {id}</h2>
{#if waitingForPlayers}
<JoinAGame
    availableCharacters={availableCharacters}
    players={players}
    on:connect={(e) => connectToTheGame(e.detail)}
></JoinAGame>
{/if}
