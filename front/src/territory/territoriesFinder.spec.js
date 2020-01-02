import { findTerritories } from './territoriesFinder';
import { createGame } from '../gameFactory';

test('initial grid', () => {
    // Given
    const { grid } = createGame();

    // When
    const gridWithTerritories = findTerritories(grid);

    // Then
    expect(gridWithTerritories.length).toBe(grid.length);
    gridWithTerritories.forEach((row, rowIndex) => {
        expect(row.length).toBe(grid[rowIndex].length)
        row.forEach(cell => {
            expect(cell.territoryId).toBe(0)
        })
    })
})

test('two territories', () => {
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
    const gridWithTerritories = findTerritories(grid);

    // Then
    let row;
    let col;
    
    for (row = 0; row < 2; row++) {
        for (col = 0; col < 3; col++) {
            expect(gridWithTerritories[row][col].territoryId).toBe(0);
        }
    }

    for (; row < gridWithTerritories.length; row++) {
        for (; col < gridWithTerritories[row].length; col++) {
            expect(gridWithTerritories[row][col].territoryId).not.toBe(0);
        }
    }
})
