import { PUT_WARRIOR, PUT_PALISADES } from './actionTypes';

const removeUsedWarrior = ({currentPlayer, players}, selectedWarrior) => {
    const warriors = players[currentPlayer].warriors
        
    players[currentPlayer].warriors = [
        ...warriors.slice(0, selectedWarrior),
        ...warriors.slice(selectedWarrior + 1),
    ]
}

const nextCurrentPlayer = ({ currentPlayer, players }) => {
    if (currentPlayer >= players.length - 1) {
        return 0
    } else {
        return currentPlayer + 1
    }
}

const putWarrior = (game, {selectedWarrior, x, y}) => {
    console.log('putWarrior', game, x, y);
    const currentPlayerInformation = game.players[game.currentPlayer]
    const currentPlayerWarriors = currentPlayerInformation.warriors
    if (currentPlayerWarriors[selectedWarrior]) {
        game.grid[x][y].warrior = {
            player: game.currentPlayer,
            playerDisplayName: currentPlayerInformation.race,
            strength: currentPlayerWarriors[selectedWarrior]
        }
        removeUsedWarrior(game, selectedWarrior)
    }
}

const putPalisade = (game) => (informationPalisade) => {
    console.log('putPalisade', game.palisades, informationPalisade);

}

export const playTurn = (actualGame, {type, selectedWarrior, x, y, palisades}) => {
    // TODO: deep copy
    const game = {...actualGame}

    if (type === PUT_WARRIOR) {
        putWarrior(game, { selectedWarrior, x, y })
        game.currentPlayer = nextCurrentPlayer(game)
    }

    if (type === PUT_PALISADES) {
        palisades.forEach(putPalisade(game))
        game.currentPlayer = nextCurrentPlayer(game)
    }

    return game
}