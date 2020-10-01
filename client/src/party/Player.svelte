<script>
    export let player = {}

    const characterImageSource = (character) => {
        return `/static/images/character-${character && character.toLowerCase()}.svg`
    }
</script>

<img
    class="character-image"
    src={characterImageSource(player.character)}
    alt={player.character}
>{player.nickname}

<style>
.character-image {
    margin-inline-end: 1em;
    width: 40px;
}
</style>