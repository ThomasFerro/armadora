<script>
    import { createEventDispatcher } from 'svelte';
    import { WARRIORS, PALISADES } from './editModes';

    export let currentPlayerWarriors;
    export let selectedWarrior;

    const dispatch = createEventDispatcher();

    const selectWarrior = (warriorIndex) => {
        dispatch('warrior-selected', { warriorIndex })
    }

    const selectPalisades = () => {
        dispatch('palisades-selected')
    }
</script>

<section class="action">
    <!-- TODO: disabled if there no palisade left -->
    <label>
        <input type=radio bind:group={selectedWarrior} value={undefined} on:input={() => selectPalisades()}>
        Palisades
    </label>
    {#each currentPlayerWarriors as currentPlayerWarrior, warriorIndex}
        <label>
            <input type=radio bind:group={selectedWarrior} value={warriorIndex} on:input={() => selectWarrior(warriorIndex)}>
            {currentPlayerWarrior}
        </label>
    {/each}
</section>