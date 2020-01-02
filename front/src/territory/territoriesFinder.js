const processCell = (grid, row, cell, nextTerritoryId) => {
    if (grid[row][cell].territoryId !== undefined) {
        return;
    }

    grid[row][cell].territoryId = nextTerritoryId;

    if (cell > 0 && !grid[row][cell].palisades.left) {
        processCell(grid, row, cell - 1, nextTerritoryId);
    }

    if (cell < grid[row].length - 1 && !grid[row][cell].palisades.right) {
        processCell(grid, row, cell + 1, nextTerritoryId);
    }

    if (row > 0 && !grid[row][cell].palisades.top) {
        processCell(grid, row - 1, cell, nextTerritoryId);
    }

    if (row < grid.length - 1 && !grid[row][cell].palisades.bottom) {
        processCell(grid, row + 1, cell, nextTerritoryId);
    }
}

export const findTerritories = (grid) => {
    
    let nextTerritoryId = 0;
    const gridWithTerritories = JSON.parse(JSON.stringify(grid))
    for (let row = 0; row < gridWithTerritories.length; row++) {
        for (let col = 0; col < gridWithTerritories[row].length; col++) {
            processCell(gridWithTerritories, row, col, nextTerritoryId)
            nextTerritoryId++
        }
    }
    return gridWithTerritories
}
