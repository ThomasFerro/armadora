import { computeResults } from './gameEngine';
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

test('simple resolution', () => {
    // Given
    const game = {
        players: [
            {
                race: 'Orc',
            },
            {
                race: 'Elf'
            }
        ],
        grid: [
            [{ ...emptyCell(), warrior: { player: 1, strength: 2 } }, { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(), pile: 6 }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }],
            [{ ...emptyCell(), warrior: { player: 0, strength: 1 } }, { ...goldCell(), pile: 7 }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(), pile: 4 }, { ...emptyCell() }, { ...goldCell(), pile: 3 }],
            [{ ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }],
            [{ ...goldCell(), pile: 5 }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(), pile: 6 }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }],
            [{ ...emptyCell() }, { ...emptyCell() }, { ...goldCell(), pile: 4 }, { ...emptyCell() }, { ...emptyCell() }, { ...emptyCell() }, { ...goldCell(), pile: 5 }, { ...emptyCell() }],
        ],
    };

    // When
    const results = computeResults(game);

    // Then
    expect(results.winner).toBe("1")
})
