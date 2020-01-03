<script>
    import { createEventDispatcher } from 'svelte';
    import { WARRIORS, PALISADES } from '../editModes';

    export let currentPlayerWarriors;
    export let selectedWarrior;
    export let hasPalisadesLeft = false;
    export let ongoingPalisadeSelection = false;

    const dispatch = createEventDispatcher();

    const selectWarrior = (warriorIndex) => {
        dispatch('warrior-selected', { warriorIndex })
    }

    const selectPalisades = () => {
        dispatch('palisades-selected')
    }

    const validatePalisadeSelection = () => {
        dispatch('validate-palisade-selection')
    }

    const cancelPalisadeSelection = () => {
        dispatch('cancel-palisade-selection')
    }
</script>

<section class="action">
    {#if ongoingPalisadeSelection}
        Ongoing palisade selection
        <button on:click={validatePalisadeSelection}>Validate</button>
        <button on:click={cancelPalisadeSelection}>Cancel</button>
    {:else}
        {#if hasPalisadesLeft}
        <label>
            <input
                type=radio bind:group={selectedWarrior}
                value={-1}
                on:input={() => selectPalisades()}
            >
            Palisades
        </label>
        {/if}
        {#each currentPlayerWarriors as currentPlayerWarrior, warriorIndex}
            <label>
                <input type=radio bind:group={selectedWarrior} value={warriorIndex} on:input={() => selectWarrior(warriorIndex)}>
                {currentPlayerWarrior}
            </label>
        {/each}
    {/if}
</section>