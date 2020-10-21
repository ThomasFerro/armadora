<script>
    import { createEventDispatcher, onDestroy } from 'svelte'
    import JoinAGame from './JoinAGame.svelte';
    import Board from './board/Board.svelte';
    import Scores from './score/Scores.svelte';
    import Player from './Player.svelte';
    import { LOADING, LOADED, ERROR } from '../loading';
    import { connectToGame, gameInformation, startGame, putWarrior, putPalisades, passTurn } from './party.service.js';

    export let id
    export let nickname

    const dispatch = createEventDispatcher()
    const leaveParty = () => {
        dispatch('leave-party')
    }

    let game
    let gameUpdateTimeout

    let partyError = ''

    const loadGameInformation = () => {
        return gameInformation(id)
            .then((updatedGame) => {
                partyError = ''
                game = updatedGame
            })
            .catch(() => {
                partyError = 'Cannot load the game\'s status'
            })
    }

    const newGameUpdateTimeout = () => {
        gameUpdateTimeout = setTimeout(() => {
            if (id) {
                loadGameInformation().finally(() => {
                    newGameUpdateTimeout()
                })
            }
        }, 1000)
    }

    $: availableCharacters = game && game.available_characters
    $: players = game && game.players
    $: waitingForPlayers = game && game.state && game.state === 'WaitingForPlayers'

    const connectToTheGame = ({ character }) => {
        partyError = ''
        connectToGame(id)({
            character,
            username: nickname
        })
            .catch(() => {
                partyError = 'Unable to connect to the game'
            })
    }

    let startTheGameStatus = ''
    $: gameIsStarting = startTheGameStatus === LOADING
    $: startTheGameLabel = gameIsStarting ? 'Starting the new game...' : 'Start the game'
    const startTheGame = () => {
        partyError = ''
        startTheGameStatus = LOADING
        startGame(id)
            .then(() => {
                startTheGameStatus = LOADED
            })
            .catch(() => {
                partyError = 'Unable to start the game'
                startTheGameStatus = ERROR
            })
    }

    $: board = game && game.board
    // TODO: Pay tech debt after doing real authent
    const sameNicknameAsConectedPlayer = (player) => player.nickname === nickname
    $: connectedPlayer = game && game.players.find(sameNicknameAsConectedPlayer)
    $: indexOfConnectecPlayer = game && game.players.indexOf(connectedPlayer)
    $: turnOfConnectedPlayer = game && indexOfConnectecPlayer === game.current_player

    $: currentPlayer = game && game.players[game.current_player]

    let actionLoadingState
    $: actionPending = actionLoadingState === LOADING
    $: actionError = actionLoadingState === ERROR
    $: isBoardActive = turnOfConnectedPlayer && actionLoadingState !== LOADING

    const putWarriorAction = (warriorData) => {
        actionLoadingState = LOADING
        putWarrior(id)(indexOfConnectecPlayer)(warriorData)
            .then(loadGameInformation)
            .then(() => {
                actionLoadingState = LOADED
            })
            .catch(() => {
                actionLoadingState = ERROR
            })
    }

    const putPalisadesAction = (palisadesData) => {
        actionLoadingState = LOADING
        putPalisades(id)(indexOfConnectecPlayer)(palisadesData)
            .then(loadGameInformation)
            .then(() => {
                actionLoadingState = LOADED
            })
            .catch(() => {
                actionLoadingState = ERROR
            })
    }

    const passTurnAction = () => {
        actionLoadingState = LOADING
        passTurn(id)(indexOfConnectecPlayer)
            .then(loadGameInformation)
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

    onDestroy(() => {
        clearTimeout(gameUpdateTimeout)
    })
</script>

<h2 class="party-title">Party {id} <button on:click={leaveParty}>⍇</button></h2>
{#if partyError}
<p class="message error-message">{partyError}</p>
{/if}
{#if actionPending}
<p class="message info-message">Sending your action...</p>
{:else if actionError}
<p class="message error-message">An error has occurred while sending your action</p>
{/if}
{#if waitingForPlayers}
    {#if !connectedPlayer}
    <JoinAGame
        availableCharacters={availableCharacters}
        on:connect={(e) => connectToTheGame(e.detail)}
    ></JoinAGame>
    {:else}
    <button on:click={startTheGame} disabled={gameIsStarting}>{startTheGameLabel}</button>
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
    active={isBoardActive}
    connectedPlayer={connectedPlayer}
    {currentPlayer}
    on:put-warrior={(e) => putWarriorAction(e.detail)}
    on:put-palisades={(e) => putPalisadesAction(e.detail)}
    on:pass-turn={passTurnAction}
></Board>
{/if}

<style>
.party-title {
    margin-block-start: 0;
    display: flex;
    align-items: center;
}

.players {
    margin-block-start: 1em;
}

.player {
    display: flex;
    align-items: center;
}
</style>
