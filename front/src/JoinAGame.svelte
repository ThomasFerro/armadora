<script>
    import { createEventDispatcher } from 'svelte'
    const dispatch = createEventDispatcher()
    export let availableCharacters = []

    let username;
    let character;
    $: canConnect = username && character
    
    const connectToTheGame = () => {
        if (!canConnect) {
            return
        }
        dispatch('connect', {
            username,
            character,
        })
    }
</script>

<form class="join-a-game" on:submit|preventDefault={connectToTheGame}>
    <label>
        Username:
        <input type="text" bind:value={username}>
    </label>
    <label>
        Character:
        <select bind:value={character}>
            <option disabled>Select a character</option>
            {#each availableCharacters as availableCharacter}
            <option value={availableCharacter}>{availableCharacter}</option>
            {/each}
        </select>
    </label>
    <input type="submit" value="Connect to the game" disabled={!canConnect}>
</form>