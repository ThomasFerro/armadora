const url = process.env.API_URL;

const fetchWithDefaultCheck = (url, options) => fetch(url, options)
    .then(response => {
        if (!response.ok) {
            throw new Error(response.statusText)
        }
        return response.json()
    });

export const createNewParty = () => fetchWithDefaultCheck(`${url}/games`,{ method: "POST" })

export const getParties = () => fetchWithDefaultCheck(`${url}/parties`)

const partyUrl = (id) => `${url}/parties/${id}`;

export const gameInformation = (id) => fetchWithDefaultCheck(partyUrl(id))

export const connectToGame = (id) => ({username, character}) => fetchWithDefaultCheck(partyUrl(id), {
    method: 'POST',
    body: JSON.stringify({
        "command_type": "JoinGame",
        "payload": {
            "Nickname": username,
            "Character": character,
        },
    })
});

export const startGame = (id) => fetchWithDefaultCheck(partyUrl(id), {
    method: 'POST',
    body: JSON.stringify({
        "command_type": "StartTheGame",
    })
});

export const putWarrior = (id) => (player) => ({x, y, strength}) => fetchWithDefaultCheck(partyUrl(id), {
    method: 'POST',
    body: JSON.stringify({
        "command_type": "PutWarrior",
        "player": player,
        "payload": {
            "Warrior": strength,
            "X": x.toString(),
            "Y": y.toString(),
        },
    })
});

export const putPalisades = (id) => (player) => ({ palisades }) => fetchWithDefaultCheck(partyUrl(id), {
    method: 'POST',
    body: JSON.stringify({
        "command_type": "PutPalisades",
        "player": player,
        "payload": {
            "Palisades": JSON.stringify(palisades)
        },
    })
});

export const passTurn = (id) => (player) => fetchWithDefaultCheck(partyUrl(id), {
    method: 'POST',
    body: JSON.stringify({
        "command_type": "PassTurn",
        "player": player,
    })
});
