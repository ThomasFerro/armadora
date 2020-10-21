<script>
  import { createEventDispatcher, onMount } from 'svelte';
import { i18n } from '../i18n/index.js';
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

  let partyNameInput = ''
</script>

<button on:click={createParty}>{$i18n('partySelection.create')}</button>
{#if creationError}
<p class="message error-message">{$i18n('partySelection.creationError')}</p>
{/if}
<details open class="parties-listing">
  <summary>{$i18n('partySelection.join')}</summary>
  
  <form on:submit|preventDefault={() => joinParty(partyNameInput)}>
    <label>
      {$i18n('partySelection.joinByName')}
      <input type="text" bind:value={partyNameInput}>
    </label>
    <input type="submit" value={$i18n('partySelection.joinParty')}>
  </form>


  {#if partiesLoadingState === LOADING}
  <p class="message info-message">{$i18n('partySelection.loadingParties')}</p>
  {:else if partiesLoadingState === LOADED}
  <ul class="parties">
    <button class="reload" on:click={loadParties}>{$i18n('partySelection.reloadParties')}</button>
    {#each parties as party}
      <li>
        <button on:click={() => joinParty(party)}>{$i18n('partySelection.joinParty')} {party}</button>
      </li>
    {/each}
  </ul>
  {:else if partiesLoadingState === ERROR}
  <p class="message error-message">{$i18n('partySelection.loadingPartiesError')}</p>
  <button class="reload" on:click={loadParties}>{$i18n('partySelection.reloadParties')}</button>
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