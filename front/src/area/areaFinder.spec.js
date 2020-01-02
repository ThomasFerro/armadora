import { findAreas } from './areaFinder';
import { createGame } from '../gameFactory';

test('initial grid', () => {
    // Given
    const { grid } = createGame();

    // When
    const gridWithAreas = findAreas(grid);

    // Then
    expect(gridWithAreas.length).toBe(grid.length);
    gridWithAreas.forEach((row, rowIndex) => {
        expect(row.length).toBe(grid[rowIndex].length)
        row.forEach(cell => {
            expect(cell.areaId).toBe(0)
        })
    })
})

test('two areas', () => {
    // Given
    const { grid } = createGame();

    grid[0][0].palisades.bottom = true;
    grid[0][1].palisades.bottom = true;
    grid[0][2].palisades.right = true;
    grid[1][0].palisades.top = true;
    grid[1][0].palisades.bottom = true;
    grid[1][1].palisades.top = true;
    grid[1][1].palisades.bottom = true;
    grid[1][2].palisades.right = true;
    grid[1][2].palisades.bottom = true;

    // When
    const gridWithAreas = findAreas(grid);

    // Then
    let row;
    let col;
    
    for (row = 0; row < 2; row++) {
        for (col = 0; col < 3; col++) {
            expect(gridWithAreas[row][col].areaId).toBe(0);
        }
    }

    for (; row < gridWithAreas.length; row++) {
        for (; col < gridWithAreas[row].length; col++) {
            expect(gridWithAreas[row][col].areaId).not.toBe(0);
        }
    }
})
