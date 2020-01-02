import { findAreas } from './area/areaFinder';
import { validGrid } from './area/areaValidator';
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
    if (game.palisadesCount > 0) {
        const cell = game.grid[x][y]
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
    }
}

const computeGridAreas = (grid) => {
    const cleanGrid = grid.map(
        row => row.map(
            cell => ({
                ...cell,
                areaId: undefined
            })
        )
    );

    return findAreas(cleanGrid);
}

export const playTurn = (actualGame, {type, selectedWarrior, x, y, palisades}) => {
    const game = JSON.parse(JSON.stringify(actualGame))

    if (type === PUT_WARRIOR) {
        if (putWarrior(game, { selectedWarrior, x, y })) {
            game.currentPlayer = nextCurrentPlayer(game)
        }
    }

    if (type === PUT_PALISADES) {
        if (game.palisadesCount >= palisades.length) {
            palisades.forEach(putPalisade(game))
            const newGrid = computeGridAreas(game.grid)
            if (!validGrid(newGrid)) {
                return actualGame;
            }
            game.palisadesCount -= palisades.length
            game.grid = newGrid
            game.currentPlayer = nextCurrentPlayer(game)
        }
    }


    return game
}