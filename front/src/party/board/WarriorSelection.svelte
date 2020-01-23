<script>
    import { createEventDispatcher, onMount } from 'svelte'

    export let warriors = {}

    let select

    const dispatch = createEventDispatcher()

    const warriorSelectionChanged = (strength) => {
        dispatch('warrior-selected', {
            strength,
        })
    }

    onMount(() => {
        warriorSelectionChanged(select.value)
    })
</script>

<select bind:this={select} class="warrior-selection" on:change={(e) => warriorSelectionChanged(e.target.value)}>
    <option disabled>Select a warrior</option>
    {#if warriors.one_point}
    <option value="1">1 ({warriors.one_point} left)</option>
    {/if}
    {#if warriors.two_points}
    <option value="2">2 ({warriors.two_points} left)</option>
    {/if}
    {#if warriors.three_points}
    <option value="3">3 ({warriors.three_points} left)</option>
    {/if}
    {#if warriors.four_points}
    <option value="4">4 ({warriors.four_points} left)</option>
    {/if}
    {#if warriors.five_points}
    <option value="5">5 ({warriors.five_points} left)</option>
    {/if}
</select>