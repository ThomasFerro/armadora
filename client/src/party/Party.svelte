<script>
    import JoinAGame from './JoinAGame.svelte';
    import Board from './board/Board.svelte';
    import Scores from './score/Scores.svelte';
    import Player from './Player.svelte';
    import { LOADING, LOADED, ERROR } from '../loading';
    import { connectToGame, gameInformation, startGame, putWarrior, putPalisades, passTurn } from './party.service.js';

    export let id = undefined;

    let game
    let gameUpdateTimeout
    // TODO: Manage real authentication
    let connected = false
    let nickname = ''

    const newGameUpdateTimeout = () => {
        gameUpdateTimeout = setTimeout(() => {
            if (id) {
                // TODO: Error + loading management
                gameInformation(id)
                    .then((updatedGame) => {
                        game = updatedGame
                    })
                    .finally(() => {
                        newGameUpdateTimeout()
                    })
            }
        }, 1000)
    }

    $: availableCharacters = game && game.available_characters
    $: players = game && game.players
    $: waitingForPlayers = game && game.state && game.state === 'WaitingForPlayers'

    const connectToTheGame = (userData) => {
        nickname = userData.username
        // TODO: Error management
        connectToGame(id)(userData)
            .then(() => {
                connected = true
            })
    }

    // TODO: Error management
    const startTheGame = () => startGame(id)

    $: board = game && game.board
    // TODO: Pay tech debt after doing real authent
    const sameNicknameAsConectedPlayer = (player) => player.nickname === nickname
    $: connectedPlayer = game && game.players.find(sameNicknameAsConectedPlayer)
    $: turnOfConnectedPlayer = game && game.players.indexOf(connectedPlayer) === game.current_player

    $: currentPlayer = game && game.players[game.current_player]

    // TODO: Display loading + error
    let actionLoadingState

    const putWarriorAction = (warriorData) => {
        actionLoadingState = LOADING
        putWarrior(id)(warriorData)
            .then(() => {
                actionLoadingState = LOADED
            })
            .catch(() => {
                actionLoadingState = ERROR
            })
    }

    const putPalisadesAction = (palisadesData) => {
        actionLoadingState = LOADING
        putPalisades(id)(palisadesData)
            .then(() => {
                actionLoadingState = LOADED
            })
            .catch(() => {
                actionLoadingState = ERROR
            })
    }

    const passTurnAction = () => {
        actionLoadingState = LOADING
        passTurn(id)
            .then(() => {
                actionLoadingState = LOADED
            })
            .catch(() => {
                actionLoadingState = ERROR
            })
    }

    $: finished = game && game.state && game.state === 'Finished'
    $: {
        clearTimeout(gameUpdateTimeout)
        if (!finished) {
            newGameUpdateTimeout()
        }
    }
    $: scores = game && game.scores
</script>

<h2>Party {id}</h2>
<!-- TODO: Loading + error -->
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
        <li class="player">
            <Player {player}></Player>
        </li>
        {/each}
    </ul>
{:else if finished}
<Scores value={scores} players={players}></Scores>
{:else}
<Board
    value={board}
    active={turnOfConnectedPlayer}
    connectedPlayer={connectedPlayer}
    {currentPlayer}
    on:put-warrior={(e) => putWarriorAction(e.detail)}
    on:put-palisades={(e) => putPalisadesAction(e.detail)}
    on:pass-turn={passTurnAction}
></Board>
{/if}

<style>
h2 {
	margin-block-start: 0;
}

.players {
    margin-block-start: 1em;
}

.player {
    display: flex;
    align-items: center;
}
</style>
