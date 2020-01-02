const processCell = (grid, row, cell, nextAreaId) => {
    if (grid[row][cell].areaId !== undefined) {
        return;
    }

    grid[row][cell].areaId = nextAreaId;

    if (cell > 0 && !grid[row][cell].palisades.left) {
        processCell(grid, row, cell - 1, nextAreaId);
    }

    if (cell < grid[row].length - 1 && !grid[row][cell].palisades.right) {
        processCell(grid, row, cell + 1, nextAreaId);
    }

    if (row > 0 && !grid[row][cell].palisades.top) {
        processCell(grid, row - 1, cell, nextAreaId);
    }

    if (row < grid.length - 1 && !grid[row][cell].palisades.bottom) {
        processCell(grid, row + 1, cell, nextAreaId);
    }
}

export const findAreas = (grid) => {
    
    let nextAreaId = 0;
    const gridWithAreas = JSON.parse(JSON.stringify(grid))
    for (let row = 0; row < gridWithAreas.length; row++) {
        for (let col = 0; col < gridWithAreas[row].length; col++) {
            processCell(gridWithAreas, row, col, nextAreaId)
            nextAreaId++
        }
    }
    return gridWithAreas
}
