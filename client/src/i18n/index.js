import { derived, writable } from 'svelte/store';
import enLabels from './en.json';
import frLabels from './fr.json';

export const labels = writable(enLabels);

const OBJECT_PROPERTY_SEPARATOR = "."

export const i18n = derived(labels, (newLabels) => {
    return (key) => {
        if (!key.includes(OBJECT_PROPERTY_SEPARATOR)) {
            return newLabels[key]
        }
        const objectToCrawl = key.split(OBJECT_PROPERTY_SEPARATOR)
        let currentPositionInLabels = newLabels
        for (let i = 0; i < objectToCrawl.length; i++) {
            currentPositionInLabels = currentPositionInLabels[objectToCrawl[i]]
            if (!currentPositionInLabels) {
                return ""
            }
        }
        return currentPositionInLabels
    }
})

export const EN_LOCALE = "en";
export const FR_LOCALE = "fr";
export let currentLocale = EN_LOCALE;

export const changeLocale = (newLocale) => {
    if (newLocale === EN_LOCALE) {
        labels.set(enLabels)
        currentLocale = newLocale
    } else if (newLocale === FR_LOCALE) {
        labels.set(frLabels)
        currentLocale = newLocale
    }
}
