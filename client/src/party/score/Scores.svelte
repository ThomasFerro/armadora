<script>
    import { i18n } from "../../i18n";

    export let value = {}
    export let players = []
</script>

<table>
    <thead>
        <tr>
            <th></th>
            <th>{$i18n('scores.username')}</th>
            <th>{$i18n('scores.character')}</th>
            <th>{$i18n('scores.totalGold')}</th>
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