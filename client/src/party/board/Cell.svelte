<script>
    export let value = {}
    export let active = false
    export let area = ""

    $: cellType = value.type

    $: warriorCharacter = value.type === 'warrior'
        ? `warrior--${value.character.toLowerCase()}`
        : '';
    
    $: goldDisplayed = value.type === 'gold'
        ? `gold--${value.gold}`
        : '';

    $: disabled = !active || value.type !== 'land'
</script>

<button class={`cell ${area} cell--${cellType} ${warriorCharacter} ${goldDisplayed}`} style={`grid-area: ${area}`} on:click {disabled}></button>

<style>
    .cell {
        width: 100%;
        height: 100%;
        border: 1px solid black;
        margin: 0;
        padding: 0;
    }

    .cell--land {
        background-color: var(--empty-land-background, #74c328e8);
    }

    .cell--land:hover {
        background-color: var(--empty-land-background-darken, #5c9c20e8);
    }

    .cell--land:disabled, .cell--warrior, .cell--gold {
        background-color: var(--empty-land-disabled-background, #74c32891);
    }

    .cell--gold {
        background-size: contain;
        background-position: center;
        background-repeat: no-repeat;
    }

    .gold--3 {
        background-image: url('/static/images/gold-3.svg');
    }

    .gold--4 {
        background-image: url('/static/images/gold-4.svg');
    }

    .gold--5 {
        background-image: url('/static/images/gold-5.svg');
    }

    .gold--6 {
        background-image: url('/static/images/gold-6.svg');
    }

    .gold--7 {
        background-image: url('/static/images/gold-7.svg');
    }

    .cell--warrior {
        background-size: 70%;
        background-position: center;
        background-repeat: no-repeat;
    }

    .warrior--goblin {
        background-image: url('/static/images/character-goblin.svg');
    }

    .warrior--elf {
        background-image: url('/static/images/character-elf.svg');
    }

    .warrior--mage {
        background-image: url('/static/images/character-mage.svg');
    }

    .warrior--orc {
        background-image: url('/static/images/character-orc.svg');
    }
</style>
