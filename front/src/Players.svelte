<script>
    export let players;
    export let currentPlayer;

    $: currentPlayerDisplay = currentPlayer < players.length ? players[currentPlayer].race : 'N/A'
</script>

Current player: {currentPlayerDisplay}
<ol>
    {#each players as player, index}
        {#if index === currentPlayer}
            <li>{player.race}: {player.warriors}</li>
        {:else}
            <li>{player.race}: {player.warriors.length}</li>
        {/if}
    {/each}
</ol>