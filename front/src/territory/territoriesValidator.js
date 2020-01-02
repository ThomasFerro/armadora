export const validGrid = (grid) => {
    const territoryCounts = grid.reduce((acc, row) => {        
        row.forEach(cell => {
            let territoryCount = acc[cell.territoryId]
            if (!territoryCount) {
                acc[cell.territoryId] = 1
            } else {
                acc[cell.territoryId] += 1
            }
        })
        return acc
    }, {});
    return !Object.values(territoryCounts).find(count => count < 4);
}
