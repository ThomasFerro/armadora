import { shuffle } from './shuffle';
import { LAND, GOLD } from '../cell/cellTypes';

const cell = () => ({
    palisades: {}
})

const emptyCell = () => ({
    ...cell(),
	type: LAND,
})

const goldCell = (pile) => ({
    ...cell(),
    type: GOLD,
    pile,
})

const warriorsConfiguration = {
    2: [ 5, 4, 3, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1 ],
    3: [ 4, 3, 2, 2, 1, 1, 1, 1, 1, 1, 1 ],
    4: [ 4, 3, 2, 1, 1, 1, 1, 1 ]
}

export const createGame = (requestedPlayers) => {
    const players = requestedPlayers.map(race => ({
        race,
        warriors: [
            ...warriorsConfiguration[requestedPlayers.length]
        ]
    }))

    const gold = shuffle([ 7, 6, 6, 5, 5, 4, 4, 3 ])

    return {
        palisadesCount: 35,
        players,
        grid: [
            [ { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(gold.pop()) }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() } ],
            [ { ...emptyCell() }, { ...goldCell(gold.pop()) }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(gold.pop()) }, { ...emptyCell() }, { ...goldCell(gold.pop()) } ],
            [ { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() } ],
            [ { ...goldCell(gold.pop()) }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(gold.pop()) }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() } ],
            [ { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(gold.pop()) }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(gold.pop()) }, { ...emptyCell() } ],
        ],
        currentPlayer: 0,
    }
}
