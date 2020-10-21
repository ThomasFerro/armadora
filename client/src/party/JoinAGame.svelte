<script>
    import { createEventDispatcher } from 'svelte'
    import { i18n } from '../i18n';
    const dispatch = createEventDispatcher()
    export let availableCharacters = []

    let character;
    $: canConnect = !!character

    $: selectableCharacters = availableCharacters.sort()
    $: {
        if (selectableCharacters.indexOf(character) === -1) {
            character = selectableCharacters && selectableCharacters[0]
        }
    }
    
    const connectToTheGame = () => {
        if (!canConnect) {
            return
        }
        dispatch('connect', {
            character,
        })
    }
</script>

<form class="join-a-game" on:submit|preventDefault={connectToTheGame}>
    <label class="character-selection">
        {$i18n('joinAGame.selectYourCharacter')}
        <select bind:value={character}>
            <option disabled>{$i18n('joinAGame.selectACharacter')}</option>
            {#each selectableCharacters as selectableCharacter (selectableCharacter)}
            <option value={selectableCharacter}>{selectableCharacter}</option>
            {/each}
        </select>
    </label>
    <input type="submit" value={$i18n('joinAGame.connect')} disabled={!canConnect}>
</form>

<style>
.join-a-game {
    display: flex;
    flex-flow: column nowrap;
}

.character-selection {
    display: flex;
    flex-flow: row wrap;
    align-items: center;
}
</style>
