<script>
    import { createEventDispatcher, onMount } from 'svelte';
    export let getConnectedPlayerInformation

    const dispatch = createEventDispatcher();
    
    let nicknameInput = ''

    const nicknameSelected = (nickname) => {
        dispatch('nickname-selected', nickname)
    }

    const loadNicknameFromLocalStorage = () => {
        if (!getConnectedPlayerInformation) {
            return
        }
        const { nickname } = getConnectedPlayerInformation()
        if (nickname) {
            nicknameSelected(nickname)
        }
    }
    onMount(loadNicknameFromLocalStorage)
</script>

<form on:submit|preventDefault={() => nicknameSelected(nicknameInput)}>
    <p>🚨🚨</p>
    <p>The game is still in active development and does not yet provide a full authentication flow.</p>

    <p>For now, your parties are linked to the nickname you choose.</p>

    <p><strong>Two players with the same nickname will act as if they play on the same account.</strong></p>

    <label>
        Nickname:
        <input type="text" bind:value={nicknameInput}>
    </label>
    <input type="submit" value="Validate nickname">
</form>
