<script>
	import { onMount } from 'svelte'
	import Party from './party/Party.svelte'
	import PartySelection from './party/PartySelection.svelte'
	let parties = []
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
	{:else}
	<Party id={currentParty}></Party>
	{/if}
</main>
