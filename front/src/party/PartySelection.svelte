<script>
  import { createEventDispatcher, onMount } from "svelte";
  import { LOADING, LOADED, ERROR } from '../loading.js';
  const dispatch = createEventDispatcher();
  export let parties = [];

  let creationError
  let partiesLoadingState

  const createParty = () => {
    creationError = null
    fetch("/games", {
      method: "POST"
    })
      .then(response => {
        return response.json();
      })
      .then(({ id }) => {
        joinParty(id);
      })
      .catch((e) => {
        creationError = e
      });
  };

  const loadParties = () => {
    partiesLoadingState = LOADING
    fetch("/parties")
      .then(response => {
        return response.json();
      })
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
<span class="error-message">An error has occurred while creating the party</span>
{/if}
{#if partiesLoadingState === LOADING}
<span class="info-message">Loading parties</span>
{:else if partiesLoadingState === LOADED}
<ul>
  {#each parties as party}
    <li>
      <button on:click={() => joinParty(party)}>Join {party}</button>
    </li>
  {/each}
</ul>
{:else if partiesLoadingState === ERROR}
<span class="error-message">An error has occurred while loading the parties</span>
<button class="reload" on:click={loadParties}>Reload</button>
{/if}
