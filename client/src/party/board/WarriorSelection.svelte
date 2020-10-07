<script>
    import { createEventDispatcher, onMount } from 'svelte'

    export let warriors = {}
    export let selectedWarrior

    const dispatch = createEventDispatcher()

    const warriorSelectionChanged = (strength) => {
        dispatch('warrior-selected', {
            strength,
        })
    }

    $: warriorSelectedClass = (strength) => selectedWarrior === strength ? 'warrior--selected': ''

    const findWeakestWarrior = () => {
        if (!warriors) {
            return undefined
        }
        if (warriors.one_point) {
            return "1"
        }
        if (warriors.two_points) {
            return "2"
        }
        if (warriors.three_points) {
            return "3"
        }
        if (warriors.four_points) {
            return "4"
        }
        if (warriors.five_points) {
            return "5"
        }
    }

    const selectWeakestWarrior = () => {
        const weakestWarrior = findWeakestWarrior()

        if (weakestWarrior) {
            warriorSelectionChanged(weakestWarrior)
        }
    }

    onMount(selectWeakestWarrior)
</script>

<article class="warrior-selection">
    <section class="warrior-selection__strength warrior-selection__strength--one-warriors">
        {#each Array(warriors.one_point || 0) as _}
        <button
            class={`warrior ${warriorSelectedClass("1")} player-action`}
            on:click={() => warriorSelectionChanged("1")}
        >1</button>
        {/each}
    </section>
    <section class="warrior-selection__strength warrior-selection__strength--two-warriors">
        {#each Array(warriors.two_points || 0) as _}
        <button
        class={`warrior ${warriorSelectedClass("2")} player-action`}
            on:click={() => warriorSelectionChanged("2")}
        >2</button>
        {/each}
    </section>
    <section class="warrior-selection__strength warrior-selection__strength--three-warriors">
        {#each Array(warriors.three_points || 0) as _}
        <button
        class={`warrior ${warriorSelectedClass("3")} player-action`}
            on:click={() => warriorSelectionChanged("3")}
        >3</button>
        {/each}
    </section>
    <section class="warrior-selection__strength warrior-selection__strength--four-warriors">
        {#each Array(warriors.four_points || 0) as _}
        <button
        class={`warrior ${warriorSelectedClass("4")} player-action`}
            on:click={() => warriorSelectionChanged("4")}
        >4</button>
        {/each}
    </section>
    <section class="warrior-selection__strength warrior-selection__strength--five-warriors">
        {#each Array(warriors.five_points || 0) as _}
        <button
        class={`warrior ${warriorSelectedClass("5")} player-action`}
            on:click={() => warriorSelectionChanged("5")}
        >5</button>
        {/each}
    </section>
</article>

<style>
.warrior--selected {
    background-color: var(--warrior-selected-background, #dab44a);
}

.warrior-selection {
    display: flex;
    --warrior-button-margin: 3em;
}

.warrior-selection__strength {
    margin-inline-end: var(--warrior-button-margin);

    display: flex;
    justify-content: center;
    align-items: center;
}

.warrior {
    margin-inline-end: calc(var(--warrior-button-margin) * -1);
}
</style>
