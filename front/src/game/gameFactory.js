import { LAND, GOLD } from '../cell/cellTypes';

const cell = () => ({
    palisades: {}
})

const emptyCell = () => ({
    ...cell(),
	type: LAND,
})

const goldCell = () => ({
    ...cell(),
	type: GOLD,
})

// TODO: Create a game with various players
// TODO: Put the gold pile randomly
export const createGame = () => ({
    palisadesCount: 5,
    players: [
        {
            race: 'Orc',
            warriors: [
                1
            ]
        },
        {
            race: 'Goblin',
            warriors: [
                1
            ]
        },
        {
            race: 'Elf',
            warriors: [
                1
            ]
        },
        {
            race: 'Mage',
            warriors: [
                1
            ]
        },
    ],
    grid: [
        [ { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(), pile: 6 }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() } ],
        [ { ...emptyCell() }, { ...goldCell(), pile: 7 }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(), pile: 4 }, { ...emptyCell() }, { ...goldCell(), pile: 3 } ],
        [ { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() } ],
        [ { ...goldCell(), pile: 5 }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(), pile: 6 }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() } ],
        [ { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(), pile: 4 }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(), pile: 5 }, { ...emptyCell() } ],
    ],
    currentPlayer: 0,
})
// export const createGame = () => ({
//     palisadesCount: 35,
//     players: [
//         {
//             race: 'Orc',
//             warriors: [
//                 1, 1, 1, 1, 1, 2, 3, 4
//             ]
//         },
//         {
//             race: 'Goblin',
//             warriors: [
//                 1, 1, 1, 1, 1, 2, 3, 4
//             ]
//         },
//         {
//             race: 'Elf',
//             warriors: [
//                 1, 1, 1, 1, 1, 2, 3, 4
//             ]
//         },
//         {
//             race: 'Mage',
//             warriors: [
//                 1, 1, 1, 1, 1, 2, 3, 4
//             ]
//         },
//     ],
//     grid: [
//         [ { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(), pile: 6 }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() } ],
//         [ { ...emptyCell() }, { ...goldCell(), pile: 7 }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(), pile: 4 }, { ...emptyCell() }, { ...goldCell(), pile: 3 } ],
//         [ { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() } ],
//         [ { ...goldCell(), pile: 5 }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(), pile: 6 }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() } ],
//         [ { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(), pile: 4 }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(), pile: 5 }, { ...emptyCell() } ],
//     ],
//     currentPlayer: 0,
// })
