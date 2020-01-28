<script>
    import JoinAGame from './JoinAGame.svelte';
    import Board from './board/Board.svelte';
    export let id = undefined;

    let partyWs
    let game
    // TODO: Manage real authentication
    let connected = false
    let nickname = ''
    
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
        nickname = username
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
    // TODO: Pay tech debt after doing real authent
    const sameNicknameAsConectedPlayer = (player) => player.nickname === nickname
    $: connectedPlayer = game && game.players.find(sameNicknameAsConectedPlayer)
    $: turnOfConnectedPlayer = game && game.players.indexOf(connectedPlayer) === game.current_player

    const putWarrior = ({x, y, strength}) => {
        // FIXME: Smothing smelly with the grid
        partyWs.send(JSON.stringify({
            "command_type": "PutWarrior",
            "payload": {
                "Warrior": strength,
                "X": x.toString(),
                "Y": y.toString(),
            },
        }))
    }

    const putPalisades = ({ palisades }) => {
        partyWs.send(JSON.stringify({
            "command_type": "PutPalisades",
            "payload": {
                "Palisades": JSON.stringify(palisades)
            },
        }))
    }
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
<Board
    value={board}
    active={turnOfConnectedPlayer}
    connectedPlayer={connectedPlayer}
    on:put-warrior={(e) => putWarrior(e.detail)}
    on:put-palisades={(e) => putPalisades(e.detail)}
></Board>
{/if}
