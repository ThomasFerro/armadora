<script>
    import { createEventDispatcher, onMount } from 'svelte';
    import { i18n } from './i18n';
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
    <p>{$i18n('nickname.authenticationWarning')}</p>

    <p>{$i18n('nickname.linkedToNickname')}</p>

    <p><strong>{$i18n('nickname.nicknameUseCase')}</strong></p>

    <label>
        {$i18n('nickname.nickname')}:
        <input type="text" bind:value={nicknameInput}>
    </label>
    <input type="submit" value={$i18n('nickname.validate')}>
</form>
