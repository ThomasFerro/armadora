// TODO: Configure ?
const url = `http://server`;

export const createNewParty = () => fetch(`${url}/games`,{ method: "POST" })
    .then(response => response.json());

export const getParties = () => fetch(`${url}/parties`)
    .then(response => response.json());

const partyUrl = (id) => `${url}/parties/${id}`;

export const gameInformation = (id) => fetch(partyUrl(id))
    .then(response => response.json());

export const connectToGame = (id) => ({username, character}) => fetch(partyUrl(id), {
    method: 'POST',
    body: JSON.stringify({
        "command_type": "JoinGame",
        "payload": {
            "Nickname": username,
            "Character": character,
        },
    })
}).then(response => response.json());

export const startGame = (id) => fetch(partyUrl(id), {
    method: 'POST',
    body: JSON.stringify({
        "command_type": "StartTheGame",
    })
}).then(response => response.json());

export const putWarrior = (id) => ({x, y, strength}) => fetch(partyUrl(id), {
    method: 'POST',
    body: JSON.stringify({
        "command_type": "PutWarrior",
        "payload": {
            "Warrior": strength,
            "X": x.toString(),
            "Y": y.toString(),
        },
    })
}).then(response => response.json());

export const putPalisades = (id) => ({ palisades }) => fetch(partyUrl(id), {
    method: 'POST',
    body: JSON.stringify({
        "command_type": "PutPalisades",
        "payload": {
            "Palisades": JSON.stringify(palisades)
        },
    })
}).then(response => response.json());

export const passTurn = (id) => fetch(partyUrl(id), {
    method: 'POST',
    body: JSON.stringify({
        "command_type": "PassTurn",
    })
})
