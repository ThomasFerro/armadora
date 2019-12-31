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
        return true
    }
    return false
}

const putPalisade = (game) => ({x, y, vertical}) => {
    
    const cell = game.grid[x][y]
    console.log('putPalisade', game.palisades, cell);
    if (vertical) {
        cell.palisades = {
            ...cell.palisades,
            right: true
        }
        const rightNeighbor = game.grid[x][y + 1]
        rightNeighbor.palisades = {
            ...rightNeighbor.palisades,
            left: true
        }
    } else {
        cell.palisades = {
            ...cell.palisades,
            bottom: true
        }
        const bottomNeighbor = game.grid[x + 1][y]
        bottomNeighbor.palisades = {
            ...bottomNeighbor.palisades,
            top: true
        }
    }
    
    if (game.palisadesCount > 0) {
        game.palisadesCount--
    }
}

export const playTurn = (actualGame, {type, selectedWarrior, x, y, palisades}) => {
    // TODO: deep copy
    const game = {...actualGame}

    if (type === PUT_WARRIOR) {
        if (putWarrior(game, { selectedWarrior, x, y })) {
            game.currentPlayer = nextCurrentPlayer(game)
        }
    }

    if (type === PUT_PALISADES) {
        palisades.forEach(putPalisade(game))
        // TODO : Error management
        game.currentPlayer = nextCurrentPlayer(game)
    }

    return game
}