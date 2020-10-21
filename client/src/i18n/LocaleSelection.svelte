<script>
    import { changeLocale, EN_LOCALE, FR_LOCALE, i18n } from ".";
    import { getPreferredLocale, setPreferredLocale } from './storage';

    let selectedLocale = getPreferredLocale()

    $: {
        if (selectedLocale) {
            changeLocale(selectedLocale)
            setPreferredLocale(selectedLocale)
        }
    }
</script>

<label>
    {$i18n('localeSelection.select')}:
    <select class="player-action" bind:value={selectedLocale}>
        <option value={EN_LOCALE} aria-label="English">🇬🇧</option>
        <option value={FR_LOCALE} aria-label="French">🇫🇷</option>
    </select>
</label>
