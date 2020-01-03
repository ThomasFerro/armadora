import { findTerritories } from '../territory/territoriesFinder';
import { validGrid } from '../territory/territoriesValidator';
import { PUT_WARRIOR, PUT_PALISADES } from '../action/actionTypes';
import { GOLD } from '../cell/cellTypes';

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

const computeGridTerritories = (grid) => {
    const cleanGrid = grid.map(
        row => row.map(
            cell => ({
                ...cell,
                territoryId: undefined
            })
        )
    );

    return findTerritories(cleanGrid);
}

const noMoreWarrior = (players) => {
    return !players.find(player => player.warriors.length !== 0)
}

const gridCellsByTerritoryId = (grid) => {
    return computeGridTerritories(grid)
        .reduce((acc, row) => {
            row.forEach(cell => {
                if (!acc[cell.territoryId]) {
                    acc[cell.territoryId] = []
                }
                acc[cell.territoryId] = [
                    ...acc[cell.territoryId],
                    cell
                ]
            })
            return acc
        }, {})
}

const sum = (count, element) => count + element

const computeTerritoryResults = (lands) => {
    const computedTerritory = lands.reduce((acc, land) => {
        if (land.type === GOLD) {
            acc.gold = acc.gold + land.pile
        } else if (land.warrior !== undefined) {
            if (!acc.playersStrength[land.warrior.player]) {
                acc.playersStrength[land.warrior.player] = []
            }
            acc.playersStrength[land.warrior.player] = [
                ...acc.playersStrength[land.warrior.player],
                land.warrior.strength
            ]
        }
        return acc
    }, {
        gold: 0,
        playersStrength: {}
    })

    const winners = Object.entries(computedTerritory.playersStrength)
        .reduce((acutalWinners, [ playerId, army ]) => {
            const armyStrength = army.reduce(sum, 0)
            if (acutalWinners.strength === 0) {
                acutalWinners.strength = armyStrength
            }
            if (acutalWinners.strength === armyStrength) {
                acutalWinners.players = [
                    ...acutalWinners.players,
                    playerId
                ]
            } else if (acutalWinners.strength < armyStrength) {
                acutalWinners.strength = armyStrength
                acutalWinners.players = [
                    playerId
                ]
            }
            return acutalWinners;
        }, {
            strength: 0,
            players: []
        })

    return {
        gold: computedTerritory.gold,
        winners: winners.players,
    };
}

const distributeGold = (territories) => {
    return territories.reduce(
        (distributedGold, {gold, winners}) => {
            const goldToDistribute = Math.floor(gold / winners.length)
            winners.forEach(winner => {
                if (!distributedGold[winner]) {
                    distributedGold[winner] = []
                }
                distributedGold[winner] = [
                    ...distributedGold[winner],
                    goldToDistribute
                ]
            })
            return distributedGold
        }
    ,  {});
}

const greatestSum = (elementsList) => {
    return elementsList.reduce((greater, [ id, elements ]) => {
        const sumOfElements = elements.reduce(sum, 0)
        if (greater === undefined) {
            return sumOfElements
        }
        return sumOfElements > greater ? sumOfElements : greater
    }, undefined)
}

const findWinner = (players) => {    
    const greatestSumOfGold = greatestSum(Object.entries(players))
    return Object
        .entries(players)
        .filter(([id, gold]) => gold.reduce(sum, 0) === greatestSumOfGold)
        // TODO: Egalité
        .map(([id]) => id)[0]
}

export const computeResults = (game) => {
    
    const computedTerritoriesResults = Object.entries(gridCellsByTerritoryId(game.grid))
        .map(([territoryId, lands]) => computeTerritoryResults(lands));

    const players = distributeGold(computedTerritoriesResults)
    
    return {
        winner: findWinner(players)
    }
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
            const newGrid = computeGridTerritories(game.grid)
            if (!validGrid(newGrid)) {
                return actualGame;
            }
            game.palisadesCount -= palisades.length
            game.grid = newGrid
            game.currentPlayer = nextCurrentPlayer(game)
        }
    }

    if (game.palisadesCount === 0 && noMoreWarrior(game.players)) {
        game.results = computeResults(game)
    }

    return game
}