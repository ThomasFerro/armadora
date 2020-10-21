export const getPreferredLocale = () => {
    return localStorage.getItem('preferred-locale')
}

export const setPreferredLocale = (locale) => {
    localStorage.setItem('preferred-locale', locale)
}