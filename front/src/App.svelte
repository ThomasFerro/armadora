<script>
	import { onMount } from 'svelte'
	import Party from './Party.svelte'
	let parties = []
	let currentParty

	const createParty = () => {
		// TODO: Manage error
		fetch("/games", {
			method: "POST"
		}).then((response) => {
			return response.json()
		}).then(({id}) => {
			joinParty(id)
		})
	}

	const loadParties = () => {
		// TODO: Manage error + reaload
		fetch("/parties").then((response) => {
			return response.json()
		}).then((data) => {
			parties = data
		})
	}

	const joinParty = (party) => {
		currentParty = party
	}

	onMount(loadParties)
</script>

<main>
	<h1>Armadöra</h1>
	{#if !currentParty}
	<button on:click={createParty}>Create party</button>
	<ul>
		{#each parties as party}
		<li><button on:click={() => joinParty(party)}>Join {party}</button></li>
		{/each}
	</ul>
	{:else}
	<Party id={currentParty}></Party>
	{/if}
</main>
