<script>
	import Licences from './Licences.svelte'
	import Party from './party/Party.svelte'
	import PartySelection from './party/PartySelection.svelte'
	let currentParty

	const joinParty = (party) => {
		currentParty = party
	}
</script>

<main>
	<h1>Armadöra</h1>
	{#if !currentParty}
	
	<PartySelection
		on:joinParty={(e) => joinParty(e.detail)}
	></PartySelection>
	<Licences></Licences>
	{:else}
	<Party id={currentParty}></Party>
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
