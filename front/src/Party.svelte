<script>
    export let id = undefined;

    let partyWs
    let game
    
    $: {
		partyWs = new WebSocket(`ws://${window.location.host}/parties/${id}`);
		partyWs.addEventListener('message', ({data}) => {
            game = data
		})
    }
</script>

<h2>Party: {id}</h2>
<pre>{game}</pre>
