export const getPartyNameFromUrl = () => {
    const urlParams = new URLSearchParams(window.location.search);
    return urlParams.get('party')
}
