<script>
    export let players;
    export let currentPlayer;
</script>

Players: {currentPlayer}
<ol>
    {#each players as player, index}
        {#if index === currentPlayer}
            <li>{player.race}: {player.warriors}</li>
        {:else}
            <li>{player.race}: {player.warriors.length}</li>
        {/if}
    {/each}
</ol>