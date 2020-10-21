<script>
	import { onMount } from 'svelte'
	import Licences from './Licences.svelte'
	import NicknameSelection from './NicknameSelection.svelte'
	import Party from './party/Party.svelte'
	import PartySelection from './party/PartySelection.svelte'

	import { disconnect, getConnectedPlayerInformation, setNickname } from './authentication'
	import { getPartyNameFromUrl } from './route'

	import { i18n } from './i18n'
	import LocaleSelection from './i18n/LocaleSelection.svelte'

	let currentParty
	let nickname

	const joinParty = (party) => {
		currentParty = party
	}

	const leaveParty = () => {
		currentParty = null
	}

	const nicknameSelected = (newNickname) => {
		setNickname(newNickname)
		nickname = newNickname
	}

	const changeNickname = () => {
		nickname = ''
		disconnect()
	}

	onMount(() => {
		currentParty = getPartyNameFromUrl()
	})
</script>

<main>
	<h1>Armadöra</h1>
	{#if !nickname}
	<LocaleSelection></LocaleSelection>
	<NicknameSelection
		{getConnectedPlayerInformation}
		on:nickname-selected={(e) => nicknameSelected(e.detail)}
	></NicknameSelection>
	{:else if !currentParty}
		<LocaleSelection></LocaleSelection>
		{$i18n('home.connectedAs')} {nickname} <button on:click={changeNickname}>{$i18n('home.changeNickname')}</button>
		<PartySelection
			on:joinParty={(e) => joinParty(e.detail)}
		></PartySelection>
		<Licences></Licences>
	{:else}
	<Party
		id={currentParty}
		on:leave-party={leaveParty}
		{nickname}
	></Party>
	{/if}
</main>

<style>
main {
	width: 100%;
	height: 100%;
	display: flex;
	flex-flow: column nowrap;
	align-items: center;
}
</style>
