<script>
  import { createEventDispatcher, onMount } from 'svelte';
  import { LOADING, LOADED, ERROR } from '../loading.js';
  import { createNewParty, getParties } from './party.service.js';
  const dispatch = createEventDispatcher();
  export let parties = [];

  let creationError
  let partiesLoadingState

  const createParty = () => {
    creationError = null
    createNewParty()
      .then(({ id }) => {
        joinParty(id);
      })
      .catch((e) => {
        console.error(e)
        creationError = e
      });
  };

  const loadParties = () => {
    partiesLoadingState = LOADING
    getParties()
      .then(data => {
        partiesLoadingState = LOADED
        parties = data;
      })
      .catch(() => {
        partiesLoadingState = ERROR
      });
  };

  const joinParty = party => {
    dispatch('joinParty', party)
  };

  onMount(loadParties);
</script>

<button on:click={createParty}>Create party</button>
{#if creationError}
<p class="message error-message">An error has occurred while creating the party</p>
{/if}
<details open class="parties-listing">
  <summary>Join a party</summary>
  {#if partiesLoadingState === LOADING}
  <p class="message info-message">Loading parties</p>
  {:else if partiesLoadingState === LOADED}
  <ul class="parties">
    <button class="reload" on:click={loadParties}>Reload</button>
    {#each parties as party}
      <li>
        <button on:click={() => joinParty(party)}>Join {party}</button>
      </li>
    {/each}
  </ul>
  {:else if partiesLoadingState === ERROR}
  <p class="message error-message">An error has occurred while loading the parties</p>
  <button class="reload" on:click={loadParties}>Reload</button>
  {/if}
</details>

<style>
.parties-listing {
  flex: 1;
  overflow: auto;
}

.parties {
  display: flex;
  flex-flow: column nowrap;
}
</style>