<script>
    import Grid from './Grid.svelte';
    import Players from './Players.svelte';

    export let players;
    export let palisadesCount;
    export let grid;
    export let currentPlayer;
</script>

<article class="game">
    <h2>Summary</h2>
    <ul>
        <li>
            <Players currentPlayer={currentPlayer} players={players}></Players>
        </li>
        <li>Palisades: {palisadesCount}</li>
        <li>
            Grid:
            <Grid value={grid}></Grid>
        </li>
    </ul>
</article>