<script>
    import { createEventDispatcher } from 'svelte';

    export let currentPlayerWarriors;
    export let selectedWarrior;

    const dispatch = createEventDispatcher();

    const palisadesEditMode = () => {
        dispatch('palisades-edit-mode')
    }

    const selectWarrior = (warriorIndex) => {
        dispatch('warrior-selected', { warriorIndex })
    }
</script>

<section class="action">
    <!-- TODO: Only show if there are palisades left -->
    <button on:click={palisadesEditMode}>Put palisades</button>
    {#each currentPlayerWarriors as currentPlayerWarrior, warriorIndex}
        <label>
            <input type=radio bind:group={selectedWarrior} value={warriorIndex} on:input={() => selectWarrior(warriorIndex)}>
            {currentPlayerWarrior}
        </label>
    {/each}
</section>