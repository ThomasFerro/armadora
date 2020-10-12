export const getConnectedPlayerInformation = () => {
    return {
        nickname: localStorage.getItem('nickname')
    }
}

export const disconnect = () => {
    localStorage.removeItem('nickname')
}

export const setNickname = (nickname) => {
    localStorage.setItem('nickname', nickname)
}
