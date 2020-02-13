<script>
    export let value = {}
    export let players = []
</script>

<table>
    <thead>
        <tr>
            <th></th>
            <th>Username</th>
            <th>Character</th>
            <th>Total gold</th>
        </tr>
    </thead>
    <tbody>
        {#each Object.entries(value) as [position, information]}
        <tr>
            <td>{position}</td>
            <td>{players[information.player].nickname}</td>
            <td>{players[information.player].character}</td>
            <td>{information.gold}</td>
        </tr>
        {/each}
    </tbody>
</table>